package diary

import (
	"fmt"
	"leetcodeNoteHelper/note"
	"strconv"
	"strings"
	"text/template"
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

type Blob struct {
	Difficulty       note.Difficulty
	ProblemCount     int
	TotalCostMinutes int
	Problems         []Problem
}

var tmpl = template.Must(template.ParseFiles("diary/digest.md.tmpl"))

func Digests(records []note.Record) string {
	difficultiesToProblems := make(map[note.Difficulty][]Problem)
	for _, r := range records {
		difficultiesToProblems[r.Difficulty] = append(difficultiesToProblems[r.Difficulty], NewProblem(r))
	}

	var sb strings.Builder
	sb.WriteString("#### leetcode digest\n\n")
	for _, difficulty := range []note.Difficulty{note.Easy, note.Medium, note.Hard} {
		problems := difficultiesToProblems[difficulty]
		if len(problems) == 0 {
			continue
		}
		var totalCostMinutes int
		for _, problem := range problems {
			totalCostMinutes += problem.CostMinutes
		}
		b := Blob{
			Difficulty:       difficulty,
			ProblemCount:     len(problems),
			TotalCostMinutes: totalCostMinutes,
			Problems:         problems,
		}

		if err := tmpl.Execute(&sb, b); err != nil {
			panic(err)
		}
	}
	return sb.String()
}
