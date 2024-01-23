package quacks

type Chip struct {
	color string
	value int
}

func NewChip(color string, value int) Chip {
	return Chip{color: color, value: value}
}
