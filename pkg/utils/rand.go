package utils

import (
	"math/rand"
	"time"
)

func RandA2B(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}

func RandPercent(p int) bool {
	if p < 1 || p > 100 {
		return true
	}

	n := RandA2B(1, 100)

	return n <= p
}
