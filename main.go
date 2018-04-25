package main

import (
	"math/rand"
	"time"
)

// TODO: exclude needs to be a slice
func getRandomValueInRange(min int, max int, exclude intSet) int {
	maxAttempts := (max - min) * 100000
	rand.Seed(time.Now().UnixNano())
	// +1 because our ranges are inclusive
	result := rand.Intn(max-min+1) + min

	for attempts := 0; exclude.has(result) && attempts < maxAttempts; attempts++ {
		result = rand.Intn(max-min+1) + min
	}

	return result
}

func getRangeSet(min int, max int) intSet {
	result := intSet{}
	for i := 0; i <= max; i++ {
		result.add(i)
	}

	return result
}

func main() {
}
