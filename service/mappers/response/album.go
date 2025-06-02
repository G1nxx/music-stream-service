package persistanceMappers

import (
	e "music-stream-service/domain/entities"
	persistance "music-stream-service/service/dtos/response"
)

func ToAlbumModel(alb e.Album) (*persistance.AlbumModel, error) {
	album, err := persistance.NewAlbumModel(&alb)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func ToAlbumInListModel(alb e.Album) (*persistance.AlbumInListModel, error) {
	album, err := persistance.NewAlbumInListModel(&alb)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func ToAlbumViewModel(alb e.Album) (*persistance.AlbumViewModel, error) {
	album, err := persistance.NewAlbumViewModel(&alb)
	if err != nil {
		return nil, err
	}
	return album, nil
}
