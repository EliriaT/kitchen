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

var Cooks = []cook{
	{Id: 1, Rank: 3, Proficiency: 4, Name: "Mike", CatchPhrase: "I like ice-creams!", FoodChan: make(chan FoodToCook, 100), ProfficiencyChan: make(chan int, 4)},
	{Id: 2, Rank: 2, Proficiency: 3, Name: "William", CatchPhrase: "So many customers these days..", FoodChan: make(chan FoodToCook, 100), ProfficiencyChan: make(chan int, 3)},
	{Id: 3, Rank: 2, Proficiency: 2, Name: "Elizabeth", CatchPhrase: "Oh! I gotta hurry!", FoodChan: make(chan FoodToCook, 100), ProfficiencyChan: make(chan int, 2)},
	{Id: 4, Rank: 1, Proficiency: 2, Name: "Andrew", CatchPhrase: "Oh! That's my favourite meal!", FoodChan: make(chan FoodToCook, 100), ProfficiencyChan: make(chan int, 2)},
}
