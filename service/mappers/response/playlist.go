package persistanceMappers

import (
	e "music-stream-service/domain/entities"
	persistance "music-stream-service/service/dtos/response"
)

func ToPlaylistModel(plst e.Playlist) (*persistance.PlaylistModel, error) {
	playlist, err := persistance.NewPlaylistModel(&plst)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func ToPlaylistInListModel(plst e.Playlist) (*persistance.PlaylistInListModel, error) {
	playlist, err := persistance.NewPlaylistInListModel(&plst)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func ToPlaylistViewModel(plst e.Playlist) (*persistance.PlaylistViewModel, error) {
	playlist, err := persistance.NewPlaylistViewModel(&plst)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}