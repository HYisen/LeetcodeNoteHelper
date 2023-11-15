package main

import (
	"flag"
	"leetcodeNoteHelper/date"
	"leetcodeNoteHelper/digester"
	"log"
)

var monthMode = flag.Bool("monthMode", false, "enable month mode that analyse to months")

// I just prefer the default value of backtraceMonths here.
var backtraceMonths = flag.Int("backtraceMonths", 6, "the amount of most recently months focused in monthMode")

func main() {
	flag.Parse()
	d := CreateDigester()

	content, err := d.Digest(date.Now())
	if err != nil {
		log.Fatal(err)
	}

	if err := Output(content); err != nil {
		log.Fatal(err)
	}
}

func CreateDigester() digester.Digester {
	if *monthMode {
		return digester.NewMonthly(*backtraceMonths)
	}
	return digester.NewDaily()
}
