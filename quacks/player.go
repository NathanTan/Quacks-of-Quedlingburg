package quacks

import (
	"fmt"
)

type Player struct {
	name                string
	bag                 Bag
	board               Board
	isDoneDrawing       bool
	rubyCount           int
	ratToken            int
	dropplet            int
	testTubeDropplet    int
	flask               bool
	explosionLimit      int
	score               int
	chooseVictoryPoints bool
	chooseBuying        bool
}

func (player Player) IsDoneDrawing(bombLimit int) bool {
	if player.isDoneDrawing {
		return true
	}

	if player.board.cherryBombValue > bombLimit {
		return true
	}

	if len(player.bag.Chips) == 0 {
		return true
	}

	return false
}

func PrintPlayerStatuses(player Player) {
	fmt.Printf("Player '%s':\n", player.name)
	fmt.Printf("\tScore - %d\n", player.score)
	fmt.Printf("\tCherry Bomb Count - %d\n", player.board.cherryBombValue)
	fmt.Printf("\tRuby Count - %d\n", player.rubyCount)
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
			false,
			0,
			0,
			0,
			0,
			true,
			7,
			0,
			false,
			false})
	}

	return players
}
