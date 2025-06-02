package response

import (
	e "music-stream-service/domain/entities"
	"time"
)

type AlbumModel struct {
	ID        int64  `json:"id" binding:"required"`
	Name      string `json:"title" binding:"required"`
	AuthorID  int64  `json:"author_id" binding:"required"`
	CoverPath string `json:"artwork" binding:"required"`
	Type      string `json:"type" binding:"required"`
}

func NewAlbumModel(album *e.Album, options ...func(*AlbumModel) (*AlbumModel, error)) (*AlbumModel, error) {
	alb := &AlbumModel{}

	alb.ID = album.ID
	alb.Name = album.Name
	alb.AuthorID = album.AuthorID
	alb.CoverPath = album.CoverPath
	alb.Type = album.Type

	for _, opt := range options {
		opt(alb)
	}

	return alb, nil
}

type AlbumInListModel struct {
	AlbumModel
	AuthorName string `json:"subtitle" binding:"required"`
}

func NewAlbumInListModel(album *e.Album, options ...func(*AlbumInListModel) (*AlbumInListModel, error)) (*AlbumInListModel, error) {
	alb := &AlbumInListModel{}

	alb.ID = album.ID
	alb.Name = album.Name
	alb.AuthorID = album.AuthorID
	alb.CoverPath = album.CoverPath
	alb.Type = album.Type
	alb.AuthorName = album.AuthorName

	for _, opt := range options {
		opt(alb)
	}

	return alb, nil
}

type AlbumViewModel struct {
	AlbumModel
	AuthorName    string    `json:"subtitle" binding:"required"`
	Length        int32     `json:"length" binding:"required"`
	RecordingDate time.Time `json:"date" binding:"required"`
}

func NewAlbumViewModel(album *e.Album, options ...func(*AlbumViewModel) (*AlbumViewModel, error)) (*AlbumViewModel, error) {
	alb := &AlbumViewModel{}

	alb.ID = album.ID
	alb.Name = album.Name
	alb.AuthorID = album.AuthorID
	alb.CoverPath = album.CoverPath
	alb.Type = album.Type
	alb.AuthorName = album.AuthorName
	alb.Length = album.Length
	alb.RecordingDate = album.RecordingDate

	for _, opt := range options {
		opt(alb)
	}

	return alb, nil
}
