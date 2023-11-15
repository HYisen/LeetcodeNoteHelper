package digester

import "leetcodeNoteHelper/date"

type Digester interface {
	Digest(now date.Date) (string, error)
}
