package main

import "fmt"

type gameResult struct {
	totalRounds           int
	firstSelectionResults int
	playerResults         map[string]int
}

func (g *game) playGames(numberOfGames int) gameResult {
	results := gameResult{
		totalRounds:           0,
		firstSelectionResults: 0,
		playerResults:         make(map[string]int),
	}
	const min = 1
	max := g.numberOfDoors

	for gameIndex := 0; gameIndex < numberOfGames; gameIndex++ {
		selectedDoor := getRandomValueInRange(min, max, intSet{})
		winningDoor := getRandomValueInRange(min, max, intSet{})

		gameExclusionSet := intSet{}
		gameExclusionSet.add(selectedDoor)
		gameExclusionSet.add(winningDoor)

		// If we chose correctly the exclusion set will only have 1 item because
		// our selection is the winner. If so, we need to pick another door to keep
		// closed.
		if gameExclusionSet.size() == 1 {
			doorToNotOpen := getRandomValueInRange(min, max, gameExclusionSet)
			gameExclusionSet.add(doorToNotOpen)
		}

		results.totalRounds++
		if selectedDoor == winningDoor {
			results.firstSelectionResults++
		}

		fmt.Printf("Game %d Winning door: %d First Selection: %d \n",
			gameIndex, winningDoor, selectedDoor)

		// all players get the same first selection
		for _, v := range g.players {
			v.saveChoice(selectedDoor)
		}

		for doorToOpen := 1; doorToOpen <= max; doorToOpen++ {
			if gameExclusionSet.has(doorToOpen) {
				continue
			}

			// fmt.Printf("\t\topen %v\n", doorToOpen)
			for _, v := range g.players {
				// show each player the empty door
				v.notifyEmpty(doorToOpen)
			}
		}

		// now let the players make their second selection
		for _, v := range g.players {
			currentName := v.getName()
			playerChoice := v.choose(min, max)
			isCorrect := (playerChoice == winningDoor)
			fmt.Printf("\t%v chose %d: %v\n", currentName, playerChoice, isCorrect)
			if isCorrect {
				results.playerResults[currentName]++
			}
			v.reset()
		}

	}

	return results
}

type game struct {
	numberOfDoors int
	players       []player
}
