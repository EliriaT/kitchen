package handlers

import (
	"encoding/json"
	kitchen_elem "github.com/EliriaT/kitchen/kitchen-elem"
	"log"
	"net/http"
)

func GetCooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonFoods, err := json.Marshal(kitchen_elem.Cooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//by default sends 200
	w.Write(jsonFoods)
}

func ReceiveOrder(w http.ResponseWriter, r *http.Request) {
	var unCookedOrder kitchen_elem.ReceivedOrd
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&unCookedOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//sending in the general orderds channel the received order from dinning-hall
	kitchen_elem.OrdersChannel <- unCookedOrder

	jsonUncookedOrder, _ := json.Marshal(unCookedOrder)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUncookedOrder)

	log.Printf("Order with ID %d, arrived at kitchen", unCookedOrder.OrderId)

}
