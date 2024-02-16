package quacks

import (
	q "quacks"
	"testing"
)

func TestRemoveChipFromBag(t *testing.T) {
	b := q.Bag{
		Chips:          []q.Chip{q.NewChip("yellow", 2), q.NewChip("yellow", 4)},
		RemainingChips: []q.Chip{q.NewChip("yellow", 2)},
	}

	b.RemoveChip(q.NewChip("yellow", 2))

	if len(b.RemainingChips) != 0 {
		t.Errorf("Error: Too many remaining chips")
	}

	if len(b.Chips) != 2 {
		t.Errorf("Error: Too many remaining chips")
	}
}

func TestRemove1ChipFromBag(t *testing.T) {
	b := q.Bag{
		Chips:          []q.Chip{q.NewChip("yellow", 2), q.NewChip("yellow", 2)},
		RemainingChips: []q.Chip{q.NewChip("yellow", 2), q.NewChip("yellow", 2)},
	}

	b.RemoveChip(q.NewChip("yellow", 2))

	if len(b.RemainingChips) != 1 {
		t.Errorf("Error: Too many remaining chips")
	}

	if len(b.Chips) != 2 {
		t.Errorf("Error: Too many remaining chips")
	}
}
