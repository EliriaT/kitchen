package kitchen_elem

import (
	"time"
)

const (
	TimeUnit         = time.Duration(float64(time.Millisecond) * 10)
	ApparatusQuantum = 10
	//URL      = "http://dinning-hall:8082/distribution"
	URL = "http://localhost:8082/distribution"
)

type apparatusType string

const (
	stoveLit apparatusType = "oven"
	ovenLit  apparatusType = "stove"
)

var (
	Ovens  Apparatus
	Stoves Apparatus

	OrdersChannel = make(chan ReceivedOrd, 20)
	CookFree      = make(chan int, 11)
	Cooks         []cook
	Foods         []Food
)
