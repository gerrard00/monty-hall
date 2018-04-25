package main

type player interface {
	choose(min int, max int) int
	getName() string
	saveChoice(choice int)
	notifyEmpty(choice int)
	reset()
}
