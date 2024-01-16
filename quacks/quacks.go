package quacks

import (
	"fmt"
)

func PlayGame(debug bool) {
	gs := GameState{}
	players := []Player{}

	for i := 1; i < 9; i++ {
		fmt.Println("Starting Round.")

		// Fortune cards

		// Rat Tails

		// Pull Chips

		// Evaluation

		// Bonus Dice

		// Special Chips

		// Rubies
		handleRubies(players, debug)

		// Victory Points
		handleVictoryPoints(players, debug)

		// Buy Chips

		/// Spend Rubys

		// End the game after nine turns
		if gs.turn > 9 {
			break
		}

		// Game is over, maybe someone surrendered
		if GameIsOver(gs) {
			break
		}
	}

}

func handleRubies(players []Player, debug bool) {
	for i := 0; i < len(players); i++ {
		if debug {
			fmt.Println("DEBUG: Assigning Rubies")
		}
		rubyCount := AssignRubies(players[i].board)
		if rubyCount {
			players[i].rubyCount = players[i].rubyCount + 1
		}
	}
}

func handleVictoryPoints(players []Player, debug bool) {
	for i := 0; i < len(players); i++ {
		if debug {
			fmt.Println("DEBUG: Assigning VPs")
		}
		_, victoryPointsEarned := GetScores(players[i].board)
		players[i].score = players[i].score + victoryPointsEarned
	}
}
