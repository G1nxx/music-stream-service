package request

type InfoModel struct {
	UserID int64 `json:"user_id" binding:"required"`
	Filter string `json:"filter" default:"All"`
}

type TracksFromAlbumModel struct {
	AlbumId string `json:"album_id" binding:"required"`
}

type TracksFromPlaylistModel struct {
	PlaylistId string `json:"playlist_id" binding:"required"`
}

type ReleasesFromArtistModel struct {
	ArtistId string `json:"artist_id" binding:"required"`
}

type LikedSongsModel struct {	
	UserId string `json:"user_id" binding:"required"`
}

type FollowStatusModel struct {	
	UserId string `json:"user_id" binding:"required"`
	ContentId string `json:"content_id" binding:"required"`
	Type string `json:"type" binding:"required"`
}
