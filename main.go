package main

import (
	q "quacks"
)

func main() {

	gs := q.CreateGameState([]string{"Nathan", "Leah", "Raymond", "Hannah"}, true)

	gs.StartGame()

	for i := 0; i < 10; i++ {
		gs.ResumePlay()
		if gs.FSM.Current() != q.BuyingInputState.String() {

			gs.Input(q.Input{Description: "", Choice: 1, Player: 1})
			gs.Input(q.Input{Description: "", Choice: 1, Player: 2})
			gs.Input(q.Input{Description: "", Choice: 1, Player: 3})
			gs.Input(q.Input{Description: "", Choice: 1, Player: 4})
		} else {
			gs.Input(q.Input{Description: "", Choice: 0, Choice2: []q.Chip{q.NewChip("Orange", 1)}, Player: 1})

		}
		if gs.GameIsOver() {
			break
		}
	}
}
