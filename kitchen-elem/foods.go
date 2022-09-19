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

type Food struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	PreparationTime  int    `json:"preparation-time"`
	Complexity       int    `json:"complexity"`
	CookingApparatus string `json:"cooking-apparatus"`
}

type OrderInKitchen struct {
	Id    int
	Foods []FoodToCook
	Wg    *sync.WaitGroup
}

func (o *OrderInKitchen) WaitForOrder(initialOrder ReceivedOrd) {
	cookingTime := time.Now()
	o.Wg.Wait()
	var cookedOrder SentOrd

	//SA FAC CA SA TRANSMIT COOKID
	var foodCookedInfo = make([]KitchenFoodInf, 0, len(initialOrder.Items))

	for _, foodID := range initialOrder.Items {
		foodCookedInfo = append(foodCookedInfo, KitchenFoodInf{
			FoodId: foodID,
			CookId: -1,
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
	//body, err := io.ReadAll(resp.Body) // Log the request body
	//if err != nil {
	//	log.Printf("Can't read the response body %s", err.Error())
	//	return
	//}
	//bodyString := string(body)
	//log.Print(bodyString)
	log.Printf("The order with id %d was sent to Dinning Hall .", cookedOrder.OrderId) // Unmarshal result

}

type FoodToCook struct {
	OrderId int
	//used to find the time for cooking in the foods list; rather should be named foodsMenu
	FoodId int
	Wg     *sync.WaitGroup
}

var Foods = []Food{
	{Id: 1, Name: "pizza", PreparationTime: 20, Complexity: 2, CookingApparatus: "oven"},
	{Id: 2, Name: "salad", PreparationTime: 10, Complexity: 1, CookingApparatus: ""},
	{Id: 3, Name: "zeama", PreparationTime: 7, Complexity: 1, CookingApparatus: "stove"},
	{Id: 4, Name: "Scallop Sashimi with Meyer Lemon Confit", PreparationTime: 32, Complexity: 3, CookingApparatus: ""},
	{Id: 5, Name: "Island Duck with Mulberry Mustard", PreparationTime: 35, Complexity: 3, CookingApparatus: "oven"},
	{Id: 6, Name: "Waffles", PreparationTime: 10, Complexity: 1, CookingApparatus: "stove"},
	{Id: 7, Name: "Aubergine", PreparationTime: 20, Complexity: 2, CookingApparatus: "oven"},
	{Id: 8, Name: "Lasagna", PreparationTime: 30, Complexity: 2, CookingApparatus: "oven"},
	{Id: 9, Name: "Burger", PreparationTime: 15, Complexity: 1, CookingApparatus: "stove"},
	{Id: 10, Name: "Gyros", PreparationTime: 15, Complexity: 1, CookingApparatus: ""},
	{Id: 11, Name: "Kebab", PreparationTime: 15, Complexity: 1, CookingApparatus: ""},
	{Id: 12, Name: "Unagi Maki", PreparationTime: 20, Complexity: 2, CookingApparatus: ""},
	{Id: 13, Name: "Tobacco Chicken", PreparationTime: 30, Complexity: 2, CookingApparatus: "oven"},
}
