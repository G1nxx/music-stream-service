package response

import (
	e "music-stream-service/domain/entities"
	"time"
)

type PlaylistModel struct {
	ID        int64  `json:"id" binding:"required"`
	Name      string `json:"title" binding:"required"`
	CreatorID int64  `json:"creator_id" binding:"required"`
	CoverPath string `json:"artwork"`
}

func NewPlaylistModel(playlist *e.Playlist, options ...func(*PlaylistModel) (*PlaylistModel, error)) (*PlaylistModel, error) {
	plst := &PlaylistModel{}

	plst.ID = playlist.ID
	plst.Name = playlist.Name
	plst.CreatorID = playlist.CreatorID
	plst.CoverPath = playlist.CoverPath

	for _, opt := range options {
		opt(plst)
	}

	return plst, nil
}

type PlaylistInListModel struct {
	PlaylistModel
	CreatorName string `json:"subtitle" binding:"required"`
	Type        string `json:"type" json-default:"Playlist"`
}

func NewPlaylistInListModel(playlist *e.Playlist, options ...func(*PlaylistInListModel) (*PlaylistInListModel, error)) (*PlaylistInListModel, error) {
	plst := &PlaylistInListModel{}

	plst.ID = playlist.ID
	plst.Name = playlist.Name
	plst.CreatorID = playlist.CreatorID
	plst.CoverPath = playlist.CoverPath
	plst.CreatorName = playlist.CreatorName
	plst.Type = "Playlist"

	for _, opt := range options {
		opt(plst)
	}

	return plst, nil
}

type PlaylistViewModel struct {
	PlaylistModel
	CreatorName  string    `json:"subtitle" binding:"required"`
	Length       int32     `json:"length" binding:"required"`
	CreationTime time.Time `json:"date" binding:"required"`
	Type         string    `json:"type" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	AttachedTo   int64     `json:"attached_to" binding:"required"`
	Saves        int64     `json:"saves"`
}

func NewPlaylistViewModel(playlist *e.Playlist, options ...func(*PlaylistViewModel) (*PlaylistViewModel, error)) (*PlaylistViewModel, error) {
	plst := &PlaylistViewModel{}

	plst.ID = playlist.ID
	plst.Name = playlist.Name
	plst.CreatorID = playlist.CreatorID
	plst.CoverPath = playlist.CoverPath
	plst.CreatorName = playlist.CreatorName
	plst.Length = playlist.Length
	plst.CreationTime = playlist.CreationTime
	plst.Type = "Playlist"
	plst.Description = playlist.Description
	plst.AttachedTo = playlist.AttachedTo
	plst.Saves = 0

	for _, opt := range options {
		opt(plst)
	}

	return plst, nil
}
