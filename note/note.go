package note

import (
	"bufio"
	"bytes"
	"fmt"
	"leetcodeNoteHelper/date"
	"strconv"
	"strings"
)

func FilePath(d date.Date) string {
	name := fmt.Sprintf("%02d%02d", d.Year%100, d.Month)
	return fmt.Sprintf("/Users/hyisen/Library/CloudStorage/OneDrive-Personal/Notes/diary/leetcode/%s.md", name)
}

func dropHeadAndEmpty(lines []string) []string {
	var ret []string
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		ret = append(ret, lines[i])
	}
	return ret
}

func ParseFile(data []byte, d date.Date) ([]Record, error) {
	dayToLines, err := extract(data, d.Day)
	if err != nil {
		return nil, err
	}
	all := dayToLines[d.Day]

	var ret []Record
	var current []string
	all = dropHeadAndEmpty(all) // Skip first line as date header and drop empty lines in MarkDown.
	for _, line := range all {
		if strings.HasPrefix(line, "## ") && len(current) > 1 {
			r, err := NewRecord(current, d)
			if err != nil {
				return nil, fmt.Errorf("parse record %s: %v", current, err)
			}
			ret = append(ret, *r)
			current = nil
		}
		current = append(current, line)
	}
	if len(current) > 1 {
		r, err := NewRecord(current, d)
		if err != nil {
			return nil, fmt.Errorf("parse record %s: %v", current, err)
		}
		ret = append(ret, *r)
	}
	return ret, nil
}

const markdownHeaderPrefix = "# "

func isDayHeader(line string) (day int, ok bool) {
	dayString, ok := strings.CutPrefix(line, markdownHeaderPrefix)
	if !ok {
		return 0, false
	}
	num, err := strconv.Atoi(dayString)
	if err != nil {
		return 0, false
	}
	if !betweenMonthDayRange(num) {
		return 0, false
	}
	return num, true
}

func betweenMonthDayRange(num int) bool {
	return num >= 1 && num <= 31
}

// extract scan data and output by section days, if optionalFilterDay not zero, dayToLines only has key on that day.
func extract(data []byte, optionalFilterDay int) (dayToLines map[int][]string, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	ret := make(map[int][]string)
	var started bool
	var sectionDay int
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, markdownHeaderPrefix) {
			started = false
		}
		if !started {
			day, ok := isDayHeader(line)
			if !ok {
				continue
			}
			if optionalFilterDay != 0 && optionalFilterDay != day {
				continue
			}
			started = true
			sectionDay = day
		}
		if started {
			ret[sectionDay] = append(ret[sectionDay], line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}
