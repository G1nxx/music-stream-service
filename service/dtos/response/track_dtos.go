package response

import e "music-stream-service/domain/entities"

type TrackModel struct {
	ID        int64  `json:"id" binding:"required"`
	Name      string `json:"title" binding:"required"`
	AlbumID   int64  `json:"album_id" binding:"required"`
	AuthorID  int64  `json:"author_id" binding:"required"`
	Length    int32  `json:"length" binding:"required"`
	Path      string `json:"url" binding:"required"`
	CoverPath string `json:"artwork" binding:"required"`
}

func NewTrackModel(track *e.Track, options ...func(*TrackModel) (*TrackModel, error)) (*TrackModel, error) {
	trk := &TrackModel{}

	trk.ID = track.ID
	trk.Name = track.Name
	trk.AlbumID = track.AlbumID
	trk.AuthorID = track.AuthorID
	trk.Length = track.Length
	trk.Path = track.Path
	trk.CoverPath = track.CoverPath

	for _, opt := range options {
		opt(trk)
	}

	return trk, nil
}

type TrackInSubsModel struct {
	TrackModel
	Number     int16   `json:"number" binding:"required"`
	AuthorName string `json:"artist" binding:"required"`
}

func NewTrackInAlbumModel(track *e.Track, options ...func(*TrackInSubsModel) (*TrackInSubsModel, error)) (*TrackInSubsModel, error) {
	trk := &TrackInSubsModel{}

	trk.ID = track.ID
	trk.Name = track.Name
	trk.AlbumID = track.AlbumID
	trk.AuthorID = track.AuthorID
	trk.Length = track.Length
	trk.Path = track.Path
	trk.CoverPath = track.CoverPath
	trk.Number = track.Number
	trk.AuthorName = track.AuthorName

	for _, opt := range options {
		opt(trk)
	}

	return trk, nil
}

type TrackPlayingModel struct {
	TrackModel
	AlbumName  string `json:"album_name" binding:"required"`
	AuthorName string `json:"author_name" binding:"required"`
}

func NewTrackPlayingModel(track *e.Track, options ...func(*TrackPlayingModel) (*TrackPlayingModel, error)) (*TrackPlayingModel, error) {
	trk := &TrackPlayingModel{}

	trk.ID = track.ID
	trk.Name = track.Name
	trk.AlbumID = track.AlbumID
	trk.AuthorID = track.AuthorID
	trk.Length = track.Length
	trk.Path = track.Path
	trk.CoverPath = track.CoverPath
	trk.AuthorName = track.AuthorName
	trk.AlbumName = track.AlbumName

	for _, opt := range options {
		opt(trk)
	}

	return trk, nil
}
