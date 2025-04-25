package objects

import "time"

type Track struct {
	ID            int64
	Name          string
	AlbumID       int64
	AuthorID      int64
	Genre         string
	Length        int32
	RecordingDate time.Time
	SizeInBytes   int32
	Number        int8
	Format        string
	Path          string
	CoverPath     string
}