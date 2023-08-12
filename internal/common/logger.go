package common

import (
	"strconv"
	"time"
)

var id int = 999

func GetTimeNow() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

func GetID() string {
	id++
	return "ST" + strconv.Itoa(id)
}
