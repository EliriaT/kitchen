package kitchen_elem

import (
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
		CookFree <- 1
		//this does not matter, I can send any number, used to control profficiency
		c.ProfficiencyChan <- food.FoodId
		//it is sent by value
		//log.Printf("Cook %d cooks food %d", c.Id, food.FoodId)

		if food.CookingApparatus == ovenLit {

			<-c.ProfficiencyChan
			<-CookFree
			//the cooks sends food to apparatus and the proceeds
			go Ovens.cookFood(food)

		} else if food.CookingApparatus == stoveLit {

			<-c.ProfficiencyChan
			<-CookFree
			//the cooks sends food to apparatus and the proceeds
			go Stoves.cookFood(food)

		} else {
			go c.cookFood(food)
		}

		//here is cheating?
		//<-CookFree
	}
}

// i send it here by value
func (c *cook) cookFood(food FoodToCook) {
	time.Sleep(TimeUnit * time.Duration(Foods[food.FoodId-1].PreparationTime))
	food.Wg.Done()
	<-c.ProfficiencyChan
	<-CookFree

}

// to make a json
var Cooks = []cook{
	{Id: 1, Rank: 3, Proficiency: 4, Name: "Mike", CatchPhrase: "I like ice-creams!", FoodChan: make(chan FoodToCook, 100), ProfficiencyChan: make(chan int, 4)},
	{Id: 2, Rank: 2, Proficiency: 3, Name: "William", CatchPhrase: "So many customers these days..", FoodChan: make(chan FoodToCook, 100), ProfficiencyChan: make(chan int, 3)},
	{Id: 3, Rank: 2, Proficiency: 2, Name: "Elizabeth", CatchPhrase: "Oh! I gotta hurry!", FoodChan: make(chan FoodToCook, 100), ProfficiencyChan: make(chan int, 2)},
	{Id: 4, Rank: 1, Proficiency: 2, Name: "Andrew", CatchPhrase: "Oh! That's my favourite meal!", FoodChan: make(chan FoodToCook, 100), ProfficiencyChan: make(chan int, 2)},
}
