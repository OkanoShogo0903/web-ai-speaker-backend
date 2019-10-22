package model

type State int

const (
	Search State = iota
	Bgm
)

func (c State) String() string {
	switch c {
	case Search:
		return "Search"
	case Bgm:
		return "Bgm"
	default:
		return "Unknown"
	}
}
