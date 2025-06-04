package request

type UserAuthModel struct {
	Login    string `json:"login"    binding:"required"`
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSubsToContentModel struct {
	UserID    string `json:"user_id" binding:"required"`
	ContentID string `json:"content_id" binding:"required"`
	Type      string `json:"type" binding:"required"`
}

type TrackAddsToPlaylistModel struct {
	TrackID    int64 `json:"track_id" binding:"required"`
	PlaylistID int64 `json:"playlist_id" binding:"required"`
}
