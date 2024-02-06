package quacks

import "fmt"

type Board struct {
	chips            []Chip
	nextPosition     int
	testTubePosition int
	cherryBombValue  int
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
		{9, 9, 1, true},
		{10, 10, 2, false},
		{11, 11, 2, false},
		{12, 12, 2, false},
		{13, 13, 2, true},
		{14, 14, 3, false},
		{15, 15, 3, false},
		{16, 15, 3, true},
		{17, 16, 3, false},
		{18, 16, 4, false},
		{19, 17, 4, false},
		{20, 17, 4, true},
		{21, 18, 4, false},
		{22, 18, 5, false},
		{23, 19, 5, false},
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

func GetChipCount(board Board, _type string, debug bool) int {
	count := 0
	for _, chip := range board.chips {
		if chip.color == _type {
			count = count + 1
		}
	}

	return count
}

func (b Board) toString() string {
	s := ""
	for _, c := range b.chips {
		s += fmt.Sprintf("%s %d, ", c.color, c.value)
	}
	return s
}

func (b *Board) placeChip(chip Chip) {
	if chip.color == White.String() {
		b.cherryBombValue = b.cherryBombValue + chip.value
	}
	b.chips = append(b.chips, chip)
	b.nextPosition = b.nextPosition + 1
}

func (b Board) getCherryBombValue() int {
	return b.cherryBombValue
}
