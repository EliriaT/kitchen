package kitchen_elem

import (
	"time"
)

type cook struct {
	Id               int    `json:"id"`
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

		if food.CookingApparatus == ovenLit {

			<-c.ProfficiencyChan
			<-CookFree
			Ovens.Accepted <- food

			//TODO trebuie conditional variable

		} else if food.CookingApparatus == stoveLit {

			<-c.ProfficiencyChan
			<-CookFree

			Stoves.Accepted <- food

		} else {
			go c.cookFood(food)
		}

		//log.Printf("Cook %d cooks food %d", c.Id, food.FoodId)

	}
}

// i send it here by value
func (c *cook) cookFood(food FoodToCook) {
	time.Sleep(TimeUnit * time.Duration(Foods[food.FoodId-1].PreparationTime))
	food.Wg.Done()
	NrFoodsQueue--
	<-c.ProfficiencyChan
	<-CookFree
}
