package diary

import (
	"fmt"
	"leetcodeNoteHelper/note"
	"strconv"
	"strings"
)

func NewProblem(r note.Record) Problem {
	return Problem{
		ID:          r.ProblemID,
		CostMinutes: int(r.End.Sub(r.Begin).Minutes()),
		Simple:      r.Simple,
	}
}

type Problem struct {
	ID          int
	CostMinutes int
	Simple      bool
}

func (p Problem) String() string {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(p.ID))
	if !p.Simple {
		sb.WriteString("[C]") // complicated as [C]
	}
	sb.WriteString(fmt.Sprintf(" %dmin", p.CostMinutes))
	return sb.String()
}

func Digests(records []note.Record) string {
	difficultiesToProblems := make(map[note.Difficulty][]Problem)
	for _, r := range records {
		difficultiesToProblems[r.Difficulty] = append(difficultiesToProblems[r.Difficulty], NewProblem(r))
	}

	var ret []string
	ret = append(ret, fmt.Sprintf("Easy %d\n", len(difficultiesToProblems[note.Easy])))
	var sb strings.Builder
	for _, p := range difficultiesToProblems[note.Easy] {
		sb.WriteString(fmt.Sprintf("\t%s", p.String()))
	}
	ret = append(ret, sb.String())
	ret = append(ret, fmt.Sprintf("Medium %d\n", len(difficultiesToProblems[note.Medium])))
	sb.Reset()
	for _, p := range difficultiesToProblems[note.Medium] {
		sb.WriteString(fmt.Sprintf("\t%s", p.String()))
	}
	ret = append(ret, sb.String())
	return strings.Join(ret, "\n")
}
