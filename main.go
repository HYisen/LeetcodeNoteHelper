package main

import (
	"fmt"
	"leetcodeNoteHelper/date"
	"leetcodeNoteHelper/note"
	"log"
	"os"
)

func main() {
	if err := Handle(date.Now()); err != nil {
		log.Fatal(err)
	}
}

func Handle(d date.Date) error {
	file, err := os.ReadFile(note.FilePath(d))
	if err != nil {
		return err
	}
	records, err := note.ParseFile(file, d)
	if err != nil {
		return err
	}
	fmt.Println(records)
	return nil
}
