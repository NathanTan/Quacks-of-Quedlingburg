package quacks

import (
	"fmt"
)

type Player struct {
	Name                   string
	bag                    Bag
	Board                  Board
	isDoneDrawing          bool
	hasCompletedTheFortune bool
	hasSpentRubies         bool
	rubyCount              int
	ratToken               int
	dropplet               int
	testTubeDropplet       int
	flask                  bool
	explosionLimit         int
	score                  int
	chooseVictoryPoints    bool
	chooseBuying           bool
	buyingPower            int
}

func (player Player) IsDoneDrawing(bombLimit int) bool {
	if player.isDoneDrawing {
		return true
	}

	fmt.Printf("Player name: %s, CherryBombValue: %d, Limit: %d\n", player.Name, player.Board.CherryBombValue, bombLimit)

	if player.Board.CherryBombValue > bombLimit {
		return true
	}

	if len(player.bag.Chips) == 0 {
		return true
	}

	return false
}

func (p *Player) moveDropper(amount int, choice int) {
	if choice == 1 {
		p.dropplet = p.dropplet + amount
	} else {
		fmt.Println("TODO")
	}
}

func (p Player) PlayersPotHasExploaded() bool {
	return p.Board.CherryBombValue > p.explosionLimit
}

func PrintPlayerStatuses(player Player) {
	fmt.Printf("Player '%s':\n", player.Name)
	fmt.Printf("\tScore - %d\n", player.score)
	fmt.Printf("\tCherry Bomb Count - %d\n", player.Board.CherryBombValue)
	fmt.Printf("\tRuby Count - %d\n", player.rubyCount)
	fmt.Printf("\tOwned Chips - %s\n\n", ChipsString(player.bag.Chips))
	fmt.Printf("\tDropper Position - %d\n", player.dropplet)

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
			Chips:            nil,
			NextPosition:     -1,
			TestTubePosition: 0,
			CherryBombValue:  0,
		},
		false,
		false,
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
