package kitchen_elem

import (
	"time"
)

type Apparatus struct {
	Name     apparatusType `json:"name"`
	Quantity int           `json:"quantity"`
	//New      chan FoodToCook
	Accepted chan FoodToCook
}

func (a Apparatus) CookFood() {

	for food := range a.Accepted {
		a.cookQuantumTime(food)
	}

}

func (a Apparatus) cookQuantumTime(food FoodToCook) {
	if food.PrepTime <= ApparatusQuantum {
		time.Sleep(TimeUnit * time.Duration(food.PrepTime))
		NrFoodsQueue--
		food.Wg.Done()

	} else {

		food.PrepTime = food.PrepTime - ApparatusQuantum

		time.Sleep(TimeUnit * ApparatusQuantum)
		a.Accepted <- food

	}

}

type apparatuses struct {
	ApparatusList []Apparatus `json:"apparatuses"`
}
