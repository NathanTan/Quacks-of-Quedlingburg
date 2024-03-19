package quacks

import (
	"context"
	"fmt"
)

func (gs *GameState) DrawChip(playerName string) {
	if gs.FSM.Current() == PreparationState.String() {
		for i, player := range gs.players {
			if player.name == playerName {
				// Stop when they player's post has exploded
				if !player.IsDoneDrawing(7) {
					pullAndPlaceChip(&gs.players[i], gs.debug)
					fmt.Printf("Cherry Bomb value is now now now %d\n", player.board.cherryBombValue)
					fmt.Printf("Remaining Chips for '%s': %d\n", playerName, len(player.bag.RemainingChips))

				}

				// Let the player decide if they want to pull more chips
				gs.FSM.Event(context.Background(), PreparationInput.String())
				gs.Awaiting = &Input{
					Description: "Please choose 1 to attempt to pull another chip, and 2 to be done",
					Options:     []string{"1", "2"},
					Player:      i,
				}
			}
		}

	}
}

func (gs *GameState) GetPlayerName(position int) string {
	return gs.players[position].name
}
func (gs *GameState) Input(input Input) Error {
	if gs.FSM.Current() == HandleFortune.String() {
		handleFortune(gs, input, gs.fortune, gs.debug)
	}
	if gs.FSM.Current() == FortuneInputState.String() {
		handleFortune(gs, input, gs.fortune, gs.debug)
		gs.FSM.Event(context.Background(), HandleFortune.String())
	}

	if gs.FSM.Current() == PreparationInputState.String() {
		// if gs.debug {
		fmt.Printf("Player %d chose %d\n", input.Player, input.Choice)
		// }/s
		// Choice of 1 is pulling another chip
		if input.Choice == 1 {
			gs.DrawChip(gs.GetPlayerName(gs.Awaiting.Player))
		}

		// Choice of 2 is done pulling
		if input.Choice == 2 {
			gs.players[input.Player].isDoneDrawing = true
			if gs.debug {
				fmt.Printf("Player '%s' is done drawing %t", gs.players[input.Player].name, gs.players[input.Player].isDoneDrawing)
			}
		}

		gs.FSM.Event(context.Background(), HandlePreparationInput.String())
	}

	if gs.FSM.Current() == BuyingInputState.String() {
		playerNumber := input.Player

		// Quick validation to prevent buying 2 chips
		if len(input.Choice2) == 2 && input.Choice2[0].color == input.Choice2[1].color {
			return Error{
				Description: "Error: Cannot buy 2 of the same type of chip",
			}
		} else if len(input.Choice2) > 2 {
			return Error{
				Description: "Error: Attempting to buy too many chip",
			}
		} else if gs.players[playerNumber].isDoneDrawing {
			gs.FSM.Event(context.Background(), EnterBuying.String())
			return Error{
				Description: "Error: Player is already done buying chips",
			}
		}

		buyingPower, _ := GetScores(gs.players[playerNumber].board)
		gs.players[playerNumber].buyingPower = buyingPower

		fmt.Printf("Buying power for %s - %d\n", gs.players[playerNumber].name, buyingPower)

		chipsCostMap := GetChipsValueMap(gs.book)

		if gs.debug {
			fmt.Println("desired buy: " + ChipsString(input.Choice2))
		}

		totalCost := 0

		for _, chip := range input.Choice2 {
			totalCost += totalCost + chipsCostMap[chip.color+fmt.Sprintf("%d", chip.value)]
		}

		fmt.Printf("Total desired chip cost: %d\n", totalCost)

		if totalCost <= buyingPower {
			for _, chip := range input.Choice2 {
				gs.players[playerNumber].bag.AddChip(chip)
				if gs.debug {
					fmt.Printf("Player '%s' bought a chip: %s\n", gs.players[playerNumber].name, chip.String())
				}
			}
		}

		gs.players[playerNumber].isDoneDrawing = true

		playersAreDoneDrawing := true
		for _, player := range gs.players {
			playersAreDoneDrawing = playersAreDoneDrawing && player.isDoneDrawing
		}

		if playersAreDoneDrawing {
			gs.FSM.Event(context.Background(), LeaveBuying.String())
		} else {
			gs.FSM.Event(context.Background(), EnterBuying.String())

		}
	}

	if gs.FSM.Current() == RubySpendingInputState.String() {
		if gs.Awaiting == nil {
			return Error{"Not waiting on any player's choice", gs.Awaiting.Player}
		}
		playerNumber := gs.Awaiting.Player

		// If the player number is too big
		if playerNumber > len(gs.players)-1 {
			if gs.debug {
				fmt.Println("Player doesn't exists")
			}
			return Error{"Player doesn't exists", gs.Awaiting.Player}
		}

		if gs.players[playerNumber].hasSpentRubies {
			if gs.debug {
				fmt.Println("Player has spent rubies")
			}

			remaingPlayerNames := gs.GetRemainingRubySpendingPlayerNames()
			if len(remaingPlayerNames) > 0 {
				stillAwaitingOnPlayer := gs.GetPlayerPosition(remaingPlayerNames[0])
				gs.Awaiting = &Input{
					Description: "Please select 1 for spending Rubies on the dropper or 2 for refilling your flask",
					Choice:      -1,
					Player:      stillAwaitingOnPlayer,
					Code:        getInputCodes()["VPOrBuying"],
				}
			}
			return Error{}
		}

		// Finished buying for the last player
		if len(gs.GetRemainingRubySpendingPlayerNames()) == 0 {
			gs.Awaiting = nil
			fmt.Println("Done with Rubies")
			gs.FSM.Event(context.Background(), EnterNextRound.String())
			return Error{}
		}

		player := gs.players[playerNumber]

		// Spend the rubies
		if player.rubyCount >= 2 && !player.hasSpentRubies {

			fmt.Printf("Spending rubies for player %s, choice: %d\n", gs.players[playerNumber].name, input.Choice)

			// Move the dropper
			if input.Choice == 1 {
				player.moveDropper(1, 1)
				fmt.Printf("DROPPER: Player %s ruby count: %d, dropper value: %d\n", gs.players[playerNumber].name, gs.players[playerNumber].rubyCount, gs.players[playerNumber].dropplet)
				gs.players[playerNumber].rubyCount -= 2
				gs.players[playerNumber].dropplet += 1
				fmt.Printf("'DROPPER': Player %s ruby count: %d, dropper value: %d\n", gs.players[playerNumber].name, gs.players[playerNumber].rubyCount, gs.players[playerNumber].dropplet)
			} else if input.Choice == 2 {
				if !player.flask {
					player.flask = true
					fmt.Printf("FLASK: Player %s ruby count: %d", gs.players[playerNumber].name, gs.players[playerNumber].rubyCount)
					gs.players[playerNumber].rubyCount -= 2
					gs.players[playerNumber].flask = true
					fmt.Printf("FLASK: Player %s ruby count: %d", gs.players[playerNumber].name, gs.players[playerNumber].rubyCount)

				}
			} else {
				gs.players[playerNumber].hasSpentRubies = true
			}

		} else {
			if gs.debug {
				fmt.Printf("Player %s doesn't have enough rubies to spend\n", player.name)
			}
			gs.players[playerNumber].hasSpentRubies = true
			// gs.Awaiting.Player += 1
		}

		// TODO: Fix around here
		if player.hasSpentRubies {
			gs.FSM.Event(context.Background(), EnterRubySpending.String())
			gs.Awaiting = nil
		}

		// if !player.hasSpentRubies {
		names := gs.GetRemainingRubySpendingPlayerNames()
		if len(names) > 0 {
			gs.Awaiting.Player = gs.GetPlayerPositionByName(names[0])
		}
		// }
	}

	return Error{}
}

func (gs *GameState) GetPlayerPositionsWithRubies() []int {
	pos := []int{}
	for i, player := range gs.players {
		fmt.Printf("Player %d, ruby count %d\n", i, player.rubyCount)
		if player.rubyCount >= 2 {
			pos = append(pos, i)
		}
	}

	return pos
}

func (gs *GameState) EndRubyBuys() {
	if gs.FSM.Current() == RubySpendingState.String() || gs.FSM.Current() == RubySpendingInputState.String() {
		gs.FSM.Event(context.Background(), EnterNextRound.String())
	}
}

func (gs *GameState) ResumePlay() {
	if gs.GameIsOver() {
		return
	}

	// Fortune cards
	if gs.FSM.Current() == FortuneState.String() {
		if gs.fortune != -1 {
			// Draw a fortune for the first time
			drawFortune(gs, gs.fortuneDeck, gs.debug)
		}

		// Check if all players have inputted for their fortune
		if len(gs.GetRemainingFortunePlayers()) == 0 {
			gs.FSM.Event(context.Background(), AssignRatTails.String())
		}
	}

	// Rat Tails
	if gs.FSM.Current() == RatTailsState.String() {
		assignRatTails(&gs.players, gs.debug)
		gs.FSM.Event(context.Background(), BeginPreparation.String())
	}

	// Draw Chips
	if gs.FSM.Current() == PreparationState.String() {
		// Shuffle Player's bags
		shufflePlayersBags(&gs.players)

		// Reset the player's board before assigning chips begins
		for i := range gs.players {
			if gs.players[i].board.nextPosition == -1 {
				gs.players[i].board.nextPosition = 0
			}
		}

		// Pull Chips (1 chip for now)
		// for i := range gs.players {
		// 	player := &gs.players[i]

		// 	if !player.IsDoneDrawing(7) {
		// 		pullAndPlaceChip(player, gs.debug)
		// 	}
		// }

		// Make it so they check if they're done, and if so move to the next game state - TODO: Pick up here
		playersAreDone := true
		for i := range gs.players {
			// if gs.debug {
			// 	fmt.Printf("Player '%s' is done drawing chips - %t\n", gs.players[i].name, gs.players[i].IsDoneDrawing(gs.bombLimit))
			// }
			playersAreDone = playersAreDone && gs.players[i].IsDoneDrawing(gs.bombLimit)
		}

		if gs.debug {
			fmt.Printf("playersAreDone: %t\n", playersAreDone)
		}

		if playersAreDone {
			gs.FSM.Event(context.Background(), EnterScoring.String())
		}
	}

	// Evaluation
	if gs.FSM.Current() == ScoringState.String() {

		// Bonus Dice
		for i := range gs.players {
			// If the player exploaded, ask for VP or buying
			if gs.players[i].PlayersPotHasExploaded() {
				gs.Awaiting = &Input{
					Description: fmt.Sprintf("Player %s has exploded their pot, please select from the options:\n\t1: Buy Chips\n\t2: Gain Victory Points", gs.players[i].name),
					Choice:      -1,
					Player:      i,
					Code:        getInputCodes()["VPOrBuying"],
				}
				gs.FSM.Event(context.Background(), HandleScoringInput.String())

			} else {
				gs.players[i].RollBonusDice(gs.debug)
			}
		}

		// Special Chips
		handleSpecialChips(&gs.players, gs.book, gs.debug)

		// Rubies
		handleRubies(&gs.players, gs.debug)

		// Victory Points
		handleVictoryPoints(gs, gs.debug)

		// Buy Chips
		gs.FSM.Event(context.Background(), EnterBuying.String())
		gs.Awaiting = nil
		// Spend Rubys

		logPlayers(gs.players)

	} else if gs.FSM.Current() == ScoringInputState.String() {
		playerId := gs.Awaiting.Player
		switch gs.Awaiting.Code {
		case getInputCodes()["VPOrBuying"]:
			if gs.Awaiting.Choice == 1 {
				gs.players[playerId].chooseBuying = true
			} else {
				gs.players[playerId].chooseVictoryPoints = true
			}
		}

		gs.FSM.Event(context.Background(), HandleScoringInput.String())
	}

	// Buying
	if gs.FSM.Current() == BuyingState.String() {
		if len(gs.GetRemainingBuyingPlayers()) == 0 {
			gs.Awaiting = nil
			gs.FSM.Event(context.Background(), EnterRubySpending.String())
			return // TODO Refactor
		} else if gs.Awaiting == nil {
			gs.Awaiting = &Input{
				"Please select a chip to buy",
				ListAvailableChips(gs.Round),
				0,
				nil,
				0,
				0,
			}
		} else if gs.Awaiting.Player < len(gs.players) && gs.Round == 9 {
			player := gs.players[gs.Awaiting.Player]

			//
			if !player.PlayersPotHasExploaded() || player.chooseBuying {
				player.score += player.buyingPower
				if gs.debug {
					fmt.Printf("Increase player '%s' score by %d to '%d'", player.name, player.buyingPower, player.score)
				}
			}
		} else if gs.Awaiting.Player < len(gs.players) {
			nextPlayer := gs.Awaiting.Player + 1 // TODO: This could be a bug and ask for the 4th player or something high
			gs.Awaiting = &Input{
				"Please select a chip to buy",
				ListAvailableChips(gs.Round),
				0,
				nil,
				nextPlayer,
				0,
			}
		}
		gs.FSM.Event(context.Background(), HandleBuying.String())
	}

	// Ruby Spending
	if gs.FSM.Current() == RubySpendingState.String() {
		stateMessage := "Please select 1 for spending Rubies on the dropper or 2 for refilling your flask"

		if gs.Awaiting != nil && gs.debug {
			fmt.Printf("Awaiting on Player '%d', len:%d\n", gs.Awaiting.Player, len(gs.players))
		}

		// Finished buying for the last player
		if gs.Awaiting != nil && gs.Awaiting.Player == len(gs.players) {
			gs.Awaiting = nil
			gs.FSM.Event(context.Background(), EnterNextRound.String())
			return
		}

		// Set Awaiting
		if gs.Awaiting == nil {
			gs.Awaiting = &Input{
				stateMessage,
				ListAvailableChips(gs.Round),
				0,
				nil,
				0,
				0,
			}
		} else if gs.Awaiting.Player < len(gs.players) {
			remainingPlayers := gs.GetRemainingRubySpendingPlayerNames()
			fmt.Printf("Remaining Players %v\n", remainingPlayers)
			if len(remainingPlayers) > 0 {

				nextPlayer := gs.getPlayerPosition(remainingPlayers[0])
				fmt.Printf("Next Player Position: %d\n", nextPlayer)

				gs.Awaiting = &Input{
					stateMessage,
					ListAvailableChips(gs.Round),
					0,
					nil,
					nextPlayer,
					0,
				}
			}
		}

		if gs.Round == 9 {
			fmt.Println("FINALROUND ------------------------------++++++.//")
			for i, player := range gs.players {
				if player.rubyCount >= 2 {
					gs.players[i].score += player.rubyCount / 2 // Go will round down here
				}
			}
			fmt.Printf("Current state: %s\n", gs.FSM.Current())
			gs.FSM.Event(context.Background(), End.String())
		} else {
			gs.FSM.Event(context.Background(), HandleRubySpending.String())
		}
	}

}

func pullAndPlaceChip(player *Player, debug bool) {

	chip := DrawChip(&(player.bag), debug)

	// No chip was pulled
	if chip.value == 0 {
		return
	}

	player.board.placeChip(chip)
	fmt.Printf("Cherry Bomb value is now now %d\n", player.board.cherryBombValue)

	if debug {
		fmt.Printf("Player %s draws a %s %d chip\n", player.name, chip.color, chip.value)
		fmt.Printf("Pot: %s\n", player.board.toString())
		fmt.Printf("RemainingChips: %s\n", player.bag.RemainingChips)
		fmt.Printf("All Chips: %s\n", player.bag.Chips)
	}
}

func (gs *GameState) StartGame() {
	gs.FSM.Event(context.Background(), Start.String())
	// Fortune cards
	drawFortune(gs, gs.fortuneDeck, gs.debug)

	// If input is required, pause for the input
	if gs.FSM.Current() == FortuneInputState.String() {
		if gs.fortune == 5 {
			gs.Awaiting = &Input{
				Description: "Choose: Take 4 victory points OR remove 1 white 1-chip from your bag" + ". (Choose 1 or 2)",
				Options:     []string{"1", "2"},
				Choice:      0,
				Choice2:     []Chip{}, // todo refactor for buying chips
				Player:      0,        // Player Position
				Code:        5,
			}
		}
		return
	}

	// Handle rat tails for fortune 4
	if gs.fortune == 4 {
		if gs.debug {
			fmt.Println("Assigning Rat Tails")
			fmt.Printf("score: %d", gs.players[0].score)
		}
		assignRatTails(&gs.players, gs.debug)
		for i := range gs.players {
			gs.players[i].ratToken += gs.players[i].ratToken
		}
		fmt.Printf("player.ratToken = %d\n", gs.players[3].ratToken)

	}

	// If no input is required, no rat tails are assigned since it's the first round
	if gs.FSM.Current() == AssignRatTails.String() {
		// No Rat Tails Assigned in the first round
		// assignRatTails(gs.players, gs.debug)
		gs.FSM.Event(context.Background(), BeginPreparation.String())
	}
}

func (gs GameState) GetRemainingPullingPlayerNames() []string {
	if gs.FSM.Current() == PreparationState.String() || gs.FSM.Current() == PreparationInputState.String() {
		names := []string{}
		for _, player := range gs.players {
			if !player.isDoneDrawing {
				names = append(names, player.name)
			}
		}
		return names
	}
	return []string{}
}

func (gs GameState) GetRemainingRubySpendingPlayerNames() []string {
	if gs.FSM.Current() == RubySpendingState.String() || gs.FSM.Current() == RubySpendingInputState.String() {
		names := []string{}
		for _, player := range gs.players {
			if gs.debug {
				fmt.Printf("%s - ruby count: %d\n", player.name, player.rubyCount)
			}
			if !player.hasSpentRubies && player.rubyCount >= 2 {
				names = append(names, player.name)
			}
		}
		return names
	}
	return []string{}
}

// TODO: is this a bug in the 2nd if???
func (gs GameState) GetRemainingBuyingPlayers() []string {
	if gs.FSM.Current() == BuyingState.String() || gs.FSM.Current() == BuyingInputState.String() {
		names := []string{}
		for _, player := range gs.players {
			if !player.isDoneDrawing {
				names = append(names, player.name)
			}
		}
		return names
	}
	return []string{}
}

func (gs GameState) GetRemainingFortunePlayers() []string {
	if gs.FSM.Current() == FortuneInputState.String() {
		names := []string{}
		for _, player := range gs.players {
			if !player.hasCompletedTheFortune {
				names = append(names, player.name)
			}
		}
		return names
	}
	return []string{}
}

func (gs GameState) GetPlayerPosition(name string) int {
	for i, player := range gs.players {
		if player.name == name {
			return i
		}
	}

	return -1
}

func (gs GameState) GetPlayerBombCountByName(name string) int {
	for _, player := range gs.players {
		if player.name == name {
			return player.board.getCherryBombValue()
		}
	}

	return -1
}

func (gs GameState) GetPlayerByName(name string) Player {
	for _, player := range gs.players {
		if player.name == name {
			return player
		}
	}

	return Player{}
}

func (gs GameState) GetPlayerPositionByName(name string) int {
	pos := -1
	for i, player := range gs.players {
		if player.name == name {
			pos = i
		}
	}

	return pos
}

func drawFortune(gs *GameState, fortuneDeck []Fortune, debug bool) {

	// gs.FSM.Event(context.Background(), ReadFortune.String())
	fmt.Println(gs.fortuneDeck)

	fortuneDeck, fortune := pop(fortuneDeck)
	if gs.debug {
		fmt.Printf("Fortune: %s", fortune.Ability)
	}
	gs.fortune = fortune.id
	if fortune.id == 2 {
		gs.FSM.Event(context.Background(), ReadFortune.String())
	} else if fortune.id == 3 {
		gs.FSM.Event(context.Background(), ReadFortune.String())
	} else if fortune.id == 5 {
		gs.FSM.Event(context.Background(), ReadFortune.String())
	}
	if debug {
		fmt.Println("Fortune for the round: " + fortune.Ability)
		fmt.Printf("Remaining Deck: %d\n", len(fortuneDeck))
	}
	if gs.FSM.Current() != FortuneInputState.String() {
		gs.FSM.Event(context.Background(), AssignRatTails.String())
	}
}

func (player *Player) RollBonusDice(debug bool) {
	// Else roll the bonus dice

	roll, description := BonusDiceRoll()
	switch roll {

	// Give a Pumpkin Chip
	case 1:
		player.bag.AddChip(NewChip(Orange.String(), 1))

		// Add 1 to their score
	case 2:
		player.score = player.score + 1

		// Add 2 to their score
	case 4:
		player.score = player.score + 2

	case 5:
		player.dropplet = player.dropplet + 1
		// TODO: Add test tube dropplet here

	case 6:
		player.rubyCount = player.rubyCount + 1

	}
	if debug {
		fmt.Println("Roll Result: " + description)
	}
}

func handleFortune(gs *GameState, input Input, fortune int, debug bool) {
	// Handle Input
	fmt.Printf("Fortune: %d, choice: %d\n", fortune, input.Choice)

	if fortune == 2 {
		if input.Choice == 1 {
			gs.players[input.Player].rubyCount = gs.players[input.Player].rubyCount + 1
		} else if input.Choice == 2 {
			gs.players[input.Player].score = gs.players[input.Player].score + 1
		}
	} else if fortune == 3 {
		for _, player := range gs.players {
			player.RollBonusDice(gs.debug)
		}
	} else if fortune == 5 {
		// Gain 4 VP
		if input.Choice == 1 {
			gs.players[input.Player].score += 4
		} else if input.Choice == 2 {
			// Remove a 1 Cherry bomb perminately
			gs.players[input.Player].bag.DeleteChip(NewChip(White.String(), 1))
		}
		gs.players[input.Player].hasCompletedTheFortune = true
	} else {
		fmt.Printf("========================\nFORTUNE NOT HANDLED YET - id: %d \n========================", fortune)
	}
}

func handleRubies(players *[]Player, debug bool) {
	for i := 0; i < len(*players); i++ {
		rubyCount := AssignRubies((*players)[i].board)
		if rubyCount {
			(*players)[i].rubyCount = (*players)[i].rubyCount + 1
		}
	}
}

func handleVictoryPoints(gs *GameState, debug bool) {
	players := gs.players
	for i := range players {
		// If the pot didn't boil over or if it did and they chose scoring
		if players[i].board.cherryBombValue <= players[i].explosionLimit || players[i].chooseVictoryPoints {
			if debug {
				fmt.Printf("DEBUG: Assigning VPs to '%s'\n", players[i].name)
			}
			_, victoryPointsEarned := GetScores(players[i].board)
			players[i].score = players[i].score + victoryPointsEarned
		}
	}
}

func logPlayers(players []Player) {
	for i := 0; i < len(players); i++ {
		PrintPlayerStatuses(players[i])
	}
}

func assignRatTails(players *[]Player, debug bool) {
	// Find player with the highest score
	highestScorePlayer := (*players)[0]
	for i := 1; i < len(*players); i++ {
		if (*players)[i].score > highestScorePlayer.score {
			highestScorePlayer = (*players)[i]
		}
	}

	// if debug {
	// 	fmt.Printf("highestScorePlayer: %s, score: %d\n", highestScorePlayer.name, highestScorePlayer.score)
	// }

	for i := 0; i < len((*players)); i++ {
		player := i % len((*players))
		leftPlayer := (*players)[player]
		rightPlayer := highestScorePlayer

		ratTailCount := countRatTails(leftPlayer.score, rightPlayer.score)
		// if debug {
		// 	fmt.Printf("Player number: %d\nleftPlayer: %s, score: %d\nrightPlayer: %s, score: %d\nRat Tail Count: %d\n",
		// 		player, leftPlayer.name, leftPlayer.score, rightPlayer.name, rightPlayer.score, ratTailCount)
		// }
		(*players)[player].ratToken = ratTailCount
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

func (gs *GameState) getPlayerPosition(name string) int {
	for i, player := range gs.players {
		if player.name == name {
			return i
		}
	}

	return -1 // Player not found
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
	for i := range ratTailList {
		if ratTailList[i] >= float32(highScore) {
			break
		}
		if ratTailList[i] > float32(lowScore) {
			tailCount = tailCount + 1
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

	if debug {
		fmt.Printf("player board: %s\n", (*players)[0].board.toString())
		fmt.Printf("player board: %s\n", (*players)[1].board.toString())
		fmt.Printf("player board: %s\n", (*players)[2].board.toString())
		fmt.Printf("player board: %s\n", (*players)[3].board.toString())
	}

	if book == 1 {
		for i := range *players {
			player := (*players)[i]
			// Handle Spiders
			chips := player.board.chips

			// If the last or second to last chip is a spider
			if chips[len(chips)-1].color == Green.String() || chips[len(chips)-2].color == Green.String() {
				// Give the player a ruby
				player.rubyCount = player.rubyCount + 1
			}

			// Handle Ghosts
			ghostCount := GetChipCount(player.board, Purple.String(), debug)
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
