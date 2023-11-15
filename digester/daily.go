package digester

import (
	"leetcodeNoteHelper/date"
	"leetcodeNoteHelper/diary"
	"leetcodeNoteHelper/note"
	"os"
)

type Daily struct {
}

func NewDaily() *Daily {
	return &Daily{}
}

func (d *Daily) Digest(now date.Date) (string, error) {
	file, err := os.ReadFile(note.FilePath(now.YearMonth()))
	if err != nil {
		return "", err
	}
	records, err := note.ParseFile(file, now)
	if err != nil {
		return "", err
	}
	return diary.Digest(records), nil
}
