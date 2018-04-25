package main

import "testing"

func TestSwitchesPlayer(t *testing.T) {
	pSwitches := switchesPlayer{}

	pSwitches.saveChoice(7)

	var switchesPlayerSecondSelection = pSwitches.choose(1, 10)

	// fmt.Printf("Switches: %d\n", switchesPlayerSecondSelection)

	if switchesPlayerSecondSelection == 7 {
		t.Error("Switch Player didn't switch selection")
	}
}
