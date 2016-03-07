package util

import (
	"math/rand"
	"time"
)

// Generate a random int between min and max, inclusive
func Random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

// Pick a random element in the slice of ints
func RandomInt(elements []int) int {
	index := Random(0, len(elements))
	return elements[index]
}
