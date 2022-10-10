package handlers

import (
	"encoding/json"
	kitchen_elem "github.com/EliriaT/kitchen/kitchen-elem"
	"github.com/EliriaT/kitchen/versionTwoElem"
	"log"
	"net/http"
)

// here i should receive the online order and send back Kitchen current info
func ReceiveOnlineOrder(w http.ResponseWriter, r *http.Request) {
	var onlineReceivedOrder versionTwoElem.OnlineReceivedOrder

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&onlineReceivedOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//sending in the general orderds channel the received online order from dinning-hall
	var sentOrderToCook kitchen_elem.ReceivedOrd

	sentOrderToCook.OrderId = onlineReceivedOrder.Id
	sentOrderToCook.TableId = -1
	sentOrderToCook.WaiterId = -1
	sentOrderToCook.Items = onlineReceivedOrder.Items
	sentOrderToCook.Priority = onlineReceivedOrder.Priority
	sentOrderToCook.MaxWait = float64(onlineReceivedOrder.MaxWait)
	sentOrderToCook.PickUpTime = onlineReceivedOrder.CreatedTime

	kitchen_elem.OrdersChannel <- sentOrderToCook

	kitchenCurrentInfo := versionTwoElem.GetCurrentKitchenInfo()
	jsonKitchenCurrentInfo, _ := json.Marshal(kitchenCurrentInfo)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonKitchenCurrentInfo)

	log.Printf("Online Order, arrived at kitchen %v \n", onlineReceivedOrder)

}
