package service

import (
	"log/slog"
	"music-stream-service/service/interfaces"
	"music-stream-service/service/repository"
)

type UserService struct {
	interfaces.UserActivity
	repo repository.UserRepository
	Log  *slog.Logger
}

func NewUserService(repo repository.UserRepository, sl *slog.Logger) (*UserService, error) {
	serv := &UserService{
		repo: repo,
		Log:  sl,
	}
	sl.Debug("user service successfully initiated")
	return serv, nil
}


func (serv *UserService) SubscribeToUser(first_userId, second_userId int64) error {
	return serv.repo.SubscribeToUser(first_userId, second_userId)
}

func (serv *UserService) UnsubscribeFromUser(first_userId, second_userId int64) error {
	return serv.repo.UnsubscribeFromUser(first_userId, second_userId)
}

func (serv *UserService) SubscribeToAlbum(userId, albumId int64) error {
	return serv.repo.SubscribeToAlbum(userId, albumId)
}

func (serv *UserService) UnsubscribeFromAlbum(userId, albumId int64) error {
	return serv.repo.UnsubscribeFromAlbum(userId, albumId)
}

func (serv *UserService)SubscribeToPlaylist(userId, playlistId int64) error {
	return serv.repo.SubscribeToPlaylist(userId, playlistId)
}

func (serv *UserService) UnsubscribeFromPlaylist(userId, playlistId int64) error {
	return serv.repo.UnsubscribeFromPlaylist(userId, playlistId)
}

func (serv *UserService) AddTrackToPlaylist(trackId, playlistId int64) error {
	return serv.repo.AddTrackToPlaylist(trackId, playlistId)
}

func (serv *UserService) RemoveTrackfromPlaylist(trackId, playlistId int64) error {
	return serv.repo.RemoveTrackfromPlaylist(trackId, playlistId)
}
