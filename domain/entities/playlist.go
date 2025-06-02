package entities

import "time"

type Playlist struct {
	ID           int64
	Name         string
	CreatorID    int64
	CreatorName  string
	Length       int32
	CreationTime time.Time
	CoverPath    string
	AttachedTo   int64
	Description  string
}
