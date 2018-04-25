package main

import (
	"testing"
)

func TestIntSet(t *testing.T) {
	var set intSet
	const targetValue = 7

	if set.has(targetValue) {
		t.Error("Has before being added")
	}

	alreadyThereBefore := set.add(targetValue)
	if alreadyThereBefore {
		t.Error("add should return false for first add")
	}

	if !set.has(targetValue) {
		t.Error("Doesn't have value after being added")
	}

	alreadyThereAfter := set.add(targetValue)
	if !alreadyThereAfter {
		t.Error("add should return true for second add")
	}
}
