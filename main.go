package main

import (
	"fmt"
	q "quacks"
)

func main() {
	playerNames := []string{"Nathan", "Leah", "Raymond", "Hannah"}

	gs := q.CreateGameState(playerNames, true)

	gs.StartGame()

	lastPlayerToPull := -1

	for i := 0; i < 100; i++ {
		gs.ResumePlay()
		fmt.Printf("Current state: %s ------------------------'\n", gs.FSM.Current())
		if gs.FSM.Current() == q.End.String() {
			fmt.Println("Game is over")
			break

		} else if gs.FSM.Current() == q.RubySpendingState.String() {
			gs.ResumePlay()
			fmt.Println(gs.Awaiting)
			players := gs.GetPlayerPositionsWithRubies()
			fmt.Printf("%d players left to spend rubies, %v\n", len(players), players)
			if len(players) > 0 {
				e := gs.Input(q.Input{
					Description: "",
					Options:     []string{},
					Choice:      1,
					Choice2:     []q.Chip{},
					Player:      players[0],
					Code:        -1,
				})

				if len(e.Description) > 0 {
					fmt.Printf("error: %s\n", e.Description)
				}
			} else {
				fmt.Println("Now ending ruby buys")
				gs.EndRubyBuys()

			}
			gs.ResumePlay()
			remainingPlayers := gs.GetRemainingRubySpendingPlayerNames()

			if len(remainingPlayers) == 0 {
				gs.EndRubyBuys()
				fmt.Println("Done buying rubies")
			}

		} else if gs.FSM.Current() == q.PreparationState.String() {

			if len(gs.GetRemainingPullingPlayerNames()[0]) > 0 {
				gs.DrawChip(gs.GetRemainingPullingPlayerNames()[0])
			}

		} else if gs.FSM.Current() == q.PreparationInputState.String() {
			// Since we are waiting

			lastPlayerToPull = (lastPlayerToPull + 1) % len(playerNames)

			// fmt.Printf("nextPlayerToPull: %d\n", lastPlayerToPull)
			if len(gs.GetRemainingPullingPlayerNames()) > 0 {
				name := gs.GetRemainingPullingPlayerNames()[0]
				fmt.Printf("nextPlayerToPull: %s, bombCount: %d\n", name, gs.GetPlayerBombCountByName(name))

				var choice int
				if gs.GetPlayerBombCountByName(name) >= 7 {
					choice = 2
				} else {
					choice = 1
				}

				gs.Input(q.Input{
					Description: gs.Awaiting.Description,
					Choice:      choice,
					Player:      gs.GetPlayerPosition(name),
				})
			}
		} else if gs.FSM.Current() == q.FortuneInputState.String() {

			gs.Input(q.Input{Description: "", Choice: 1, Player: 0})
			gs.ResumePlay()

			gs.Input(q.Input{Description: "", Choice: 1, Player: 1})
			gs.ResumePlay()

			gs.Input(q.Input{Description: "", Choice: 1, Player: 2})
			gs.ResumePlay()

			gs.Input(q.Input{Description: "", Choice: 1, Player: 3})
			gs.ResumePlay()

		} else if gs.FSM.Current() != q.BuyingInputState.String() && gs.FSM.Current() != q.RubySpendingInputState.String() {

			gs.Input(q.Input{Description: "", Choice: 1, Player: 0})
			gs.Input(q.Input{Description: "", Choice: 1, Player: 1})
			gs.Input(q.Input{Description: "", Choice: 1, Player: 2})
			gs.Input(q.Input{Description: "", Choice: 1, Player: 3})

		} else if gs.FSM.Current() == q.ScoringState.String() {
			gs.Input(q.Input{Description: "", Choice: 0, Player: 1})
			gs.Input(q.Input{Description: "", Choice: 1, Player: 2})
			gs.Input(q.Input{Description: "", Choice: 1, Player: 3})
			gs.Input(q.Input{Description: "", Choice: 1, Player: 4})

		} else if gs.FSM.Current() == q.BuyingInputState.String() {
			error := gs.Input(q.Input{Description: "", Choice: 0, Choice2: []q.Chip{q.NewChip(q.Orange.String(), 1)}, Player: 0})
			gs.ResumePlay()

			error = gs.Input(q.Input{Description: "", Choice: 0, Choice2: []q.Chip{q.NewChip(q.Orange.String(), 1)}, Player: 1})
			gs.ResumePlay()

			error = gs.Input(q.Input{Description: "", Choice: 0, Choice2: []q.Chip{q.NewChip(q.Orange.String(), 1)}, Player: 2})
			gs.ResumePlay()
			// fmt.Printf("Players that still need to buy: %s\n", gs.GetRemainingBuyingPlayers())

			error = gs.Input(q.Input{Description: "", Choice: 0, Choice2: []q.Chip{q.NewChip(q.Orange.String(), 1)}, Player: 3})
			if len(error.Description) > 0 {
				fmt.Println(error.Description)
			}
			gs.ResumePlay()

			// for buying phase
			// fmt.Printf("Players that still need to buy: %s\n", gs.GetRemainingBuyingPlayers())

		} else if gs.FSM.Current() == q.RubySpendingInputState.String() {
			remainingPlayers := gs.GetRemainingRubySpendingPlayerNames()

			if len(remainingPlayers) == 0 {
				gs.EndRubyBuys()
				fmt.Println("Done buying rubies")
			} else if gs.Awaiting != nil {

				fmt.Printf("remaing Players: %s", remainingPlayers)
				fmt.Printf("Awaiting on Player '%d' to '%s'\n", gs.Awaiting.Player, gs.Awaiting.Description)
				// for _, playerName := range remainingPlayers {
				gs.Input(q.Input{Description: "", Choice: 1, Player: gs.GetPlayerPositionByName(remainingPlayers[0])})
				gs.ResumePlay()
				gs.Input(q.Input{Description: "", Choice: -1, Player: gs.GetPlayerPositionByName(remainingPlayers[0])})
				gs.ResumePlay()
			}
			// }

			// gs.Input(q.Input{Description: "", Choice: 1, Player: 0})
			// gs.ResumePlay()
			// gs.Input(q.Input{Description: "", Choice: -1, Player: 0})
			// gs.ResumePlay()

			// gs.Input(q.Input{Description: "", Choice: 1, Player: 1})
			// gs.ResumePlay()
			// gs.Input(q.Input{Description: "", Choice: -1, Player: 1})
			// gs.ResumePlay()

			// gs.Input(q.Input{Description: "", Choice: 1, Player: 2})
			// gs.ResumePlay()
			// gs.Input(q.Input{Description: "", Choice: -1, Player: 2})
			// gs.ResumePlay()

			// gs.Input(q.Input{Description: "", Choice: 1, Player: 3})
			// gs.ResumePlay()
			// gs.Input(q.Input{Description: "", Choice: -1, Player: 3})
			// gs.ResumePlay()

		}

		if gs.GameIsOver() {
			fmt.Println("game is over")
			break
		}
	}

	gs.PrintRound()
}
