package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const timeUnit = 50 * time.Millisecond

func getCooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cooks)
}

func receivedOrder(c *gin.Context) {
	var unCookedOrder receivedOrd
	if err := c.BindJSON(&unCookedOrder); err != nil {
		log.Printf(err.Error())
	}

	log.Printf("Order with ID %d, arrived at kitchen", unCookedOrder.OrderId)
	ordersChannel <- unCookedOrder.OrderId
	orderMap[unCookedOrder.OrderId] = unCookedOrder
	c.IndentedJSON(http.StatusCreated, unCookedOrder)
}

func main() {

	rand.Seed(time.Now().UnixNano())
	router := gin.Default()
	router.GET("/cooks", getCooks)
	router.POST("/order", receivedOrder)
	//router.POST("/order", serveOrder)

	router.Run("localhost:8082")

}
