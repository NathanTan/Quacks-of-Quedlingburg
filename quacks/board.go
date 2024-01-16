package quacks

type Board struct {
	nextPosition     int
	testTubePosition int
}

type Slot struct {
	position     int
	pointValue   int
	vpPointValue int
	hasRuby      bool
}

func getStandardBoard() []Slot {
	return []Slot{
		{0, 0, 0, false},
		{1, 1, 0, false},
		{2, 2, 0, false},
		{3, 3, 0, false},
		{4, 4, 0, false},
		{5, 5, 0, true},
		{6, 6, 1, false},
		{7, 7, 1, false},
		{8, 8, 1, false},
	}
}

func SetUpBoard(droppletPosition int, ratTailCount int) Board {
	board := Board{nextPosition: droppletPosition + ratTailCount + 1}

	return board
}

// Buying Points and Victory Points.
func GetScores(board Board) (int, int) {
	standardBoard := getStandardBoard()
	buyingPoints := standardBoard[board.nextPosition].pointValue
	victoryPoints := standardBoard[board.nextPosition].vpPointValue
	return buyingPoints, victoryPoints
}

func AssignRubies(board Board) bool {
	standardBoard := getStandardBoard()
	rubys := standardBoard[board.nextPosition].hasRuby
	return rubys
}
