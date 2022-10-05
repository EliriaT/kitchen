package main

import (
	"encoding/json"
	"github.com/EliriaT/kitchen/dataStructures"
	"github.com/EliriaT/kitchen/kitchen-elem"
	"github.com/gorilla/mux"
	"runtime"

	//_ "go.uber.org/automaxprocs"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

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

	//sending in the general orderds channel the received order from dinning-hall
	kitchen_elem.OrdersChannel <- unCookedOrder

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUncookedOrder)

	log.Printf("Order with ID %d, arrived at kitchen", unCookedOrder.OrderId)

}

func listenForOrders() {

	autoLockMutex := false
	var lowestPriority uint8 = 5 //highest is 1
	queue := dataStructures.NewHierarchicalQueue(lowestPriority, autoLockMutex)

	//constantly listening for this channel
	for order := range kitchen_elem.OrdersChannel {

		//generating a kitchen order
		kitchenOrder := kitchen_elem.OrderInKitchen{
			Id:            order.OrderId,
			Foods:         make([]kitchen_elem.FoodToCook, 0, len(order.Items)),
			ReceivedOrder: order,
			// new creates a pointer to WaitGroup
			Wg:       new(sync.WaitGroup),
			Priority: uint8(order.Priority),
		}

		//to wait for food prep, we add number of foods to wait for
		kitchenOrder.Wg.Add(len(order.Items))

		//generating the food to cook by cooks, this will be sent to cooks
		for _, foodId := range order.Items {
			newFood := kitchen_elem.FoodToCook{
				OrderId:          order.OrderId,
				FoodId:           foodId,
				CookingApparatus: kitchen_elem.Foods[foodId-1].CookingApparatus,
				PrepTime:         kitchen_elem.Foods[foodId-1].PreparationTime,
				//kitchenOrder.Wg is already a pointer
				Wg: kitchenOrder.Wg,
			}
			kitchenOrder.Foods = append(kitchenOrder.Foods, newFood)
		}

		//push to the priority queue
		queue.Enqueue(kitchenOrder, kitchenOrder.Priority)

		//If no cook is free, then take another order
		// TODO HERE BETTER USE PROFFICIENCY CHANNEL ?
		select {
		case kitchen_elem.CookFree <- 1:
			<-kitchen_elem.CookFree

		default:
			continue
		}

		//TODO have here a load balancing

		//take the order with best priority (1 the best))
		orderInterface, _ := queue.Dequeue()
		kitchenOrder, _ = orderInterface.(kitchen_elem.OrderInKitchen)

		kitchen_elem.SendFoodsToCooks(kitchenOrder)
		// waiting for the foods to be prepared
		go kitchenOrder.WaitForOrder(kitchenOrder.Foods)

	}
}

func main() {
	//fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(6)

	rand.Seed(time.Now().UnixNano())
	kitchen_elem.StartWorkDay()

	for i := 0; i < kitchen_elem.Stoves.Quantity; i++ {

		go kitchen_elem.Stoves.CookFood()
	}

	for i := 0; i < kitchen_elem.Ovens.Quantity; i++ {

		go kitchen_elem.Ovens.CookFood()
	}
	go listenForOrders()
	for i, _ := range kitchen_elem.Cooks {
		go kitchen_elem.Cooks[i].ListenForFood()
	}

	r := mux.NewRouter()
	r.HandleFunc("/", getCooks).Methods("GET")
	r.HandleFunc("/order", receiveOrder).Methods("POST")

	log.Println("Kitchen server started..")
	log.Println("Quantum Apparatus: ", kitchen_elem.ApparatusQuantum)
	log.Fatal(http.ListenAndServe(":8080", r))

}
