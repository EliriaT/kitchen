package kitchen_elem

import (
	"time"
)

const (
	TimeUnit         = time.Duration(float64(time.Millisecond) * 25)
	ApparatusQuantum = 10

	CookingApparatus   = 3
	CooksProffieciency = 11
)

type apparatusType string

const (
	stoveLit apparatusType = "oven"
	ovenLit  apparatusType = "stove"
)

var (
	//URL              = "http://dinning-hall:8082/distribution"
	URL  = "http://localhost:8082/distribution"
	Port string

	Ovens        Apparatus
	Stoves       Apparatus
	NrFoodsQueue = 0

	OrdersChannel = make(chan ReceivedOrd, 20)
	CookFree      = make(chan int, 11)
	Cooks         []cook
	Foods         []Food
)
