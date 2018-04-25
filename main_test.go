package main

import "testing"

func TestGetRandomValueInRangeHonorsRange(t *testing.T) {
	const min = 1
	const max = 10

	result := getRandomValueInRange(min, max, intSet{})

	if result < min || result > max {
		t.Errorf("Random value %d was out of range %d - %d\n",
			result, min, max)
	}
}

func TestGetRandomValueInRangeHonorsExclusion(t *testing.T) {
	const min = 1
	const max = 10
	const excludedValue = 7

	excludedSet := intSet{}
	excludedSet.add(excludedValue)

	result := getRandomValueInRange(min, max, excludedSet)

	if result == excludedValue {
		t.Errorf("Random value %d matched excluded value.\n", result)
	}
}

func TestGetRandomValueInRangeRandomness(t *testing.T) {
	const min = 1
	const max = 10
	theRange := max - min + 1
	valuesFound := intSet{}

	const numberOfTests = 1000

	for i := 0; i < numberOfTests; i++ {
		current := getRandomValueInRange(min, max, intSet{})
		valuesFound.add(current)
		if valuesFound.size() == theRange {
			// spew.Dump(valuesFound)
			return
		}
	}

	t.Errorf("Only got %d of %d possible values", valuesFound.size(), theRange)
}
