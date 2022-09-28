package kitchen_elem

import (
	"time"
)

const (
	TimeUnit = time.Duration(float64(time.Millisecond) * 10)
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
	//used for priority
	CookFree = make(chan int, 11)
)
