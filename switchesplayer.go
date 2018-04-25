package main

type switchesPlayer struct {
	normalPlayer
	name string
}

func (sp switchesPlayer) getName() string {
	return "Switches"
}

func (sp *switchesPlayer) choose(min int, max int) int {
	result := getRandomValueInRange(min, max, sp.exclusionSet)
	sp.normalPlayer.saveChoice(result)
	return result
}
