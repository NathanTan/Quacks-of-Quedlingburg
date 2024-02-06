package main

import (
	"fmt"
	q "quacks"
)

func main() {

	gs := q.CreateGameState([]string{"Nathan", "Leah", "Raymond", "Hannah"}, true)

	gs.StartGame()

	for i := 0; i < 100; i++ {
		gs.ResumePlay()
		fmt.Printf("Current state: %s ------------------------'\n", gs.FSM.Current())
		if gs.FSM.Current() == q.End.String() {
			fmt.Println("Game is over")

		} else if gs.FSM.Current() != q.BuyingInputState.String() && gs.FSM.Current() != q.RubySpendingInputState.String() {

			gs.Input(q.Input{Description: "", Choice: 1, Player: 1})
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
			fmt.Printf("Players that still need to buy: %s\n", gs.GetRemainingBuyingPlayers())
		} else if gs.FSM.Current() == q.RubySpendingInputState.String() {
			fmt.Printf("Awaiting on Player '%d' to '%s'\n", gs.Awaiting.Player, gs.Awaiting.Description)
			gs.Input(q.Input{Description: "", Choice: 1, Player: 0})
			gs.ResumePlay()

			gs.Input(q.Input{Description: "", Choice: 1, Player: 1})
			gs.ResumePlay()

			gs.Input(q.Input{Description: "", Choice: 1, Player: 2})
			gs.ResumePlay()

			gs.Input(q.Input{Description: "", Choice: 1, Player: 3})
			gs.ResumePlay()

		}
		if gs.GameIsOver() {
			break
		}
	}
}
