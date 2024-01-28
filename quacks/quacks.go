package quacks

import (
	"context"
	"fmt"
)

func (gs GameState) Input(input Input) {
	if gs.FSM.Current() == HandleFortune.String() {
		handleFortune(&gs, input, gs.fortune, gs.debug)
	}

}

func (gs GameState) StartGame() {
	fortuneDeck := createFortunes()
	gs.FSM.Event(context.Background(), "start")
	for i := 1; i < 9; i++ {
		fmt.Println("Starting Round.")

		// Fortune cards
		drawFortune(gs, fortuneDeck, gs.debug)

		// Rat Tails
		assignRatTails(gs.players, gs.debug)

		if gs.debug {
			fmt.Println("Drawing Chips \n")
		}

		// Shuffle Player's bags
		shufflePlayersBags(&gs.players)

		// Pull Chips (1 chip for now)
		for _, player := range gs.players {
			if len(player.bag.Chips) > 0 {
				chip := DrawChip(&(player.bag), gs.debug)
				player.board.chips = append(player.board.chips, chip)
				if gs.debug {
					fmt.Printf("Player %s draws a %s %d chip\n", player.name, chip.color, chip.value)
				}
			}
		}

		// Evaluation

		// Bonus Dice
		for i := 0; i < len(gs.players); i++ {

			roll, description := BonusDiceRoll()
			switch roll {

			// Give a Pumpkin Chip
			case 1:
				AddChip(&gs.players[i].bag, NewChip("orange", 1))

				// Add 1 to their score
			case 2:
				gs.players[i].score = gs.players[i].score + 1

				// Add 2 to their score
			case 4:
				gs.players[i].score = gs.players[i].score + 2

			case 5:
				gs.players[i].dropplet = gs.players[i].dropplet + 1
				// TODO: Add test tube dropplet here

			case 6:
				gs.players[i].rubyCount = gs.players[i].rubyCount + 1

			}
			if gs.debug {
				fmt.Println("Roll Result: " + description)
			}
		}

		// Special Chips
		handleSpecialChips(&gs.players, gs.book, gs.debug)

		// Rubies
		handleRubies(gs.players, gs.debug)

		// Victory Points
		handleVictoryPoints(gs.players, gs.debug)

		// Buy Chips

		/// Spend Rubys

		logPlayers(gs.players)

		if gs.debug {
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

}

func drawFortune(gs GameState, fortuneDeck []Fortune, debug bool) {

	fmt.Println(fortuneDeck)
	gs.FSM.Event(context.Background(), "read_fortune")

	fortuneDeck, fortune := pop(fortuneDeck)
	if fortune.id == 2 {
		gs.FSM.Event(context.Background(), GetHandleFortuneString())
	}
	fmt.Println("Fortune for the round: " + fortune.Ability)
	if debug {
		fmt.Println(fmt.Sprintf("Remaining Deck: %d", len(fortuneDeck)))
	}
}

func handleFortune(gs *GameState, input Input, fortune int, debug bool) {
	// Handle Input
	if fortune == 2 {
		if input.Choice == 1 {
			gs.players[input.Player].rubyCount = gs.players[input.Player].rubyCount + 1
		} else if input.Choice == 2 {
			gs.players[input.Player].score = gs.players[input.Player].score + 1
		}
	}
	return

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

func handleSpecialChips(players *[]Player, book int, debug bool) {

	// Handle Moths
	playerMothCounts := []int{}

	// Count all the moths
	for _, player := range *players {
		playerMothCounts = append(playerMothCounts, GetChipCount(player.board, "black", debug))
	}

	if len(*players) > 2 {
		for i, player := range *players {
			nextValue := playerMothCounts[(i+1)%len(*players)]
			prevValue := playerMothCounts[(i-1+len(*players))%len(*players)]

			if playerMothCounts[i] > nextValue && playerMothCounts[i] > prevValue {
				player.rubyCount = player.rubyCount + 1
				// TODO: Add option for increasing test tube dropplet
				player.dropplet = player.dropplet + 1
			} else if playerMothCounts[i] > nextValue || playerMothCounts[i] > prevValue {
				// TODO: Add option for increasing test tube dropplet
				player.dropplet = player.dropplet + 1
			}
		}
	} else {
		for i, player := range *players {
			nextValue := playerMothCounts[(i+1)%len(*players)]

			if playerMothCounts[i] > nextValue {
				player.rubyCount = player.rubyCount + 1
				// TODO: Add option for increasing test tube dropplet
				player.dropplet = player.dropplet + 1
			} else if playerMothCounts[i] == nextValue {
				// TODO: Add option for increasing test tube dropplet
				player.dropplet = player.dropplet + 1
			}
		}
	}

	if book == 1 {
		for _, player := range *players {
			// Handle Spiders
			chips := player.board.chips

			// If the last or second to last chip is a spider
			if chips[len(chips)].color == "green" || chips[len(chips)-1].color == "green" {
				// Give the player a ruby
				player.rubyCount = player.rubyCount + 1
			} else {
				if debug {
					fmt.Println("No Spider")
				}
			}

			// Handle Ghosts
			ghostCount := GetChipCount(player.board, "purple", debug)
			if ghostCount == 1 {
				player.score = player.score + 1
			} else if ghostCount == 2 {
				player.score = player.score + 2
				player.rubyCount = player.rubyCount + 1
			} else if ghostCount > 2 {
				player.score = player.score + 2
				player.dropplet = player.dropplet + 1
				// TODO: Add input for choice of dropplet or testTubeDropplet
			}
		}
	}
}
