package quacks

type Fortune struct {
	Ability string
	id      int
}

func createFortunes() []Fortune {
	return []Fortune{
		{"Temp", 0},
		{"Temp2", 1}
	}
}

func pop(slice []Fortune) ([]Fortune, Fortune) {

	element := slice[len(slice)]
	if len(slice) > 0 {
		slice = slice[:len(slice)-1]
	}

	return slice, element
}
