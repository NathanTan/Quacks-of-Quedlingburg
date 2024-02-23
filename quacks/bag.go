package quacks

import (
	"math/rand"
)

type Bag struct {
	Chips          []Chip
	RemainingChips []Chip
}

func (bag *Bag) AddChip(chip Chip) {
	bag.Chips = append(bag.Chips, chip)
}

func DrawChip(bag *Bag, debug bool) Chip {
	if len(bag.RemainingChips) == 0 {
		return Chip{value: 0, color: ""}
	}

	slice := bag.RemainingChips
	// rand.Seed(time.Now().UnixNano())
	// index := rand.Intn(len(arr))

	// Set arr to the
	// Pop the last element
	slice, lastElement := slice[:len(slice)-1], slice[len(slice)-1]

	bag.RemainingChips = slice

	bag.RemoveChip(lastElement)

	return lastElement
}

func (bag *Bag) RemoveChip(chip Chip) {
	for i := 0; i < len(bag.RemainingChips); i++ {
		if bag.RemainingChips[i].value == chip.value && bag.RemainingChips[i].color == chip.color {
			// Remove the element at index i from people.
			bag.RemainingChips = append(bag.RemainingChips[:i], bag.RemainingChips[i+1:]...)
			i-- // Decrement i since we just removed an element.
			break
		}
	}
}

func (bag *Bag) DeleteChip(chip Chip) {
	for i := 0; i < len(bag.Chips); i++ {
		if bag.Chips[i].value == chip.value && bag.Chips[i].color == chip.color {
			// Remove the element at index i from people.
			bag.Chips = append(bag.Chips[:i], bag.Chips[i+1:]...)
			i-- // Decrement i since we just removed an element.
			break
		}
	}
}

func shufflePlayersBags(players *[]Player) {
	for _, player := range *players {
		shufflePlayersBag(&player)
	}
}

func shufflePlayersBag(player *Player) {
	slice := player.bag.RemainingChips
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}
