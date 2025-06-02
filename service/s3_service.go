package service

import (
	"log/slog"
	"music-stream-service/service/interfaces"
	"music-stream-service/service/repository"
)

type S3Service struct {
	interfaces.S3Service
	repo repository.S3Storage
	Log  *slog.Logger
}

func NewS3Service(repo repository.S3Storage, sl *slog.Logger) (*S3Service, error) {
	serv := &S3Service{
		repo: repo,
		Log:  sl,
	}
	sl.Debug("s3 service successfully initiated")
	return serv, nil
}

func (s *S3Service) Upload(bucket, fileName string, fileContent []byte) error {
	return s.repo.Upload(bucket, fileName, fileContent)
}

func (s *S3Service) Download(bucket, fileName string) ([]byte, error) {
	return s.repo.Download(bucket, fileName)
}

func (s *S3Service) Delete(bucket, fileName string) error {
	return s.repo.Delete(bucket, fileName)
}