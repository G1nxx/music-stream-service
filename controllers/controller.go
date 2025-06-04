package controllers

import (
	"music-stream-service/controllers/middleware"
	service "music-stream-service/service"
	"net/http"
)

type Controller struct {
	AuthController      AuthController
	UserController      UserController
	MusicInfoController MusicInfoController
}

func (controller *Controller) RegisterRoutes(mux *http.ServeMux) {
	controller.AuthController.registerAuthorization(mux)
	controller.UserController.registerUserActivity(mux)
	controller.MusicInfoController.registerMusicInfo(mux)
}

func (authController *AuthController) registerAuthorization(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/signup", authController.signUp)
}

func (UserController *UserController) registerUserActivity(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/user/subs/to_content", UserController.SubscribeToContent)
	mux.HandleFunc("POST /api/user/subs/from_content", UserController.UnsubscribeFromContent)
	mux.HandleFunc("POST /api/user/subs/track_to_playlist", UserController.AddTrackToPlaylist)
	mux.HandleFunc("POST /api/user/subs/track_from_playlist", UserController.RemoveTrackFromPlaylist)
}

func (MusicInfoController *MusicInfoController) registerMusicInfo(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/user/get_subscriptions", MusicInfoController.GetAllSubscriptionsInListByFilter)
	mux.HandleFunc("POST /api/user/get_album_info", MusicInfoController.GetAlbumInfoByAlbumId)
	mux.HandleFunc("POST /api/user/get_playlist_info", MusicInfoController.GetPlaylistInfoByPlaylistId)
	mux.HandleFunc("POST /api/user/get_artist_info", MusicInfoController.GetArtistInfoByArtistId)
	mux.HandleFunc("POST /api/user/get_liked_songs", MusicInfoController.GetLikedSongsByUserId)
	mux.HandleFunc("POST /api/user/get_follow_status", MusicInfoController.GetFollowStatusByUserId)
}

func NewController(serv *service.Service) *Controller {
	middleware := middleware.NewMiddleware(serv)
	return &Controller{
		AuthController:      *NewAuthController(serv.AuthService, *middleware),
		UserController:      *NewUserController(serv.UserService, *middleware),
		MusicInfoController: *NewMusicInfoController(serv.MusicInfoService, *middleware),
	}
}