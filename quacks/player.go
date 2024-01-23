package quacks

import (
	"fmt"
)

type Player struct {
	name             string
	bag              Bag
	board            Board
	rubyCount        int
	ratToken         int
	dropplet         int
	testTubeDropplet int
	flask            bool
	explosionLimit   int
	score            int
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
			Bag{
				Chips: []Chip{
					NewChip("white", 1),
					NewChip("white", 1),
					NewChip("white", 1),
					NewChip("white", 1),
					NewChip("white", 2),
					NewChip("white", 2),
					NewChip("white", 3),
					NewChip("green", 1),
					NewChip("orange", 1),
				},
				RemainingChips: []Chip{
					NewChip("white", 1),
					NewChip("white", 1),
					NewChip("white", 1),
					NewChip("white", 1),
					NewChip("white", 2),
					NewChip("white", 2),
					NewChip("white", 3),
					NewChip("green", 1),
					NewChip("orange", 1),
				},
			},
			Board{},
			0,
			0,
			0,
			0,
			true,
			7,
			0})
	}

	return players
}
