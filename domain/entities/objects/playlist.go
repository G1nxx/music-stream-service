package objects

import "time"

type Playlist struct {
	ID			  int64
	Name          string
	Discription   string
	CreatorID     int64
	CreationTime  time.Time
	TrackIDs      []int64
	SubscriberIDs []int64
}
