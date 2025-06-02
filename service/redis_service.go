package service

import (
	"log/slog"
	"music-stream-service/service/interfaces"
	"music-stream-service/service/repository"
	"time"
)

type RedisService struct {
	interfaces.RedisService
	repo repository.RedisStorage
	Log  *slog.Logger
}

func NewRedisService(repo repository.RedisStorage, sl *slog.Logger) (*RedisService, error) {
	serv := &RedisService{
		repo: repo,
		Log:  sl,
	}
	sl.Debug("s3 service successfully initiated")
	return serv, nil
}

func (rs *RedisService) Set(key string, value interface{}, expiration time.Duration) error {
	return rs.repo.Set(key, value, expiration)
}

func (rs *RedisService) Get(key string) (string, error) {
	return rs.repo.Get(key)
}
