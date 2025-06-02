package entities

import "time"

type Album struct {
	ID            int64
	Name          string
	AuthorID      int64
	AuthorName    string
	Genre         string
	Length        int32
	RecordingDate time.Time
	CoverPath     string
	Type          string
}
