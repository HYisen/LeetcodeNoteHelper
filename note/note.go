package note

import (
	"bufio"
	"bytes"
	"fmt"
	"leetcodeNoteHelper/date"
	"strconv"
	"strings"
)

func FilePath(ym date.YearMonth) string {
	name := fmt.Sprintf("%02d%02d", ym.Year%100, ym.Month)
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

// ParseFileWithFilter scans data and output Record s group by dayOfMonth.
// If optionalFilterDay not zero, only that date's Record s would be return.
func ParseFileWithFilter(data []byte, ym date.YearMonth, optionalFilterDay int) ([][]Record, error) {
	dayToLines, err := extract(data, optionalFilterDay)
	if err != nil {
		return nil, err
	}

	var ret [][]Record
	for day, all := range dayToLines {
		// previous extract function has filtered all key day here satisfies optionalFilterDay criteria.
		records, err := parseLines(ym.Date(day), all)
		if err != nil {
			return nil, err
		}
		ret = append(ret, records)
	}
	return ret, nil
}

func parseLines(d date.Date, all []string) ([]Record, error) {
	var records []Record
	var current []string
	all = dropHeadAndEmpty(all) // Skip first line as date header and drop empty lines in MarkDown.
	for _, line := range all {
		if strings.HasPrefix(line, "## ") && len(current) > 1 {
			r, err := NewRecord(current, d)
			if err != nil {
				return nil, fmt.Errorf("parse record %s: %v", current, err)
			}
			records = append(records, *r)
			current = nil
		}
		current = append(current, line)
	}
	if len(current) > 1 {
		r, err := NewRecord(current, d)
		if err != nil {
			return nil, fmt.Errorf("parse record %s: %v", current, err)
		}
		records = append(records, *r)
	}
	return records, nil
}

func ParseFile(data []byte, d date.Date) ([]Record, error) {
	whole, err := ParseFileWithFilter(data, d.YearMonth(), d.Day)
	if err != nil {
		return nil, err
	}
	// If no target date in data.
	if len(whole) == 0 {
		return nil, nil
	}
	// As d.Day is specified in ParseFileWithFilter, shall be impossible if len(whole) > 1.
	return whole[0], err
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

// extract scans data and output by section days, if optionalFilterDay not zero, dayToLines only has key on that day.
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
