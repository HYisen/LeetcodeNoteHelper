package digester

import (
	"fmt"
	"leetcodeNoteHelper/date"
	"leetcodeNoteHelper/note"
	"os"
	"strings"
)

type Monthly struct {
	backtraceMonths int
}

func NewMonthly(backtraceMonths int) *Monthly {
	return &Monthly{backtraceMonths: backtraceMonths}
}

func (m *Monthly) Digest(now date.Date) (string, error) {
	var lines []string
	for i := 0; i < m.backtraceMonths; i++ {
		ym := now.Add(0, -i, 0).YearMonth()
		file, err := os.ReadFile(note.FilePath(ym))
		if err != nil {
			// allow no data in month mode
			if !os.IsNotExist(err) {
				return "", err
			}
		}
		whole, err := note.ParseFileWithFilter(file, ym, 0)
		if err != nil {
			return "", fmt.Errorf("on %v: %v", ym, err)
		}
		lines = append(lines, digest(ym, whole))
	}
	return strings.Join(lines, "\n\n"), nil
}

func digest(ym date.YearMonth, whole [][]note.Record) string {
	difficultyToRecords := make(map[note.Difficulty][]note.Record)
	var hours float64
	for _, records := range whole {
		for _, record := range records {
			difficultyToRecords[record.Difficulty] = append(difficultyToRecords[record.Difficulty], record)
			hours += record.End.Sub(record.Begin).Hours()
		}
	}

	return fmt.Sprintf(
		"%d %s %d|%d|%d %.1f hr",
		ym.Year,
		ym.Month.String()[:3],
		len(difficultyToRecords[note.Easy]),
		len(difficultyToRecords[note.Medium]),
		len(difficultyToRecords[note.Hard]),
		hours,
	)
}
