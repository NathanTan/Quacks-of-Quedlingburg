package quacks

import (
	"fmt"
	"testing"
)

func TestFortune4(t *testing.T) {
	playerNames := []string{"Nathan", "Leah", "Raymond", "Hannah"}

	gs := CreateGameState(playerNames, true)

	// Set up the fortune deck for the test
	gs.fortuneDeck = append(gs.fortuneDeck, Fortune{
		"Double the number of rat tails in this round.",
		4,
	})

	fmt.Println(gs.fortuneDeck)

	// Set a high score so rat tails happen
	gs.players[0].score = 10

	gs.StartGame()

	if gs.FSM.Current() != RatTailsState.String() {
		fmt.Printf("Starting State: %s\n", gs.FSM.Current())
		t.Errorf("Game does not wait for fortune input")
	}

	if gs.players[0].ratToken != 0 {
		t.Errorf("Player has the wrong amount of rat tails")
	}

	for i, player := range gs.players {
		expected := 6
		if i != 0 && player.ratToken != expected {
			t.Errorf("Player has the wrong amount of rat tails. Expected: %d, Actual: %d", expected, player.ratToken)
		}
	}
}

func TestFortune5(t *testing.T) {
	playerNames := []string{"Nathan", "Leah", "Raymond", "Hannah"}

	gs := CreateGameState(playerNames, true)

	// Set up the fortune deck for the test
	gs.fortuneDeck = append(gs.fortuneDeck, Fortune{
		"Choose: Take 4 victory points OR remove 1 white 1-chip from your bag",
		5,
	})

	fmt.Println(gs.fortuneDeck)

	// Set a high score so rat tails happen

	gs.StartGame()

	if gs.FSM.Current() != FortuneInputState.String() {
		fmt.Printf("Starting State: %s\n", gs.FSM.Current())
		t.Errorf("Game does not wait for fortune input")
	}

	gs.Input(Input{
		Description: "Choose: Take 4 victory points OR remove 1 white 1-chip from your bag" + ". (Choose 1 or 2)",
		Options:     []string{"1", "2"},
		Choice:      1,
		Choice2:     []Chip{}, // todo refactor for buying chips
		Player:      0,        // Player Position
		Code:        5,
	})

	gs.ResumePlay()

	if gs.players[0].score != 4 {
		t.Errorf("Input for player 0 ignored.")
	}

	for i, player := range gs.players {
		expected := 0
		if i != 0 && player.score != expected {
			t.Errorf("Player has the wrong amount of VP. Expected: %d, Actual: %d", expected, player.ratToken)
		}
	}

	if len(gs.GetRemainingFortunePlayers()) != 3 {
		t.Errorf("Game has the wrong amount of players that have done their fortune inputs")
	}

	for i := 1; i < len(gs.players)-1; i++ {
		gs.Input(Input{
			Description: "Choose: Take 4 victory points OR remove 1 white 1-chip from your bag" + ". (Choose 1 or 2)",
			Options:     []string{"1", "2"},
			Choice:      1,
			Choice2:     []Chip{}, // todo refactor for buying chips
			Player:      i,        // Player Position
			Code:        5,
		})

		gs.ResumePlay()
	}

	if gs.players[1].score != 4 && gs.players[2].score != 4 {
		t.Errorf("Input for player 1 or 2 ignored.")
	}

	if len(gs.GetRemainingFortunePlayers()) != 1 {
		t.Errorf("Game has the wrong amount of players that have done their fortune inputs")
	}

	gs.Input(Input{
		Description: "Choose: Take 4 victory points OR remove 1 white 1-chip from your bag" + ". (Choose 1 or 2)",
		Options:     []string{"1", "2"},
		Choice:      2,
		Choice2:     []Chip{}, // todo refactor for buying chips
		Player:      3,        // Player Position
		Code:        5,
	})

	gs.ResumePlay()

	if len(gs.players[3].bag.Chips) != 8 {
		t.Errorf("Game has the wrong amount of chips for players 3: %d", len(gs.players[3].bag.Chips))
	}
}
