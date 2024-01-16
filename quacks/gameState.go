package quacks

type GameState struct {
	turn    int
	fortune int
	winner  int
}

func GameIsOver(gs GameState) bool {
	if gs.winner > 0 {
		return true
	}
	return false
}
