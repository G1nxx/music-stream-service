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
	serv.Log.Debug("Getting to subscribe to user",
		"first_userId", first_userId,
		"second_userId", second_userId,
	)
	err := serv.repo.SubscribeToUser(first_userId, second_userId)
	if err != nil {
		serv.Log.Info("Subscription failed",
			"first_userId", first_userId,
			"second_userId", second_userId,
			"error", err.Error(),
		)
		return err
	}
	serv.Log.Info("Subscription completed succseessfull",
		"first_userId", first_userId,
		"second_userId", second_userId,
	)
	return nil
}

func (serv *UserService) UnsubscribeFromUser(first_userId, second_userId int64) error {
	serv.Log.Debug("Getting to unsubscribe from user",
		"first_userId", first_userId,
		"second_userId", second_userId,
	)
	err := serv.repo.UnsubscribeFromUser(first_userId, second_userId)
	if err != nil {
		serv.Log.Info("Unsubscription failed",
			"first_userId", first_userId,
			"second_userId", second_userId,
			"error", err.Error(),
		)
		return err
	}
	serv.Log.Info("Unsubscription completed succseessfull",
		"first_userId", first_userId,
		"second_userId", second_userId,
	)
	return nil
}

func (serv *UserService) SubscribeToAlbum(userId, albumId int64) error {
	serv.Log.Debug("Getting to subscribe to album",
		"userId", userId,
		"albumId", albumId,
	)
	err := serv.repo.SubscribeToAlbum(userId, albumId)
	if err != nil {
		serv.Log.Info("Subscription failed",
			"userId", userId,
			"albumId", albumId,
			"error", err.Error(),
		)
		return err
	}
	serv.Log.Info("Subscription completed succseessfull",
		"userId", userId,
		"albumId", albumId,
	)
	return nil
}

func (serv *UserService) UnsubscribeFromAlbum(userId, albumId int64) error {
	serv.Log.Debug("Getting to unsubscribe from album",
		"userId", userId,
		"albumId", albumId,
	)
	err := serv.repo.UnsubscribeFromAlbum(userId, albumId)
	if err != nil {
		serv.Log.Info("Unubscription failed",
			"userId", userId,
			"albumId", albumId,
			"error", err.Error(),
		)
		return err
	}
	serv.Log.Info("Unubscription completed succseessfull",
		"userId", userId,
		"albumId", albumId,
	)
	return nil
}

func (serv *UserService) SubscribeToPlaylist(userId, playlistId int64) error {
	serv.Log.Debug("Getting to subscribe to album",
		"userId", userId,
		"playlistId", playlistId,
	)
	err := serv.repo.SubscribeToPlaylist(userId, playlistId)
	if err != nil {
		serv.Log.Info("Subscription failed",
			"userId", userId,
			"playlistId", playlistId,
			"error", err.Error(),
		)
		return err
	}
	serv.Log.Info("Subscription completed succseessfull",
		"userId", userId,
		"playlistId", playlistId,
	)
	return nil
}

func (serv *UserService) UnsubscribeFromPlaylist(userId, playlistId int64) error {
	serv.Log.Debug("Getting to unsubscribe from album",
		"userId", userId,
		"playlistId", playlistId,
	)
	err := serv.repo.UnsubscribeFromPlaylist(userId, playlistId)
	if err != nil {
		serv.Log.Info("Unsubscription failed",
			"userId", userId,
			"playlistId", playlistId,
			"error", err.Error(),
		)
		return err
	}
	serv.Log.Info("Unsubscription completed succseessfull",
		"userId", userId,
		"playlistId", playlistId,
	)
	return nil
}

func (serv *UserService) AddTrackToPlaylist(trackId, playlistId int64) error {
	return serv.repo.AddTrackToPlaylist(trackId, playlistId)
}

func (serv *UserService) RemoveTrackfromPlaylist(trackId, playlistId int64) error {
	return serv.repo.RemoveTrackfromPlaylist(trackId, playlistId)
}
