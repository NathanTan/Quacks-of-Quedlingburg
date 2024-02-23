package quacks

import (
	// q "quacks"
	"fmt"
	"testing"
)

func TestRatTailCount(t *testing.T) {

	res := countRatTails(1, 10)

	if res != 3 {
		t.Errorf("Incorrect rat tail count: %d", res)
	}

	res2 := countRatTails(2, 10)

	if res2 != 2 {
		t.Errorf("Incorrect rat tail count: %d", res)
	}

	res3 := countRatTails(100, 10)

	if res3 != 0 {
		t.Errorf("Incorrect rat tail count: %d", res)
	}
}

func TestAddition(t *testing.T) {

	bag := Bag{
		Chips: []Chip{
			NewChip("blue", 4),
			NewChip("green", 1),
		},
		RemainingChips: []Chip{
			NewChip("blue", 4),
			NewChip("green", 1),
		},
	}

	if len(bag.RemainingChips) != 2 {
		t.Errorf("Error")
	}

	chip := DrawChip(&bag, true)

	if len(bag.RemainingChips) != 1 {
		t.Errorf("Error 2")
	}
	if len(bag.Chips) != 2 {
		t.Errorf("Error 3")
	}
	if chip != NewChip("green", 1) {
		t.Errorf("Error 4")
	}
}

func TestRollDiceForPlayer(t *testing.T) {

	player := CreateNewPlayer("test")

	player.RollBonusDice(true)

	if player.rubyCount == 0 &&
		player.score == 0 &&
		player.dropplet == 0 &&
		len(player.board.chips) == 9 {
		t.Errorf("Bonus die had no impact")
	}
}

func TestStartGame(t *testing.T) {
	playerNames := []string{"Nathan", "Leah", "Raymond", "Hannah"}

	gs := CreateGameState(playerNames, true)

	if gs.FSM.Current() != ClosedState.String() {
		t.Errorf("Game does not begin with Closed State")
	}

	gs.StartGame()

	if gs.FSM.Current() != AssignRatTails.String() {
		fmt.Printf("Starting State: %s\n", gs.FSM.Current())
		t.Errorf("Game does not begin with Rat tails")
	}

}

func TestStartGameWithInputFortune(t *testing.T) {
	playerNames := []string{"Nathan", "Leah", "Raymond", "Hannah"}

	gs := CreateGameState(playerNames, true)

	// Set up the fortune deck for the test

	gs.fortuneDeck = createFortunes()

	gs.fortuneDeck = append(gs.fortuneDeck, Fortune{
		"Everyone rolls the die once and gets the bonus shown",
		3,
	})

	fmt.Println(gs.fortuneDeck)

	gs.StartGame()

	if len(gs.fortuneDeck) < 4 {
		t.Errorf("Not enough cards")
	}

	if gs.FSM.Current() != FortuneInputState.String() {
		fmt.Printf("Starting State: %s\n", gs.FSM.Current())
		t.Errorf("Game does not wait for fortune input")
	}
}
