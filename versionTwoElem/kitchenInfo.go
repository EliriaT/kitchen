package versionTwoElem

import kitchen_elem "github.com/EliriaT/kitchen/kitchen-elem"

type KitchenInfo struct {
	CookingApparatus  int `json:"cookingApparatus"`
	CooksProfficiency int `json:"cooksProfficiency"`
	NrFoodsQueue      int `json:"nrFoodsQueue"`
}

func GetCurrentKitchenInfo() KitchenInfo {
	var kitchenInf KitchenInfo
	kitchenInf.CookingApparatus = kitchen_elem.CookingApparatus
	kitchenInf.CooksProfficiency = kitchen_elem.CooksProffieciency
	kitchenInf.NrFoodsQueue = kitchen_elem.NrFoodsQueue
	return kitchenInf
}
