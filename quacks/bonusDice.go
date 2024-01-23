package quacks

import (
	"math/rand"
)

func BonusDiceRoll() (int, string) {
	// Generate a random number between 0 and 5
	roll := rand.Intn(6)
	roll = roll + 1

	switch roll {
	case 1:
		return 1, "Collect A Pumpkin Chip"
	case 2:
	case 3:
		return 2, "1 Victory Point"
	case 4:
		return 4, "2 Victory Points"
	case 5:
		return 5, "Move a Dropplet"
	case 6:
		return 6, "Collect a Ruby"
	case 7:
		return 7, "error 7"
	}
	return 0, "error"
}
