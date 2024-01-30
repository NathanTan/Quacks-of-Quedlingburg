package main

import (
	q "quacks"
)

func main() {

	gs := q.CreateGameState([]string{"Nathan", "Leah", "Raymond", "Hannah"}, true)

	gs.StartGame()

	for i := 0; i < 10; i++ {
		gs.ResumePlay()
		gs.Input(q.Input{Description: "", Choice: 1, Player: 1})
		gs.Input(q.Input{Description: "", Choice: 1, Player: 2})
		gs.Input(q.Input{Description: "", Choice: 1, Player: 3})
		gs.Input(q.Input{Description: "", Choice: 1, Player: 4})
		if gs.GameIsOver() {
			break
		}
	}
}
