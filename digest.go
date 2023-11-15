package main

import (
	"leetcodeNoteHelper/date"
	"leetcodeNoteHelper/diary"
	"leetcodeNoteHelper/note"
	"log"
	"os"
)

type DigestFunc func(now date.Date) (string, error)

func CreateDigester(monthMode bool) DigestFunc {
	if monthMode {
		log.Println("monthly")
		return MonthlyDigest
	}
	return DailyDigest
}

func MonthlyDigest(_ date.Date) (string, error) {
	panic("implement")
}

func DailyDigest(d date.Date) (string, error) {
	file, err := os.ReadFile(note.FilePath(d))
	if err != nil {
		return "", err
	}
	records, err := note.ParseFile(file, d)
	if err != nil {
		return "", err
	}
	return diary.Digest(records), nil
}
