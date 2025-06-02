package request

type UserAuthModel struct {
	Login    string `json:"login"    binding:"required"`
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSubsToUserModel struct {
	FirstUserID  int64 `json:"first_user_id" binding:"required"`
	SecondUserID int64 `json:"second_user_id" binding:"required"`
}

type UserSubsToAlbumModel struct {
	UserID  int64 `json:"user_id" binding:"required"`
	AlbumID int64 `json:"album_id" binding:"required"`
}

type UserSubsToPlaylistModel struct {
	UserID     int64 `json:"user_id" binding:"required"`
	PlaylistID int64 `json:"playlist_id" binding:"required"`
}

type TrackAddsToPlaylistModel struct {
	TrackID    int64 `json:"track_id" binding:"required"`
	PlaylistID int64 `json:"playlist_id" binding:"required"`
}
