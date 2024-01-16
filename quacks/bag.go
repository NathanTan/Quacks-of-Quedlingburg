package quacks

import (
	"fmt"
	"math/rand"
	"time"
)

type Bag struct {
	Chips          []Chip
	RemainingChips []Chip
}

func AddChip(bag Bag, chip Chip) {
	bag.Chips = append(bag.Chips, chip)

}

func DrawChip(bag Bag, debug bool) {

	arr := bag.RemainingChips
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(arr))
	arr[index] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	bag.RemainingChips = arr
	if debug {
		fmt.Println("DEBUG:")
		fmt.Println(arr)
		fmt.Println(bag.RemainingChips)
	}
}

func RemoveChip(bag Bag, chip Chip) {
	// TODO: Implement
	fmt.Println("Not Yet Implemented.")
}
