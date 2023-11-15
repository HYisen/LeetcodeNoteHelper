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
		{"multi-line format", args{
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
