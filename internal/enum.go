package internal

type Hand int

const (
	Rock Hand = iota + 1
	Paper
	Scissors
)

func (h Hand) String() string {
	switch h {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		return "Unknown"
	}
}
