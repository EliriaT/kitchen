package kitchen_elem

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)

type Apparatus struct {
	Name      apparatusType `json:"name"`
	Quantity  int           `json:"quantity"`
	QuantChan chan int
}

func (a Apparatus) cookFood(food FoodToCook) {
	a.QuantChan <- 1
	time.Sleep(TimeUnit * time.Duration(Foods[food.FoodId-1].PreparationTime))
	<-a.QuantChan
	food.Wg.Done()

}

type apparatuses struct {
	ApparatusList []Apparatus `json:"apparatuses"`
}

func InitiateApparatus() {

	file, err := os.Open("./jsonConfig/apparatus.json")
	if err != nil {
		log.Fatal("Error opening apparatus.json", err)
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	var apparList apparatuses
	err = json.Unmarshal(byteValue, &apparList)
	if err != nil {
		log.Fatal("Error unmarshling apparatus.json", err)
	}

	for i := range apparList.ApparatusList {

		if apparList.ApparatusList[i].Name == stoveLit {
			Stoves = apparList.ApparatusList[i]
			Stoves.QuantChan = make(chan int, apparList.ApparatusList[i].Quantity)

		} else if apparList.ApparatusList[i].Name == ovenLit {
			Ovens = apparList.ApparatusList[i]
			Ovens.QuantChan = make(chan int, apparList.ApparatusList[i].Quantity)

		}
	}
}
