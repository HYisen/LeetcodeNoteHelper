package date

import (
	"time"
)

func Now() Date {
	d := time.Now()
	// Let's say the short first period after midnight belongs to previous date.
	if BelongToYesterday(d) {
		// It's not OO but FP. Don't ask my how I figure out it.
		d = d.AddDate(0, 0, -1)
	}
	return Date{
		Year:  d.Year(),
		Month: d.Month(),
		Day:   d.Day(),
	}
}

func BelongToYesterday(t time.Time) bool {
	// I won't get up before 0400 so use the time threshold here.
	return t.Hour() < 4
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}
