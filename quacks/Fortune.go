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
	}
}

func pop(slice []Fortune) ([]Fortune, Fortune) {

	element := slice[len(slice)-1]
	if len(slice) > 0 {
		slice = slice[:len(slice)-1]
	}

	return slice, element
}
