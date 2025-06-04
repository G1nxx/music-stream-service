package usecases

import (
	e "music-stream-service/domain/entities"
	"time"
)

type Authorization interface {
	AddUser(e.User) error
	GetUser(username, password string) (*e.User, error)
	GetUserById(id int64) (*e.User, error)
}

type User interface {
	SubscribeToUser(first_userId, second_userId int64) error
	UnsubscribeFromUser(first_userId, second_userId int64) error

	SubscribeToAlbum(userId, albumId int64) error
	UnsubscribeFromAlbum(userId, albumId int64) error

	SubscribeToPlaylist(userId, playlistId int64) error
	UnsubscribeFromPlaylist(userId, playlistId int64) error

	AddTrackToPlaylist(trackId, playlistId int64) error
	RemoveTrackfromPlaylist(trackId, playlistId int64) error

	//CreatePlaylist()
}

type S3 interface {
	Upload(bucket, fileName string, fileContent []byte) error
	Download(bucket, fileName string) ([]byte, error)
	Delete(bucket, fileName string) error
}

type Redis interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
}

type MusicInfo interface {
	GetAllAlbums(userId int64) ([]e.Album, error)
	GetAllPlaylists(userId int64) ([]e.Playlist, error)
	GetAllArtists(userId int64) ([]e.User, error)

	GetTracksFromAlbum(albumId int64) ([]e.Track, error)
	GetAlbum(albumId int64) (*e.Album, error)
	
	GetTracksFromPlaylist(playlistId int64) ([]e.Track, error)
	GetPlaylist(playlistId int64) (*e.Playlist, error)
	GetPlaylistSaves(playlistId int64) (int64, error)

	GetReleasesFromArtist(artistId int64) ([]e.Album, error)
	GetArtist(artistId int64) (*e.User, error)
	GetArtistAttachmentId(artistId int64) (int64, error)

	GetLikedSongsId(userId int64) (int64, error)
	
	GetIsFollowedArtist(uId, cId int64) (bool, error)
	GetIsFollowedAlbum(uId, cId int64) (bool, error)
	GetIsFollowedPlaylist(uId, cId int64) (bool, error)
}
