package quacks

type Fortune struct {
	Ability string
	id      int
}

func createFortunes() []Fortune {
	return []Fortune{
		{"Temp", 0},
		{"Temp2", 1},
		{"Get a ruby or +1 VP", 2},
		{"Get a ruby or +1 VP", 3},
		{"Double the number of rat tails this round.", 4},
		{"Choose: Take 4 victory points OR remove 1 white 1-chip from your bag", 5},
	}
}

func pop(slice []Fortune) ([]Fortune, Fortune) {

	element := slice[len(slice)-1]
	if len(slice) > 0 {
		slice = slice[:len(slice)-1]
	}

	return slice, element
}
