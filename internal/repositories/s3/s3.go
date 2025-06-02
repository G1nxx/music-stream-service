package s3

import (
	"bytes"
	_ "context"
	_ "errors"
	"fmt"
	_ "io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	cfg "music-stream-service/internal/config"
	"music-stream-service/service/repository"
)

type S3Storage struct {
	repository.S3Storage
	s3Client   *s3.S3
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func NewS3Storage(cfg cfg.S3Config) *S3Storage {
	awsConfig := &aws.Config{
		Region: aws.String(cfg.Region),
		Credentials: credentials.NewStaticCredentials(
			cfg.AccessKeyID,
			cfg.SecretAccessKey,
			"",
		),
	}

	if cfg.Endpoint != "" {
		awsConfig.Endpoint = aws.String(cfg.Endpoint)
		awsConfig.DisableSSL = aws.Bool(cfg.DisableSSL)
		awsConfig.S3ForcePathStyle = aws.Bool(cfg.ForcePathStyle)
	}

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		log.Fatalf("Failed to create S3 session: %v", err)
	}

	svc := s3.New(sess)

	// ok, err := bucketExists(context.Background(), svc, "tracks")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// if !ok {
	// 	_, err = svc.CreateBucket(&s3.CreateBucketInput{
	// 		Bucket: aws.String("tracks"),
	// 	})
	// 	if err != nil {
	// 		log.Fatalf("Failed to create \"tracks\" bucket: %v", err)
	// 	}
	// }

	return &S3Storage{
		s3Client:   svc,
		uploader:   s3manager.NewUploaderWithClient(svc),
		downloader: s3manager.NewDownloaderWithClient(svc),
	}
}

// func bucketExists(ctx context.Context, client *s3., bucketName string) (bool, error) {
// 	_, err := client.HeadBucket(ctx, &s3.HeadBucketInput{
// 		Bucket: aws.String(bucketName),
// 	})
// 	if err != nil {
// 		var notFound *s3.NotFound
// 		if errors.As(err, &notFound) {
// 			return false, nil
// 		}
// 		return false, fmt.Errorf("failed to check bucket existence: %w", err)
// 	}
// 	return true, nil
// }

func (s *S3Storage) Upload(bucket, fileName string, fileContent []byte) error {
	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(fileContent),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %w", err)
	}
	return nil
}

func (s *S3Storage) Download(bucket, fileName string) ([]byte, error) {
	buf := aws.NewWriteAtBuffer([]byte{})

	_, err := s.downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download file from S3: %w", err)
	}

	return buf.Bytes(), nil
}

func (s *S3Storage) Delete(bucket, fileName string) error {
	_, err := s.s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %w", err)
	}

	// err = s.s3Client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
	// 	Bucket: aws.String(bucket),
	// 	Key:    aws.String(fileName),
	// })

	return err
}
