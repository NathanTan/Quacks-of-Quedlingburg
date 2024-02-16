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
	buyingPower         int
}

func (player Player) IsDoneDrawing(bombLimit int) bool {
	if player.isDoneDrawing {
		return true
	}

	fmt.Printf("Player name: %s, CherryBombValue: %d, Limit: %d\n", player.name, player.board.cherryBombValue, bombLimit)

	if player.board.cherryBombValue > bombLimit {
		return true
	}

	if len(player.bag.Chips) == 0 {
		return true
	}

	return false
}

func (p Player) moveDropper(amount int, choice int) {
	if choice == 1 {
		p.dropplet = p.dropplet + amount
	} else {
		fmt.Println("TODO")
	}
}

func PrintPlayerStatuses(player Player) {
	fmt.Printf("Player '%s':\n", player.name)
	fmt.Printf("\tScore - %d\n", player.score)
	fmt.Printf("\tCherry Bomb Count - %d\n", player.board.cherryBombValue)
	fmt.Printf("\tRuby Count - %d\n", player.rubyCount)
	fmt.Printf("\tOwned Chips - %s\n\n", ChipsString(player.bag.Chips))

	fmt.Println(player)
}

func setUpPlayers(names []string) []Player {
	players := []Player{}
	for i := 0; i < len(names); i++ {
		players = append(players, CreateNewPlayer(names[i]))
	}

	return players
}

func CreateNewPlayer(name string) Player {
	return Player{
		name,
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
		Board{
			chips:            nil,
			nextPosition:     -1,
			testTubePosition: 0,
			cherryBombValue:  0,
		},
		false,
		0,
		0,
		0,
		0,
		true,
		7,
		0,
		false,
		false,
		0,
	}
}
