package interfaces

type S3Service interface {
	Upload(bucket, fileName string, fileContent []byte) error
	Download(bucket, fileName string) ([]byte, error)
	Delete(bucket, fileName string) error
}
