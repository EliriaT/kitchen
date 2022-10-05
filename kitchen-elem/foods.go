package kitchen_elem

import (
	"sync"
)

type Food struct {
	Id               int           `json:"id"`
	Name             string        `json:"name"`
	PreparationTime  int           `json:"preparation-time"`
	Complexity       int           `json:"complexity"`
	CookingApparatus apparatusType `json:"cooking-apparatus"`
}

// info of cooked food sent to dinning-hall in response
type KitchenFoodInf struct {
	FoodId int `json:"food_id"`
	CookId int `json:"cook_id"`
}

type FoodToCook struct {
	OrderId          int
	FoodId           int
	CookId           int
	PrepTime         int
	CookingApparatus apparatusType
	Wg               *sync.WaitGroup
}
