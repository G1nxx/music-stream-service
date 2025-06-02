package interfaces

import (
	resp "music-stream-service/service/dtos/response"
)

type MusicInfoActivity interface {
	GetAllPlaylistsInList(userId int64) ([]resp.PlaylistInListModel, error)
	GetAllAlbumsInList(userId int64) ([]resp.AlbumInListModel, error)
	GetAllArtistsInList(userId int64) ([]resp.ArtistModel, error)

	GetTracksFromAlbum(albumId int64) ([]resp.TrackInSubsModel, error)
	GetAlbumInfo(albumId int64) (*resp.AlbumViewModel, error)

	GetTracksFromPlaylist(playlistId int64) ([]resp.TrackInSubsModel, error)
	GetPlaylistInfo(playlistId int64) (*resp.PlaylistViewModel, error)
	GetPlaylistSaves(playlistId int64) (int64, error)

	GetReleasesFromArtist(artistId int64) ([]resp.AlbumInListModel, error)
	GetArtistInfo(artistId int64) (*resp.ArtistModel, error)
	GetArtistAttachment(artistId int64) (*resp.PlaylistViewModel, error)
}
