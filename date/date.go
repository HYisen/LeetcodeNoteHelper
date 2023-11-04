package date

import (
	"time"
)

func Now() Date {
	d := time.Now()
	// Let's say the short first period after midnight belongs to previous date.
	// I won't get up before 0400 so use the time threshold here.
	if d.Hour() < 4 {
		// It's not OO but FP. Don't ask my how I figure out it.
		d = d.AddDate(0, 0, -1)
	}
	return Date{
		Year:  d.Year(),
		Month: d.Month(),
		Day:   d.Day(),
	}
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}
