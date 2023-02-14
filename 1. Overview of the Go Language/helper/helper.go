package helper

import (
	"math/rand"
	"time"
)

func RandomNumber(maxNumber int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(maxNumber)
	return value
}