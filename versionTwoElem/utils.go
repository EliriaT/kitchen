package versionTwoElem

import kitchen_elem "github.com/EliriaT/kitchen/kitchen-elem"

func GetCurrentKitchenInfo() KitchenInfo {
	var kitchenInf KitchenInfo
	kitchenInf.CookingApparatus = kitchen_elem.CookingApparatus
	kitchenInf.CooksProfficiency = kitchen_elem.CooksProffieciency
	kitchenInf.NrFoodsQueue = kitchen_elem.NrFoodsQueue
	return kitchenInf
}
