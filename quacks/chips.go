package quacks

type Chip struct {
	color string
	value int
}

func NewChip(color string, value int) Chip {
	return Chip{color: color, value: value}
}

type ChipType int

const (
	Orange ChipType = iota
	Yellow
	Blue
	Red
	Green
	Purple
	Black
	White
	LocoWeed
)

func (t ChipType) String() string {
	switch t {
	case Yellow:
		return "yellow"
	case Blue:
		return "blue"
	case Red:
		return "red"
	case Green:
		return "green"
	case Purple:
		return "purple"
	case White:
		return "white"
	case Black:
		return "black"
	case LocoWeed:
		return "locoWeed"
	default:
		return "unknown"
	}
}
