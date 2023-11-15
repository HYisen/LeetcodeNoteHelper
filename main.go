package main

import (
	"flag"
	"leetcodeNoteHelper/date"
	"log"
)

var monthMode = flag.Bool("monthMode", false, "enable month mode that analyse to months")

func main() {
	flag.Parse()
	digest := CreateDigester(*monthMode)

	content, err := digest(date.Now())
	if err != nil {
		log.Fatal(err)
	}

	if err := Output(content); err != nil {
		log.Fatal(err)
	}
}
