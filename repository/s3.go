package repository

import (
	"thumbnail-generator-worker/config"
	infra "thumbnail-generator-worker/infra/aws"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Repository struct {
	Uploader *s3manager.Uploader
	Bucket   string
}

var settings = config.GetSettings()

func NewS3Repository() *S3Repository {
	session := infra.GetSessionAWS()
	uploader := s3manager.NewUploader(session, func(u *s3manager.Uploader) {
		u.PartSize = 64 * 1024 * 1024
	})
	return &S3Repository{
		Uploader: uploader,
		Bucket:   settings["BUCKET"],
	}
}

func (s *S3Repository) Save() error {
	return nil
}

func (s *S3Repository) Get() error {
	return nil
}
