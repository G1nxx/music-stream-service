package objects

import "time"

type Album struct {
	ID            int64
	Name          string
	Author        string
	Genre         string
	Length        int32
	RecordingDate time.Time
	CoverPath     string
	Type          string
	TrackIDs      []int64
	SubscriberIDs []int64
}
