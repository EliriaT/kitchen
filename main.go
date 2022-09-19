package main

import (
	"github.com/EliriaT/kitchen/kitchen-elem"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const TimeUnit = 50 * time.Millisecond

func getCooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, kitchen_elem.Cooks)
}

// handler function for post request
func receivedOrder(c *gin.Context) {
	var unCookedOrder kitchen_elem.ReceivedOrd
	if err := c.BindJSON(&unCookedOrder); err != nil {
		log.Printf(err.Error())
	}

	log.Printf("Order with ID %d, arrived at kitchen", unCookedOrder.OrderId)

	//locking to avoid  concurrent map read and map write error
	//kitchen_elem.OrderMapMutex.Lock()
	//kitchen_elem.OrderMap[unCookedOrder.OrderId] = unCookedOrder

	//kitchen_elem.OrderMapMutex.Unlock()

	kitchen_elem.OrdersChannel <- unCookedOrder
	c.IndentedJSON(http.StatusCreated, unCookedOrder)
}

func distributeFood() {
	for order := range kitchen_elem.OrdersChannel {
		// I dont need the Foods list probably
		//creating as a pointer because it is further transmitted intro a channel. the channel should be of type pointer?
		kitchenOrder := kitchen_elem.OrderInKitchen{
			Id:    order.OrderId,
			Foods: make([]kitchen_elem.FoodToCook, len(order.Items)),
			// new creates a pointer to WaitGroup
			Wg: new(sync.WaitGroup),
		}
		kitchenOrder.Wg.Add(len(order.Items))
		//here also sent be value

		for _, foodID := range order.Items {
			//creating a pointer because changes should be saved?
			newFood := kitchen_elem.FoodToCook{
				OrderId: order.OrderId,
				FoodId:  foodID,
				//kitchenOrder.Wg is already a pointer
				Wg: kitchenOrder.Wg,
			}

			switch complexity := kitchen_elem.Foods[foodID-1].Complexity; complexity {
			case 3:
				kitchen_elem.Cooks[0].FoodChan <- newFood

			case 2:
				if rand.Intn(2) == 0 {
					kitchen_elem.Cooks[1].FoodChan <- newFood
				} else {
					kitchen_elem.Cooks[2].FoodChan <- newFood
				}

			case 1:
				kitchen_elem.Cooks[3].FoodChan <- newFood

			}
		}

		go kitchenOrder.WaitForOrder(order)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())
	router := gin.Default()
	router.GET("/cooks", getCooks)
	router.POST("/order", receivedOrder)

	go distributeFood()
	for i, _ := range kitchen_elem.Cooks {
		go kitchen_elem.Cooks[i].ListenForFood()
	}
	router.Run(":8080")

}
