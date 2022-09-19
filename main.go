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

	kitchen_elem.OrdersChannel <- unCookedOrder
	c.IndentedJSON(http.StatusCreated, unCookedOrder)
}

func distributeFood() {
	for order := range kitchen_elem.OrdersChannel {

		kitchenOrder := kitchen_elem.OrderInKitchen{
			Id:    order.OrderId,
			Foods: make([]kitchen_elem.FoodToCook, 0, len(order.Items)),
			// new creates a pointer to WaitGroup
			Wg: new(sync.WaitGroup),
		}

		kitchenOrder.Wg.Add(len(order.Items))

		for _, foodID := range order.Items {

			newFood := kitchen_elem.FoodToCook{
				OrderId: order.OrderId,
				FoodId:  foodID,

				//kitchenOrder.Wg is already a pointer
				Wg: kitchenOrder.Wg,
			}

			switch complexity := kitchen_elem.Foods[foodID-1].Complexity; complexity {
			//i can use here another factor; cooks's freeness , depending on a channel if he is free(the profficiency channel) or not; but for this i should sort by complexity descending
			case 3:
				kitchen_elem.Cooks[0].FoodChan <- newFood
				newFood.CookId = 0

			case 2:
				if rand.Intn(2) == 0 {
					kitchen_elem.Cooks[1].FoodChan <- newFood
					newFood.CookId = 1
				} else {
					kitchen_elem.Cooks[2].FoodChan <- newFood
					newFood.CookId = 2
				}

			case 1:
				kitchen_elem.Cooks[3].FoodChan <- newFood
				newFood.CookId = 3

			}
			kitchenOrder.Foods = append(kitchenOrder.Foods, newFood)
		}

		go kitchenOrder.WaitForOrder(order, kitchenOrder.Foods)
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
