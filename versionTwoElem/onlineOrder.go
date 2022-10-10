package versionTwoElem

import "time"

type OnlineReceivedOrder struct {
	Id          int       `json:"id"`
	Items       []int     `json:"items"`
	Priority    int       `json:"priority"`
	MaxWait     float32   `json:"max_wait"`
	CreatedTime time.Time `json:"created_time"`
}
