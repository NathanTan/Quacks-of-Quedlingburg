package quacks_test

import (
	q "quacks"
	"testing"
)

func TestAddition(t *testing.T) {

	bag := q.Bag{
		Chips: []q.Chip{
			q.Chip{color: "blue", value: 4},
			q.Chip{color: "green", value: 1},
		},
		RemainingChips: []q.Chip{
			q.Chip{color: "blue", value: 4},
			q.Chip{color: "green", value: 1},
		},
	}

	q.DrawChip(bag, true)

	result := 2 + 2
	if result != 4 {
		t.Errorf("Expected 2 + 2 to equal 4, but got %d instead", result)
	}
}
