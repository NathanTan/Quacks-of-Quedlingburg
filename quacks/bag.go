package quacks

import (
	"fmt"
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

	slice := bag.RemainingChips
	// rand.Seed(time.Now().UnixNano())
	// index := rand.Intn(len(arr))

	// Set arr to the
	// Pop the last element
	slice, lastElement := slice[:len(slice)-1], slice[len(slice)-1]

	bag.RemainingChips = slice

	return lastElement
}

func RemoveChip(bag Bag, chip Chip) {
	// TODO: Implement
	fmt.Println("Not Yet Implemented.")
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
