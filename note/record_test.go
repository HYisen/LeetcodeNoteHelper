package note

import (
	"leetcodeNoteHelper/date"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestNewRecord(t *testing.T) {
	type args struct {
		lines []string
		d     date.Date
	}
	tests := []struct {
		name    string
		args    args
		want    *Record
		wantErr bool
	}{
		{"multi-line no begin", args{
			lines: strings.Split("## 0904\n\nM 1136\n\nTLE 71/91\n\n1255 hacked", "\n\n"),
			d: date.Date{
				Year:  2020,
				Month: time.January,
				Day:   1,
			},
		}, &Record{
			ProblemID:  904,
			Difficulty: Medium,
			Simple:     false,
			Begin:      time.Date(2020, time.January, 1, 11, 36, 0, 0, time.Local),
			End:        time.Date(2020, time.January, 1, 12, 55, 0, 0, time.Local),
		}, false},
		{"multi-line not forward", args{
			lines: []string{
				"## 0898",
				"M 2352",
				"2357 TLE 70/83 brute force no other idea",
				"80/83 optimization one",
				"0011",
			},
			d: date.Date{
				Year:  2020,
				Month: time.January,
				Day:   1,
			},
		}, &Record{
			ProblemID:  898,
			Difficulty: Medium,
			Simple:     false,
			Begin:      time.Date(2020, time.January, 1, 23, 52, 0, 0, time.Local),
			End:        time.Date(2020, time.January, 2, 0, 11, 0, 0, time.Local),
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRecord(tt.args.lines, tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}
