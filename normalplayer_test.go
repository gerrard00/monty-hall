package main

import "testing"

func TestPlayerSticks(t *testing.T) {
	const selection = 7
	pNormal := normalPlayer{}

	pNormal.saveChoice(selection)

	var normalPlayerSecondSelection = pNormal.choose(1, 10)

	if normalPlayerSecondSelection != selection {
		t.Error("Player changed selection")
	}
}
