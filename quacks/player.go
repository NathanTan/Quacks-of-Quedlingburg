package quacks

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
