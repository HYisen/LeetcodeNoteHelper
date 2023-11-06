package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"leetcodeNoteHelper/date"
	"leetcodeNoteHelper/diary"
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
	digest := diary.Digest(records)
	fmt.Println(digest)
	// The origin design is just append to the markdown diary file.
	// But I found my editor Typora is slow to react and start a race condition in normal work flow.
	// So switch to along with the standard out display, copy the text to clipboard.
	// Another solution is use Unix philosophy as pipe the output to tools such as gclip,
	// which could be worse in my imagination on Windows as PowerShell scripts shall not run through a double click.
	return writeToClipboard(digest)
}

func writeToClipboard(s string) error {
	if err := clipboard.Init(); err != nil {
		return err
	}

	data := []byte(s)
	log.Printf("going to write %d bytes to clipboard", len(data))
	clipboard.Write(clipboard.FmtText, data)
	// The return channel is not monitored to check whether written data is overwritten by others.
	// As it's nothing but a best-effort design.
	return nil
}
