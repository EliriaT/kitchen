package kitchen_elem

import (
	"log"
	"time"
)

type cook struct {
	Id               int    `json:"cook_id"`
	Rank             int    `json:"rank"`
	Proficiency      int    `json:"proficiency"`
	Name             string `json:"name"`
	CatchPhrase      string `json:"catch_phrase"`
	FoodChan         chan FoodToCook
	ProfficiencyChan chan int
}

func (c *cook) ListenForFood() {
	for food := range c.FoodChan {
		//this does not matter, i can send any number, used to control profficiency
		c.ProfficiencyChan <- food.FoodId
		//it is sent by value
		log.Printf("Cook %d cooks food %d", c.Id, food.FoodId)
		go c.cookFood(food)
	}
}

// i send it here by value
func (c *cook) cookFood(food FoodToCook) {
	time.Sleep(TimeUnit * time.Duration(Foods[food.FoodId-1].PreparationTime))
	food.Wg.Done()
	<-c.ProfficiencyChan

}

// continously listening to OrdersChannel
//func (c *cook) LookUpOrders() {
//
//	for i := range OrdersChannel {
//		cookedOrder := c.cookOrder(i)
//		c.sendOrder(cookedOrder)
//
//	}
//}

// Each cook , makes an entire order. Not optimal, but for simplicity
// maybe the RWMutex wasn't installed?
//func (c *cook) cookOrder(orderId int) SentOrd {
//	cookingTime := time.Now()
//	var cookedOrder SentOrd
//	OrderMapMutex.Lock()
//	initialOrder := OrderMap[orderId]
//	OrderMapMutex.Unlock()
//
//	//Wait for the food to cook; foodID is the actual index
//	for _, foodID := range initialOrder.Items {
//		time.Sleep(TimeUnit * time.Duration(Foods[foodID-1].PreparationTime))
//	}
//
//	var foodCookedInfo = make([]KitchenFoodInf, 0, len(initialOrder.Items)) //can be done faster
//	for _, foodID := range initialOrder.Items {
//		foodCookedInfo = append(foodCookedInfo, KitchenFoodInf{
//			FoodId: foodID,
//			CookId: c.Id,
//		})
//	}
//
//	//remove the order from the list ; maybe i can combine with previous lock..?
//	OrderMapMutex.Lock()
//	delete(OrderMap, orderId)
//	OrderMapMutex.Unlock()
//
//	cookedOrder.OrderId = initialOrder.OrderId
//	cookedOrder.TableId = initialOrder.TableId
//	cookedOrder.WaiterId = initialOrder.WaiterId
//	cookedOrder.Items = initialOrder.Items
//	cookedOrder.Priority = initialOrder.Priority
//	cookedOrder.MaxWait = initialOrder.MaxWait
//	cookedOrder.PickUpTime = initialOrder.PickUpTime
//	cookedOrder.CookingTime = time.Since(cookingTime)
//	cookedOrder.CookingDetails = foodCookedInfo
//	//log.Printf("cooking details: %+v", cookedOrder.CookingDetails)
//	return cookedOrder
//}
//
//// send order back to dinning hall server
//func (c *cook) sendOrder(cookedOrder SentOrd) {
//	reqBody, err := json.Marshal(cookedOrder)
//	if err != nil {
//		log.Printf(err.Error())
//		return
//	}
//
//	resp, err := http.Post("http://localhost:8082/distribution", "application/json", bytes.NewBuffer(reqBody))
//
//	if err != nil {
//		log.Printf("Request Failed: %s", err.Error())
//		return
//	}
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			log.Printf(err.Error())
//			return
//		}
//	}(resp.Body)
//	//body, err := io.ReadAll(resp.Body) // Log the request body
//	//if err != nil {
//	//	log.Printf("Can't read the response body %s", err.Error())
//	//	return
//	//}
//	//bodyString := string(body)
//	//log.Print(bodyString)
//	log.Printf("The order with id %d was sent to Dinning Hall .", cookedOrder.OrderId) // Unmarshal result
//
//}

var Cooks = []cook{
	{Id: 1, Rank: 3, Proficiency: 4, Name: "Mike", CatchPhrase: "I like ice-creams!", FoodChan: make(chan FoodToCook, 100), ProfficiencyChan: make(chan int, 4)},
	{Id: 2, Rank: 2, Proficiency: 3, Name: "William", CatchPhrase: "So many customers these days..", FoodChan: make(chan FoodToCook, 100), ProfficiencyChan: make(chan int, 3)},
	{Id: 3, Rank: 2, Proficiency: 2, Name: "Elizabeth", CatchPhrase: "Oh! I gotta hurry!", FoodChan: make(chan FoodToCook, 100), ProfficiencyChan: make(chan int, 2)},
	{Id: 4, Rank: 1, Proficiency: 2, Name: "Andrew", CatchPhrase: "Oh! That's my favourite meal!", FoodChan: make(chan FoodToCook, 100), ProfficiencyChan: make(chan int, 2)},
}
