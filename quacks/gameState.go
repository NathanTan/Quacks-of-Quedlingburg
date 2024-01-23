package quacks

type GameState struct {
	turn    int
	fortune int
	winner  int
	book    int
}

func CreateGameStates() GameState {
	return GameState{
		0,
		0,
		0,
		1,
	}
}

func GameIsOver(gs GameState) bool {
	if gs.winner > 0 {
		return true
	}
	return false
}
