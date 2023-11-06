package note

import "fmt"

type Difficulty int

const (
	Unknown Difficulty = iota
	Easy
	Medium
	Hard
)

func parseDifficulty(s string) (Difficulty, error) {
	switch s {
	case "E":
		return Easy, nil
	case "M":
		return Medium, nil
	case "H":
		return Hard, nil
	default:
		return Unknown, fmt.Errorf("unknown difficulty string %s", s)
	}
}

func (d Difficulty) String() string {
	switch d {
	case Unknown:
		return "Unknown"
	case Easy:
		return "Easy"
	case Medium:
		return "Medium"
	case Hard:
		return "Hard"
	}
	panic(fmt.Errorf("unimplemented on Difficuly %d", d))
}
