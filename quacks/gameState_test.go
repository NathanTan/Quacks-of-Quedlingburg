package quacks

import (
	"context"
	"fmt"
	"testing"
)

func TestInitalGameState(t *testing.T) {

	playerNames := []string{"Nathan", "Leah", "Raymond", "Hannah"}

	gs := CreateGameState(playerNames, "game123", true)

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

	gs := CreateGameState(playerNames, "game123", true)

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

	gs := CreateGameState(playerNames, "game123", true)

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

// func TestHandlePurpleFortune2VP(t *testing.T) {
func TestSpendingRubies(t *testing.T) {
	playerNames := []string{"Nathan", "Leah", "Raymond", "Hannah"}

	gs := CreateGameState(playerNames, "game123", true)

	// Given all the players Rubies to spend
	for i := range gs.players {
		gs.players[i].rubyCount = 2
	}

	// Set up the fortune deck for the test
	gs.fortuneDeck = append(gs.fortuneDeck, Fortune{
		"Test",
		0,
	})

	gs.StartGame()

	forwardState(gs, RubySpendingState.String())

	fmt.Println(gs.FSM.Current())

	// Given all the players Rubies to spend
	for _, player := range gs.players {
		player.rubyCount = 2
	}

	gs.ResumePlay()

	for i := 0; i < len(playerNames); i++ {

		remainingPlayers := gs.GetRemainingRubySpendingPlayerNames()
		if len(remainingPlayers) == 0 {
			gs.EndRubyBuys()
			fmt.Println("Done buying rubies")
		} else if gs.Awaiting != nil {

			fmt.Printf("remaing Players: %s\n", remainingPlayers)
			fmt.Printf("Awaiting on Player '%d' to '%s'\n", gs.Awaiting.Player, gs.Awaiting.Description)
			// for _, playerName := range remainingPlayers {
			gs.Input(Input{Description: "", Choice: 1, Player: gs.GetPlayerPositionByName(remainingPlayers[0])})
			gs.ResumePlay()
			gs.Input(Input{Description: "", Choice: -1, Player: gs.GetPlayerPositionByName(remainingPlayers[0])})
			gs.ResumePlay()
		}
	}

	// Check that every other player spent their rubies
	for i, player := range gs.players {
		if i%2 == 0 && player.rubyCount != 0 {
			t.Errorf("Player has unspent rubies")
		}
		if i%2 == 0 && player.dropplet != 1 {
			t.Errorf("Player didn't increase their dropplet value")
		}
		if i%2 == 1 && player.rubyCount != 2 {
			t.Errorf("Player has spent rubies")
		}
		if i%2 == 1 && player.dropplet > 0 {
			t.Errorf("Player %d increased their dropplet value", i)
		}
	}

	// TODO: Pick up here
	// if gs.FSM.Current() != FortuneInputState.String() {
	// 	fmt.Printf("Starting State: %s\n", gs.FSM.Current())
	// 	t.Errorf("Game does not wait for fortune input")
	// }

	// Choice  is get a VP
	// addInputForAllPlayers(gs, 2)

	// for _, player := range gs.players {
	// 	if player.score != 1 {
	// 		t.Errorf("Players aren't given rubies for fortune 2")
	// 	}
	// }
}

// Use case is when there is at least 1 player left that can spend their
// rubies but we are awaiting on the wrong player.
func TestSpendRubiesOnInvalidPlayer(t *testing.T) {
	playerNames := []string{"Nathan", "Leah", "Raymond", "Hannah"}

	gs := CreateGameState(playerNames "game123",, true)

	// Given all the players Rubies to spend
	for i := range gs.players {
		gs.players[i].rubyCount = 2
	}

	// Set up the fortune deck for the test
	gs.fortuneDeck = append(gs.fortuneDeck, Fortune{
		"Test",
		0,
	})

	gs.StartGame()

	forwardState(gs, RubySpendingState.String())

	fmt.Println(gs.FSM.Current())

	// Given all the players Rubies to spend
	for i := range gs.players {
		gs.players[i].rubyCount = 1
	}

	// Player 3 gets to spend rubies
	gs.players[2].rubyCount = 2

	gs.ResumePlay()

	for i := 0; i < len(playerNames); i++ {

		remainingPlayers := gs.GetRemainingRubySpendingPlayerNames()
		if len(remainingPlayers) == 0 {
			gs.EndRubyBuys()
			fmt.Println("Done buying rubies")
		} else if gs.Awaiting != nil {

			fmt.Printf("remaing Players: %s\n", remainingPlayers)
			fmt.Printf("Awaiting on Player '%d' to '%s'\n", gs.Awaiting.Player, gs.Awaiting.Description)
			// for _, playerName := range remainingPlayers {
			gs.Input(Input{Description: "", Choice: 1, Player: gs.GetPlayerPositionByName(remainingPlayers[0])})
			gs.ResumePlay()
			gs.Input(Input{Description: "", Choice: 1, Player: gs.GetPlayerPositionByName(remainingPlayers[0])})
			gs.ResumePlay()
		}
	}

	// Check that every other player spent their rubies
	for i, player := range gs.players {
		if i != 2 && player.dropplet > 0 {
			t.Errorf("Player increased their dropplet value")
		}
		if i == 2 && player.rubyCount == 2 {
			t.Errorf("Player has unspent rubies")
		}
	}
}

func addInputForAllPlayers(gs *GameState, choice int) {
	for i := range gs.players {
		gs.Input(Input{Description: "", Choice: choice, Player: i})
		gs.ResumePlay()
	}
}

func forwardState(gs *GameState, stateDestination string) {
	fmt.Printf("Forwarding State. Current State: %s\n", gs.FSM.Current())

	gs.FSM.Event(context.Background(), ReadFortune.String())
	fmt.Printf("Current State1: %s\n", gs.FSM.Current())

	gs.FSM.Event(context.Background(), AssignRatTails.String())
	fmt.Printf("Current State2: %s\n", gs.FSM.Current())

	gs.FSM.Event(context.Background(), BeginPreparation.String())
	fmt.Printf("Current State3: %s\n", gs.FSM.Current())

	gs.FSM.Event(context.Background(), EnterScoring.String())
	fmt.Printf("Current State: %s\n", gs.FSM.Current())

	gs.FSM.Event(context.Background(), EnterBuying.String())
	fmt.Printf("Current State: %s\n", gs.FSM.Current())

	gs.FSM.Event(context.Background(), EnterRubySpending.String())
	fmt.Printf("Current State: %s\n", gs.FSM.Current())
	if stateDestination == RubySpendingState.String() {
		return
	}
	if stateDestination == RubySpendingInputState.String() {
		return
	}
}
