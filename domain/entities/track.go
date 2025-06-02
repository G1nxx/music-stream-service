package entities

import "time"

type Track struct {
	ID            int64
	Name          string
	AlbumID       int64
	AlbumName     string
	AuthorID      int64
	AuthorName    string
	Genre         string
	Length        int32
	RecordingDate time.Time
	SizeInBytes   int32
	Number        int16
	Format        string
	Path          string
	CoverPath     string
}
