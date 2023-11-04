package note

import (
	"bufio"
	"bytes"
	"fmt"
	"leetcodeNoteHelper/date"
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
	all, err := extractFieldByDate(data, d)
	if err != nil {
		return nil, err
	}

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

func extractFieldByDate(data []byte, d date.Date) ([]string, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var started bool
	dayHeader := fmt.Sprintf("# %02d", d.Day)
	var ret []string
	for scanner.Scan() {
		line := scanner.Text()
		// fast skip to target date
		if !started {
			if line == dayHeader {
				started = true
				ret = append(ret, line)
			}
			continue
		}
		// break when meets next date
		if strings.HasPrefix(line, "# ") {
			break
		}
		ret = append(ret, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}
