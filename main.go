package main

import (
	"github.com/EliriaT/kitchen/kitchen-elem"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	kitchen_elem.OrdersChannel <- unCookedOrder.OrderId
	kitchen_elem.OrderMap[unCookedOrder.OrderId] = unCookedOrder
	c.IndentedJSON(http.StatusCreated, unCookedOrder)
}

func main() {

	router := gin.Default()
	router.GET("/cooks", getCooks)
	router.POST("/order", receivedOrder)
	//router.POST("/order", serveOrder)
	for i, _ := range kitchen_elem.Cooks {
		go kitchen_elem.Cooks[i].LookUpOrders()
	}
	router.Run(":8080")

}
