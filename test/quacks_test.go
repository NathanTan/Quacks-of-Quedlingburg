package quacks_test

import (
	q "quacks"
	"testing"
)

func TestAddition(t *testing.T) {

	bag := q.Bag{
		Chips: []q.Chip{
			q.NewChip("blue", 4),
			q.NewChip("green", 1),
		},
		RemainingChips: []q.Chip{
			q.NewChip("blue", 4),
			q.NewChip("green", 1),
		},
	}

	if len(bag.RemainingChips) != 2 {
		t.Errorf("Error")
	}

	chip := q.DrawChip(&bag, true)

	if len(bag.RemainingChips) != 1 {
		t.Errorf("Error 2")
	}
	if len(bag.Chips) != 2 {
		t.Errorf("Error 3")
	}
	if chip != q.NewChip("green", 1) {
		t.Errorf("Error 4")
	}
}
