package main

type normalPlayer struct {
	hasSelected      bool
	currentSelection int
	exclusionSet     intSet
}

func (p *normalPlayer) getName() string {
	return "Normal"
}

func (p *normalPlayer) reset() {
	p.hasSelected = false
	p.currentSelection = 0
	p.exclusionSet = intSet{}
}

func (p *normalPlayer) saveChoice(choice int) {
	// fmt.Printf("Save %d\n", choice)
	p.hasSelected = true
	p.currentSelection = choice
	p.exclusionSet.add(choice)
}

func (p *normalPlayer) notifyEmpty(choice int) {
	p.exclusionSet.add(choice)
}

func (p *normalPlayer) choose(min int, max int) int {
	var result int
	// fmt.Printf("%v\t%v\n", p.hasSelected, p.currentSelection)
	if p.hasSelected {
		result = p.currentSelection
	} else {
		result = getRandomValueInRange(min, max, intSet{})
	}

	p.saveChoice(result)

	return result
}
