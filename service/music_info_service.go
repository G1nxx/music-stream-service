package service

import (
	"log/slog"
	resp "music-stream-service/service/dtos/response"
	"music-stream-service/service/interfaces"
	mpr "music-stream-service/service/mappers/response"
	"music-stream-service/service/repository"
)

type MusicInfoService struct {
	interfaces.MusicInfoActivity
	repo repository.MusicInfoRepository
	Log  *slog.Logger
}

func NewMusicInfoService(repo repository.MusicInfoRepository, sl *slog.Logger) (*MusicInfoService, error) {
	serv := &MusicInfoService{
		repo: repo,
		Log:  sl,
	}
	sl.Debug("user service successfully initiated")
	return serv, nil
}

func (mi *MusicInfoService) GetAllPlaylistsInList(userId int64) ([]resp.PlaylistInListModel, error) {
	mi.Log.Debug("Getting all playlists of user",
		"userId", userId,
	)

	plsts, err := mi.repo.GetAllPlaylists(userId)
	if err != nil {
		mi.Log.Error("Failed to get user playlists",
			"userId", userId,
			"error", err,
		)
		return nil, err
	}

	mi.Log.Info("Playlists of user collected",
		"userId", userId,
		"count", len(plsts),
	)

	models := make([]resp.PlaylistInListModel, 0, len(plsts))

	for _, el := range plsts {
		plst, err := mpr.ToPlaylistInListModel(el)
		if err != nil {
			mi.Log.Debug("The playlist of user is not correct",
				"userId", userId,
				"id", el.ID,
				"error", err,
			)
			continue
		}
		models = append(models, *plst)
	}

	mi.Log.Info("Playlists of user sent",
		"userId", userId,
		"count", len(models),
	)
	return models, nil
}

func (mi *MusicInfoService) GetAllAlbumsInList(userId int64) ([]resp.AlbumInListModel, error) {
	mi.Log.Debug("Getting all albums of user",
		"userId", userId,
	)

	albs, err := mi.repo.GetAllAlbums(userId)
	if err != nil {
		mi.Log.Error("Failed to get user albums",
			"userId", userId,
			"error", err,
		)
		return nil, err
	}

	mi.Log.Info("Albums of user collected",
		"userId", userId,
		"count", len(albs),
	)

	models := make([]resp.AlbumInListModel, 0, len(albs))

	for _, el := range albs {
		alb, err := mpr.ToAlbumInListModel(el)
		if err != nil {
			mi.Log.Debug("The album of user is not correct",
				"userId", userId,
				"id", el.ID,
				"error", err,
			)
			continue
		}
		models = append(models, *alb)
	}

	mi.Log.Info("Albums of user sent",
		"userId", userId,
		"count", len(models),
	)
	return models, nil
}

func (mi *MusicInfoService) GetAllArtistsInList(userId int64) ([]resp.ArtistModel, error) {
	mi.Log.Debug("Getting all artists of user",
		"userId", userId,
	)

	artsts, err := mi.repo.GetAllArtists(userId)
	if err != nil {
		mi.Log.Error("Failed to get user artists",
			"userId", userId,
			"error", err,
		)
		return nil, err
	}

	mi.Log.Info("Artists of user collected",
		"userId", userId,
		"count", len(artsts),
	)

	models := make([]resp.ArtistModel, 0, len(artsts))

	for _, el := range artsts {
		artist, err := mpr.ToArtistModel(el)
		if err != nil {
			mi.Log.Debug("The artist of user is not correct",
				"userId", userId,
				"id", el.ID,
				"error", err,
			)
			continue
		}
		models = append(models, *artist)
	}

	mi.Log.Info("Artists of user sent",
		"userId", userId,
		"count", len(models),
	)
	return models, nil
}

func (mi *MusicInfoService) GetTracksFromAlbum(albumId int64) ([]resp.TrackInSubsModel, error) {
	mi.Log.Debug("Getting all tracks of album",
		"albumId", albumId,
	)

	trks, err := mi.repo.GetTracksFromAlbum(albumId)
	if err != nil {
		mi.Log.Error("Failed to get tracks of album",
			"albumId", albumId,
			"error", err,
		)
		return nil, err
	}

	mi.Log.Info("Tracks of album collected",
		"albumId", albumId,
		"count", len(trks),
	)

	models := make([]resp.TrackInSubsModel, 0, len(trks))

	for _, el := range trks {
		alb, err := mpr.ToTrackInSubsModel(el)
		if err != nil {
			mi.Log.Debug("The track for album is not correct",
				"albumId", albumId,
				"id", el.ID,
				"error", err,
			)
			continue
		}
		models = append(models, *alb)
	}

	mi.Log.Info("Tracks of album sent",
		"albumId", albumId,
		"count", len(models),
	)
	return models, nil
}

func (mi *MusicInfoService) GetAlbumInfo(albumId int64) (*resp.AlbumViewModel, error) {
	mi.Log.Debug("Getting album info",
		"albumId", albumId,
	)

	alb, err := mi.repo.GetAlbum(albumId)
	if err != nil {
		mi.Log.Error("Failed to get album info",
			"albumId", albumId,
			"error", err,
		)
		return nil, err
	}

	mi.Log.Info("Album info collected",
		"albumId", albumId,
	)

	model, err := mpr.ToAlbumViewModel(*alb)
	if err != nil {
		mi.Log.Debug("The album info is not correct",
			"albumId", albumId,
			"error", err,
		)
	}

	mi.Log.Info("Album info sent",
		"albumId", albumId,
	)
	return model, nil
}

func (mi *MusicInfoService) GetTracksFromPlaylist(playlistId int64) ([]resp.TrackInSubsModel, error) {
	mi.Log.Debug("Getting all tracks of playlist",
		"playlistId", playlistId,
	)

	trks, err := mi.repo.GetTracksFromPlaylist(playlistId)
	if err != nil {
		mi.Log.Error("Failed to get tracks of playlist",
			"playlistId", playlistId,
			"error", err,
		)
		return nil, err
	}

	mi.Log.Info("Tracks of playlist collected",
		"playlistId", playlistId,
		"count", len(trks),
	)

	models := make([]resp.TrackInSubsModel, 0, len(trks))

	for _, el := range trks {
		alb, err := mpr.ToTrackInSubsModel(el)
		if err != nil {
			mi.Log.Debug("The track of playlist is not correct",
				"playlistId", playlistId,
				"id", el.ID,
				"error", err,
			)
			continue
		}
		models = append(models, *alb)
	}

	mi.Log.Info("Tracks of playlist sent",
		"playlistId", playlistId,
		"count", len(models),
	)
	return models, nil
}

func (mi *MusicInfoService) GetPlaylistInfo(playlistId int64) (*resp.PlaylistViewModel, error) {
	mi.Log.Debug("Getting playlist info",
		"playlistId", playlistId,
	)

	plst, err := mi.repo.GetPlaylist(playlistId)
	if err != nil {
		mi.Log.Error("Failed to get playlist info",
			"playlistId", playlistId,
			"error", err,
		)
		return nil, err
	}

	mi.Log.Info("Playlist info collected",
		"playlistId", playlistId,
	)

	model, err := mpr.ToPlaylistViewModel(*plst)
	if err != nil {
		mi.Log.Debug("The playlist info is not correct",
			"playlistId", playlistId,
			"error", err,
		)
	}

	mi.Log.Info("Playlist info sent",
		"playlistId", playlistId,
	)
	return model, nil
}

func (mi *MusicInfoService) GetPlaylistSaves(playlistId int64) (int64, error) {
	mi.Log.Debug("Getting playlist saves",
		"playlistId", playlistId,
	)

	saves, err := mi.repo.GetPlaylistSaves(playlistId)
	if err != nil {
		mi.Log.Error("Failed to get playlist saves",
			"playlistId", playlistId,
			"error", err,
		)
		return 0, err
	}

	mi.Log.Info("Playlist saves collected and sent",
		"playlistId", playlistId,
	)
	return saves, nil
}

func (mi *MusicInfoService) GetReleasesFromArtist(artistId int64) ([]resp.AlbumInListModel, error) {
	mi.Log.Debug("Getting all releases of artist",
		"userId", artistId,
	)

	albs, err := mi.repo.GetReleasesFromArtist(artistId)
	if err != nil {
		mi.Log.Error("Failed to get artist releases",
			"artistId", artistId,
			"error", err,
		)
		return nil, err
	}

	mi.Log.Info("Releases of artist collected",
		"artistId", artistId,
		"count", len(albs),
	)

	models := make([]resp.AlbumInListModel, 0, len(albs))

	for _, el := range albs {
		alb, err := mpr.ToAlbumInListModel(el)
		if err != nil {
			mi.Log.Debug("The release of artist is not correct",
				"artistId", artistId,
				"releaseId", el.ID,
				"error", err,
			)
			continue
		}
		models = append(models, *alb)
	}

	mi.Log.Info("Releases of artist sent",
		"artistId", artistId,
		"count", len(models),
	)
	return models, nil
}

func (mi *MusicInfoService) GetArtistInfo(artistId int64) (*resp.ArtistModel, error) {
	mi.Log.Debug("Getting artist info",
		"artistId", artistId,
	)

	artst, err := mi.repo.GetArtist(artistId)
	if err != nil {
		mi.Log.Error("Failed to get artist info",
			"artistId", artistId,
			"error", err,
		)
		return nil, err
	}

	mi.Log.Info("Playlist info collected",
		"artistId", artistId,
	)

	model, err := mpr.ToArtistModel(*artst)
	if err != nil {
		mi.Log.Debug("The artist info is not correct",
			"artistId", artistId,
			"error", err,
		)
	}

	mi.Log.Info("Artist info sent",
		"artistId", artistId,
	)
	return model, nil
}

func (mi *MusicInfoService) GetArtistAttachment(artistId int64) (*resp.PlaylistViewModel, error) {
	mi.Log.Debug("Getting artist attachment",
		"artistId", artistId,
	)

	plstID, err := mi.repo.GetArtistAttachmentId(artistId)
	if err != nil {
		mi.Log.Error("Failed to get attached playlist id",
			"artistId", artistId,
			"error", err,
		)
		return nil, err
	}

	mi.Log.Info("Attached playlist id collected",
		"artistId", artistId,
		"playlistId", plstID,
	)

	plst, err := mi.repo.GetPlaylist(plstID)
	if err != nil {
		mi.Log.Error("Failed to get attached playlist info",
			"artistId", artistId,
			"playlistId", plst.ID,
			"error", err,
		)
		return nil, err
	}

	mi.Log.Info("Attached playlist info collected",
		"artistId", artistId,
		"playlistId", plst.ID,
	)

	model, err := mpr.ToPlaylistViewModel(*plst)
	if err != nil {
		mi.Log.Debug("The attached playlist info is not correct",
			"artistId", artistId,
			"playlistId", plst.ID,
			"error", err,
		)
	}

	mi.Log.Info("Attached playlist info sent",
		"artistId", artistId,
		"playlistId", plst.ID,
	)
	return model, nil
}

func (mi *MusicInfoService) GetLikedSongs(userId int64) (*resp.PlaylistViewModel, error) {
	mi.Log.Debug("Getting liked songs playlist",
		"userId", userId,
	)

	plstID, err := mi.repo.GetLikedSongsId(userId)
	if err != nil {
		mi.Log.Error("Failed to get liked songs playlist id",
			"userId", userId,
			"error", err,
		)
		return nil, err
	}

	mi.Log.Info("Liked songs playlist id collected",
		"userId", userId,
		"playlistId", plstID,
	)

	plst, err := mi.repo.GetPlaylist(plstID)
	if err != nil {
		mi.Log.Error("Failed to get liked songs playlist info",
			"userId", userId,
			"playlistId", plst.ID,
			"error", err,
		)
		return nil, err
	}

	mi.Log.Info("Liked songs playlist info collected",
		"userId", userId,
		"playlistId", plst.ID,
	)

	model, err := mpr.ToPlaylistViewModel(*plst)
	if err != nil {
		mi.Log.Debug("The liked songs playlist info is not correct",
			"userId", userId,
			"playlistId", plst.ID,
			"error", err,
		)
	}

	mi.Log.Info("Liked songs playlist info sent",
		"userId", userId,
		"playlistId", plst.ID,
	)
	return model, nil
}

func (mi *MusicInfoService) GetIsFollowedArtist(uId, cId int64) (bool, error) {
	mi.Log.Debug("Getting is followed artist",
		"userId", uId,
		"artistId", cId,
	)

	isFoloving, err := mi.repo.GetIsFollowedArtist(uId, cId)
	if err != nil {
		mi.Log.Error("Failed to get is followed artist",
			"userId", uId,
			"artistId", cId,
			"error", err,
		)
		return false, err
	}

	mi.Log.Info("Is followed artist collected and sent",
		"userId", uId,
		"artistId", cId,
		"isFollowing", isFoloving,
	)

	return isFoloving, nil
}

func (mi *MusicInfoService) GetIsFollowedAlbum(uId, cId int64) (bool, error) {
	mi.Log.Debug("Getting is followed album",
		"userId", uId,
		"albumId", cId,
	)

	isFoloving, err := mi.repo.GetIsFollowedAlbum(uId, cId)
	if err != nil {
		mi.Log.Error("Failed to get is followed album",
			"userId", uId,
			"albumId", cId,
			"error", err,
		)
		return false, err
	}

	mi.Log.Info("Is followed album collected and sent",
		"userId", uId,
		"albumId", cId,
		"isFollowing", isFoloving,
	)

	return isFoloving, nil
}

func (mi *MusicInfoService) GetIsFollowedPlaylist(uId, cId int64) (bool, error) {
	mi.Log.Debug("Getting is followed playlist",
		"userId", uId,
		"playlistId", cId,
	)

	isFoloving, err := mi.repo.GetIsFollowedPlaylist(uId, cId)
	if err != nil {
		mi.Log.Error("Failed to get is followed playlist",
			"userId", uId,
			"playlistId", cId,
			"error", err,
		)
		return false, err
	}

	mi.Log.Info("Is followed playlist collected and sent",
		"userId", uId,
		"playlistId", cId,
		"isFollowing", isFoloving,
	)

	return isFoloving, nil
}
