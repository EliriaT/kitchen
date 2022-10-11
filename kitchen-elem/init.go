package kitchen_elem

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func StartWorkDay() {
	initiate_Congif()
	initiate_Cooks()
	initiateApparatus()
	initiate_Foods()

}

func initiate_Congif() {
	var config map[string]string

	file, err := os.Open("./jsonConfig/config.json")
	if err != nil {
		log.Fatal("Error opening config.json ", err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	_ = json.Unmarshal(byteValue, &config)
	URL = config["address"]
	Port = config["listenning_port"]

}

func initiate_Cooks() {
	file, err := os.Open("./jsonConfig/cooks.json")
	if err != nil {
		log.Fatal("Error opening cooks.json", err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	_ = json.Unmarshal(byteValue, &Cooks)

	for i := range Cooks {
		Cooks[i].ProfficiencyChan = make(chan int, Cooks[i].Proficiency)
		Cooks[i].FoodChan = make(chan FoodToCook, 100)
	}

}

func initiate_Foods() {
	file, err := os.Open("./jsonConfig/foods.json")
	if err != nil {
		log.Fatal("Error opening foods.json", err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	_ = json.Unmarshal(byteValue, &Foods)

}

func initiateApparatus() {

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

	var stove, oven Apparatus

	for i := range apparList.ApparatusList {

		if apparList.ApparatusList[i].Name == stoveLit {

			stove = apparList.ApparatusList[i]
			stove.Accepted = make(chan FoodToCook, 100)

			Stoves = stove

		} else if apparList.ApparatusList[i].Name == ovenLit {

			oven = apparList.ApparatusList[i]
			oven.Accepted = make(chan FoodToCook, 100)

			Ovens = oven

		}
	}
}
