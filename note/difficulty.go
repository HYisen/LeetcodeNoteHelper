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
