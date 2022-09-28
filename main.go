package main

import (
	"encoding/json"
	"github.com/EliriaT/kitchen/kitchen-elem"
	"github.com/gorilla/mux"
	"runtime"

	//pq "github.com/kyroy/priority-queue"
	//_ "go.uber.org/automaxprocs"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

//
//var pq = make(kitchen_elem.PriorityQueue, 10)

func getCooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonFoods, err := json.Marshal(kitchen_elem.Cooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//by default sends 200
	w.Write(jsonFoods)
}

func receiveOrder(w http.ResponseWriter, r *http.Request) {
	var unCookedOrder kitchen_elem.ReceivedOrd
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&unCookedOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	jsonUncookedOrder, _ := json.Marshal(unCookedOrder)

	kitchen_elem.OrdersChannel <- unCookedOrder

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUncookedOrder)

	log.Printf("Order with ID %d, arrived at kitchen", unCookedOrder.OrderId)

}

func listenForOrders() {

	//queue := pq.NewPriorityQueue()

	//here to use smth like a signal to notify when a cook is free. As long as a cook is free we can distribute the order with the most priority
	for order := range kitchen_elem.OrdersChannel {

		//generating a kitchen order
		kitchenOrder := kitchen_elem.OrderInKitchen{
			Id:            order.OrderId,
			Foods:         make([]kitchen_elem.FoodToCook, 0, len(order.Items)),
			ReceivedOrder: order,
			// new creates a pointer to WaitGroup
			Wg:       new(sync.WaitGroup),
			Priority: order.Priority,
		}
		//to wait for food prep
		kitchenOrder.Wg.Add(len(order.Items))

		//generating the food to cook
		for _, foodId := range order.Items {
			newFood := kitchen_elem.FoodToCook{
				OrderId:          order.OrderId,
				FoodId:           foodId,
				CookingApparatus: kitchen_elem.Foods[foodId-1].CookingApparatus,
				//kitchenOrder.Wg is already a pointer
				Wg: kitchenOrder.Wg,
			}
			kitchenOrder.Foods = append(kitchenOrder.Foods, newFood)
		}
		//log.Println(kitchenOrder)
		//push to the priority queue

		//queue.Insert(kitchenOrder, float64(kitchenOrder.Priority))

		//If no cook is free, then take another order
		//select {
		//case kitchen_elem.CookFree <- 1:
		//	<-kitchen_elem.CookFree
		//default:
		//	continue
		//}

		//if len(kitchen_elem.CookFree) == 11 {
		//	continue
		////}

		//take the order with best priority (1 the best))  [IT SHOULD ALSO BE SORTED BY TIME?]
		//Further we should work ONLY with order !
		//order := queue.PopLowest().(kitchen_elem.OrderInKitchen)
		//log.Println(order)
		order := kitchenOrder
		for _, foodToCook := range order.Foods {

			foodID := foodToCook.FoodId

			switch complexity := kitchen_elem.Foods[foodID-1].Complexity; complexity {
			//i can use here another factor; cooks's freeness , depending on a channel if he is free(the profficiency channel) or not; but for this i should sort by complexity descending
			case 3:
				select {
				case kitchen_elem.Cooks[0].ProfficiencyChan <- 1:
					<-kitchen_elem.Cooks[0].ProfficiencyChan
					foodToCook.CookId = 1
					kitchen_elem.Cooks[0].FoodChan <- foodToCook

				default:
					foodToCook.CookId = 1
					kitchen_elem.Cooks[0].FoodChan <- foodToCook
				}

			case 2:
				select {
				case kitchen_elem.Cooks[1].ProfficiencyChan <- 1:
					<-kitchen_elem.Cooks[1].ProfficiencyChan
					foodToCook.CookId = 2
					kitchen_elem.Cooks[1].FoodChan <- foodToCook

				case kitchen_elem.Cooks[2].ProfficiencyChan <- 1:
					<-kitchen_elem.Cooks[2].ProfficiencyChan
					foodToCook.CookId = 3
					kitchen_elem.Cooks[2].FoodChan <- foodToCook

				case kitchen_elem.Cooks[0].ProfficiencyChan <- 1:
					<-kitchen_elem.Cooks[0].ProfficiencyChan
					foodToCook.CookId = 1
					kitchen_elem.Cooks[0].FoodChan <- foodToCook

				default:
					foodToCook.CookId = 2
					kitchen_elem.Cooks[1].FoodChan <- foodToCook
				}
				//if rand.Intn(2) == 0 {
				//	foodToCook.CookId = 2
				//	kitchen_elem.Cooks[1].FoodChan <- foodToCook
				//
				//} else {
				//	foodToCook.CookId = 3
				//	kitchen_elem.Cooks[2].FoodChan <- foodToCook
				//
				//}

			case 1:
				select {
				case kitchen_elem.Cooks[1].ProfficiencyChan <- 1:
					<-kitchen_elem.Cooks[1].ProfficiencyChan
					foodToCook.CookId = 2
					kitchen_elem.Cooks[1].FoodChan <- foodToCook

				case kitchen_elem.Cooks[2].ProfficiencyChan <- 1:
					<-kitchen_elem.Cooks[2].ProfficiencyChan
					foodToCook.CookId = 3
					kitchen_elem.Cooks[2].FoodChan <- foodToCook

				case kitchen_elem.Cooks[0].ProfficiencyChan <- 1:
					<-kitchen_elem.Cooks[0].ProfficiencyChan
					foodToCook.CookId = 1
					kitchen_elem.Cooks[0].FoodChan <- foodToCook

				case kitchen_elem.Cooks[3].ProfficiencyChan <- 1:
					<-kitchen_elem.Cooks[3].ProfficiencyChan
					foodToCook.CookId = 4
					kitchen_elem.Cooks[3].FoodChan <- foodToCook
				default:
					//foodToCook.CookId = 1
					//kitchen_elem.Cooks[0].FoodChan <- foodToCook
					foodToCook.CookId = 4
					kitchen_elem.Cooks[3].FoodChan <- foodToCook
				}

			}

		}
		//after all foods were sent to cooks, we can wait for it to be prepared
		go kitchenOrder.WaitForOrder(order.Foods)
	}
}

func main() {
	//fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(8)
	rand.Seed(time.Now().UnixNano())
	kitchen_elem.InitiateApparatus()

	r := mux.NewRouter()
	r.HandleFunc("/", getCooks).Methods("GET")
	r.HandleFunc("/order", receiveOrder).Methods("POST")

	go listenForOrders()
	for i, _ := range kitchen_elem.Cooks {
		go kitchen_elem.Cooks[i].ListenForFood()
	}
	log.Println("Kitchen server started..")
	log.Fatal(http.ListenAndServe(":8080", r))

}
