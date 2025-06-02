package repository

import (
	"log/slog"
	"music-stream-service/domain/usecases"
)

type Repository struct {
	AuthRepo      AuthorizationRepository
	UserRepo      UserRepository
	S3Storage     S3Storage
	RedisStorage  RedisStorage
	MusicInfoRepo MusicInfoRepository
	Log           *slog.Logger
}

type AuthorizationRepository interface {
	usecases.Authorization
}

type UserRepository interface {
	usecases.User
}

type S3Storage interface {
	usecases.S3
}

type RedisStorage interface {
	usecases.Redis
}

type MusicInfoRepository interface {
	usecases.MusicInfo
}
