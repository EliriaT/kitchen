package kitchen_elem

import "time"

// the order response sent back to dinning-hall
type SentOrd struct {
	OrderId        int              `json:"order_id"`
	TableId        int              `json:"table_id"`
	WaiterId       int              `json:"waiter_id"`
	Items          []int            `json:"items"`
	Priority       int              `json:"priority"`
	MaxWait        int              `json:"max_wait"`
	PickUpTime     time.Time        `json:"pick_up_time"`
	CookingTime    time.Duration    `json:"cooking_time"`
	CookingDetails []KitchenFoodInf `json:"cooking_details"`
}

// info of cooked food
type KitchenFoodInf struct {
	FoodId int `json:"food_id"`
	CookId int `json:"cook_id"`
}

// the response received from dinning hall
type ReceivedOrd struct {
	OrderId    int       `json:"order_id"`
	TableId    int       `json:"table_id"`
	WaiterId   int       `json:"waiter_id"`
	Items      []int     `json:"items"`
	Priority   int       `json:"priority"`
	MaxWait    int       `json:"max_wait"`
	PickUpTime time.Time `json:"pick_up_time"`
}

// a map of order ID with key, and the order as value
var OrderMap = make(map[int]ReceivedOrd)

// waiter's goroutine receive orders on channel
var OrdersChannel = make(chan int, 10)
