package quacks

import (
	"fmt"
)

type Player struct {
	name           string
	bag            Bag
	board          Board
	rubyCount      int
	ratToken       int
	dropplet       int
	flask          bool
	explosionLimit int
	score          int
}

func PrintPlayerStatuses(player Player) {
	fmt.Println("Player:")
	fmt.Println(player)
}

func setUpPlayers(names []string) []Player {
	players := []Player{}
	for i := 0; i < len(names); i++ {
		players = append(players, Player{
			names[i],
			Bag{},
			Board{},
			0,
			0,
			0,
			true,
			7,
			0})
	}

	return players
}
