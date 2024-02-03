package quacks

import (
	"fmt"
	q "quacks"
	"testing"
)

func TestBonusDiceRoll(t *testing.T) {
	for i := 0; i < 100; i++ {

		value, desc := q.BonusDiceRoll()

		fmt.Println(desc)
		if value < 1 {
			t.Errorf("Error")
		}

	}
}
