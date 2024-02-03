package quacks

import "fmt"

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

func ListAvailableChips(round int) []string {
	chips := []string{
		Blue.String(),
		Red.String(),
		Green.String(),
		Black.String(),
	}

	if round >= 2 {
		chips = append(chips, Yellow.String())
	}

	if round >= 3 {
		chips = append(chips, Purple.String())
	}

	return chips
}

func GetChipsValueMap(book int) map[string]int {
	var mapp map[string]int
	if book == 1 {
		mapp = map[string]int{
			Orange.String() + "1": 3,
			Blue.String() + "1":   5,
			Blue.String() + "2":   10,
			Blue.String() + "4":   19,
			Red.String() + "1":    6,
			Red.String() + "2":    10,
			Red.String() + "4":    16,
			Yellow.String() + "1": 8,
			Yellow.String() + "2": 12,
			Yellow.String() + "4": 18,
			Green.String() + "1":  4,
			Green.String() + "2":  8,
			Green.String() + "4":  14,
			Purple.String() + "1": 9,
		}
	} else if book == 2 {
		mapp = nil
	}

	mapp[Black.String()+"1"] = 10

	return mapp
}

func (c Chip) String() string {
	return fmt.Sprintf("%s_%d", c.color, c.value)
}
