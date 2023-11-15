package digester

import (
	"leetcodeNoteHelper/date"
)

type Monthly struct {
	backtraceMonths int
}

func NewMonthly(backtraceMonths int) *Monthly {
	return &Monthly{backtraceMonths: backtraceMonths}
}

func (m *Monthly) Digest(now date.Date) (string, error) {
	panic("implement me")
}
