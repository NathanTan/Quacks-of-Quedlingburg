package main

import (
	q "quacks"
)

func main() {

	gs := q.CreateGameState([]string{"Nathan", "Leah", "Raymond", "Hannah"}, true)

	gs.StartGame()

}
