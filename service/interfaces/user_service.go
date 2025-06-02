package interfaces

type UserActivity interface {
	SubscribeToUser(first_userId, second_userId int64) error
	UnsubscribeFromUser(first_userId, second_userId int64) error

	SubscribeToAlbum(userId, albumId int64) error
	UnsubscribeFromAlbum(userId, albumId int64) error

	SubscribeToPlaylist(userId, playlistId int64) error
	UnsubscribeFromPlaylist(userId, playlistId int64) error

	AddTrackToPlaylist(trackId, playlistId int64) error
	RemoveTrackfromPlaylist(trackId, playlistId int64) error
}