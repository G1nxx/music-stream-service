package service

import (
	"database/sql"
	"log/slog"
	"music-stream-service/internal/config"
	"music-stream-service/internal/repositories/postgresql"
	_ "music-stream-service/internal/repositories/redis"
	"music-stream-service/internal/repositories/s3"
	serviceInterfaces "music-stream-service/service/interfaces"
	"music-stream-service/service/repository"
)

type Service struct {
	AuthService      serviceInterfaces.Authorization
	UserService      serviceInterfaces.UserActivity
	S3Service        serviceInterfaces.S3Service
	RedisService     serviceInterfaces.RedisService
	MusicInfoService serviceInterfaces.MusicInfoActivity
	serviceInterfaces.TokenAuth
	Log *slog.Logger
}

func NewRepository(db *sql.DB, cfg *config.Config, sl *slog.Logger) *repository.Repository {

	repo := &repository.Repository{
		AuthRepo:  postgresql.NewAuthPostgres(db),
		UserRepo:  postgresql.NewUserPostgres(db),
		S3Storage: s3.NewS3Storage(cfg.S3Config),
		//RedisStorage: redis.NewRedisStorage(cfg.RedisConfig),
		MusicInfoRepo: postgresql.NewMusicInfoPostgres(db),
		Log:           sl,
	}
	sl.Debug("repository successfully initiated")
	return repo
}

func NewService(repo repository.Repository, sl *slog.Logger) (*Service, error) {
	AuthService, err := NewAuthService(repo.AuthRepo, sl)
	if err != nil {
		return nil, err
	}

	UserService, err := NewUserService(repo.UserRepo, sl)
	if err != nil {
		return nil, err
	}

	S3Service, err := NewS3Service(repo.S3Storage, sl)
	if err != nil {
		return nil, err
	}

	RedisService, err := NewRedisService(repo.RedisStorage, sl)
	if err != nil {
		return nil, err
	}

	MusicInfoService, err := NewMusicInfoService(repo.MusicInfoRepo, sl)
	if err != nil {
		return nil, err
	}

	serv := &Service{
		AuthService:      AuthService,
		UserService:      UserService,
		S3Service:        S3Service,
		RedisService:     RedisService,
		MusicInfoService: MusicInfoService,
		Log:              sl,
	}
	sl.Info("service successfully initiated")
	return serv, nil
}
