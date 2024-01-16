package quacks

import (
	"fmt"
)

type Player struct {
	name           string
	bag            Bag
	board          Board
	rubyCount      int
	ratToken       int
	dropplet       int
	flask          bool
	explosionLimit int
	score          int
}

func PrintPlayerStatuses(player Player) {
	fmt.Println("Player:")
	fmt.Println(player)
}
