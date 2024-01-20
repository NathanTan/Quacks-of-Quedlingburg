package quacks

import (
	"fmt"
)

func PlayGame(playerNames []string, debug bool) GameState {
	gs := GameState{}
	players := setUpPlayers(playerNames)
	fortuneDeck := createFortunes()
	for i := 1; i < 9; i++ {
		fmt.Println("Starting Round.")

		// Fortune cards
		fmt.Println(fortuneDeck)
		fortuneDeck, fortune := pop(fortuneDeck)
		fmt.Println("Fortune for the round: " + fortune.Ability)
		if debug {
			fmt.Println(fmt.Sprintf("Remaining Deck: %d", len(fortuneDeck)))
		}

		// Rat Tails
		assignRatTails(players, debug)

		// Pull Chips
		for i := 0; i < len(players); i++ {
			DrawChip(players[i].bag, true)
		}

		// Evaluation

		// Bonus Dice

		// Special Chips

		// Rubies
		handleRubies(players, debug)

		// Victory Points
		handleVictoryPoints(players, debug)

		// Buy Chips

		/// Spend Rubys

		logPlayers(players)

		if debug {
			break
		}

		// End the game after nine turns
		if gs.turn > 9 {
			break
		}

		// Game is over, maybe someone surrendered
		if GameIsOver(gs) {
			break
		}
	}

	return gs

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

func logPlayers(players []Player) {
	for i := 0; i < len(players); i++ {
		PrintPlayerStatuses(players[i])
	}
}

func assignRatTails(players []Player, debug bool) {
	// Find player with the highest score
	highestScorePlayer := players[0]
	for i := 1; i < len(players); i++ {
		if players[i].score > highestScorePlayer.score {
			highestScorePlayer = players[i]
		}
	}

	for i := 0; i < len(players); i++ {
		player := i % len(players)
		fmt.Printf("Player number: %d\n", player)
		leftPlayer := players[player]
		rightPlayer := highestScorePlayer
		fmt.Printf("leftPlayer: %s, rightPlayer: %s\n", leftPlayer.name, rightPlayer.name)

		ratTailCount := countRatTails(leftPlayer.score, rightPlayer.score)
		leftPlayer.ratToken = ratTailCount
		if debug {
			fmt.Printf("Player %q rat tail count: %d\n", leftPlayer.name, ratTailCount)
		}
	}

}

func getOtherPlayerPosition(players []Player, counter int) int {
	if counter+1 <= len(players)-1 {
		return counter + 1
	}
	return 0
}

func countRatTails(lowScore int, highScore int) int {
	if lowScore >= highScore {
		return 0 // early exit
	}

	// brute solution
	ratTailList := []float32{1.5, 4.5, 7.5, 10.5, 12.5, 14.5, 16.5, 18.5, 20.5,
		22.5, 24.5, 26.5, 28.5, 30.5, 32.5, 34.5, 36.5, 38.5, 40.5,
		42.5, 44.5, 46.5, 48.5, 51.5, 54.5, 57.5, 60.5, 62.5, 64.5,
		66.5, 68.5, 70.5, 72.5, 74.5, 76.5, 78.5, 80.5, 82.5, 84.5,
		86.5, 88.5, 90.5, 92.5, 94.5, 96.5, 98.5, 101.5, 104.5}

	tailCount := 0
	for i := 0; i < len(ratTailList); i++ {
		if ratTailList[i] > float32(lowScore) && ratTailList[i] < float32(highScore) {
			tailCount = tailCount + 1
		}

		if ratTailList[i] > float32(highScore) {
			break
		}
	}

	return tailCount
}
