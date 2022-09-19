package kitchen_elem

import (
	"sync"
)

type Food struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	PreparationTime  int    `json:"preparation-time"`
	Complexity       int    `json:"complexity"`
	CookingApparatus string `json:"cooking-apparatus"`
}

// info of cooked food
type KitchenFoodInf struct {
	FoodId int `json:"food_id"`
	CookId int `json:"cook_id"`
}

type FoodToCook struct {
	OrderId int
	//used to find the time for cooking in the foods list; rather should be named foodsMenu
	FoodId int
	CookId int
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
