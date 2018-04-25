package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestPlayGames(t *testing.T) {
	normalp := &normalPlayer{}
	switchesp := &switchesPlayer{}
	g := game{
		numberOfDoors: 500,
		players: []player{
			normalp,
			switchesp,
		},
	}

	results := g.playGames(10000)
	spew.Dump(results)
}
