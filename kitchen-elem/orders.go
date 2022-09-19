package kitchen_elem

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

// the order response sent back to dinning-hall
type SentOrd struct {
	OrderId        int              `json:"order_id"`
	TableId        int              `json:"table_id"`
	WaiterId       int              `json:"waiter_id"`
	Items          []int            `json:"items"`
	Priority       int              `json:"priority"`
	MaxWait        float64          `json:"max_wait"`
	PickUpTime     time.Time        `json:"pick_up_time"`
	CookingTime    time.Duration    `json:"cooking_time"`
	CookingDetails []KitchenFoodInf `json:"cooking_details"`
}

// the response received from dinning hall
type ReceivedOrd struct {
	OrderId    int       `json:"order_id"`
	TableId    int       `json:"table_id"`
	WaiterId   int       `json:"waiter_id"`
	Items      []int     `json:"items"`
	Priority   int       `json:"priority"`
	MaxWait    float64   `json:"max_wait"`
	PickUpTime time.Time `json:"pick_up_time"`
}

type OrderInKitchen struct {
	Id    int
	Foods []FoodToCook
	Wg    *sync.WaitGroup
}

func (o *OrderInKitchen) WaitForOrder(initialOrder ReceivedOrd, cookedFoods []FoodToCook) {
	cookingTime := time.Now()
	o.Wg.Wait()
	var cookedOrder SentOrd

	//SA FAC CA SA TRANSMIT COOKID
	var foodCookedInfo = make([]KitchenFoodInf, 0, len(initialOrder.Items))

	for _, food := range cookedFoods {
		foodCookedInfo = append(foodCookedInfo, KitchenFoodInf{
			FoodId: food.FoodId,
			CookId: food.CookId,
		})
	}

	cookedOrder.OrderId = initialOrder.OrderId
	cookedOrder.TableId = initialOrder.TableId
	cookedOrder.WaiterId = initialOrder.WaiterId
	cookedOrder.Items = initialOrder.Items
	cookedOrder.Priority = initialOrder.Priority
	cookedOrder.MaxWait = initialOrder.MaxWait
	cookedOrder.PickUpTime = initialOrder.PickUpTime
	cookedOrder.CookingTime = time.Since(cookingTime)
	cookedOrder.CookingDetails = foodCookedInfo

	o.sendOrder(cookedOrder)
}

func (o *OrderInKitchen) sendOrder(cookedOrder SentOrd) {
	reqBody, err := json.Marshal(cookedOrder)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	resp, err := http.Post("http://localhost:8082/distribution", "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		log.Printf("Request Failed: %s", err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf(err.Error())
			return
		}
	}(resp.Body)

	log.Printf("The order with id %d was sent to Dinning Hall .", cookedOrder.OrderId) // Unmarshal result

}

// waiter's goroutine receive orders on channel
var OrdersChannel = make(chan ReceivedOrd, 20)
