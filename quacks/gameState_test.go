package quacks

import (
	"fmt"
	"testing"
)

func TestInitalGameState(t *testing.T) {

	playerNames := []string{"Nathan", "Leah", "Raymond", "Hannah"}

	gs := CreateGameState(playerNames, true)

	if len(GetTopPlayers(gs.players)) != len(playerNames) {
		t.Errorf("Players don't start with the same amount of points")
	}

	if gs.GameIsOver() == true {
		t.Errorf("Game is unexpectedly over")
	}

	gs.players[3].score = 1
	if gs.GetPlayersByScore()[0] != playerNames[3] {
		t.Errorf("GetPlayersByScore doesn't sort correctly")
	}
}

func TestHandlePurpleFortune2Rubies(t *testing.T) {
	playerNames := []string{"Nathan", "Leah", "Raymond", "Hannah"}

	gs := CreateGameState(playerNames, true)

	// Set up the fortune deck for the test
	gs.fortuneDeck = append(gs.fortuneDeck, Fortune{
		"Rubies",
		2,
	})

	fmt.Println(gs.fortuneDeck)

	gs.StartGame()

	if gs.FSM.Current() != FortuneInputState.String() {
		fmt.Printf("Starting State: %s\n", gs.FSM.Current())
		t.Errorf("Game does not wait for fortune input")
	}

	// Choice 1 is get a ruby
	addInputForAllPlayers(gs, 1)

	for _, player := range gs.players {
		if player.rubyCount != 1 {
			t.Errorf("Players aren't given rubies for fortune 2")
		}
	}
}

func TestHandlePurpleFortune2VP(t *testing.T) {
	playerNames := []string{"Nathan", "Leah", "Raymond", "Hannah"}

	gs := CreateGameState(playerNames, true)

	// Set up the fortune deck for the test
	gs.fortuneDeck = append(gs.fortuneDeck, Fortune{
		"Get a ruby or +1 VP",
		2,
	})

	fmt.Println(gs.fortuneDeck)

	gs.StartGame()

	if gs.FSM.Current() != FortuneInputState.String() {
		fmt.Printf("Starting State: %s\n", gs.FSM.Current())
		t.Errorf("Game does not wait for fortune input")
	}

	// Choice  is get a VP
	addInputForAllPlayers(gs, 2)

	for _, player := range gs.players {
		if player.score != 1 {
			t.Errorf("Players aren't given rubies for fortune 2")
		}
	}
}

func addInputForAllPlayers(gs *GameState, choice int) {
	for i := range gs.players {
		gs.Input(Input{Description: "", Choice: choice, Player: i})
		gs.ResumePlay()
	}
}
