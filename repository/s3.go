package repository

import (
	"mime/multipart"
	"thumbnail-generator-worker/config"
	infra "thumbnail-generator-worker/infra/aws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Repository struct {
	sess   *session.Session
	Bucket string
}

var settings = config.GetSettings()

func NewS3Repository() *S3Repository {
	return &S3Repository{
		sess:   infra.GetSessionAWS(),
		Bucket: settings["BUCKET"],
	}
}

func (s *S3Repository) Upload(file *multipart.FileHeader, ID string) error {
	/*
		f, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open file: %s", err.Error())
		}
		defer f.Close()

		destination := fmt.Sprintf("raw/%s/%s", ID, file.Filename)

		_, err = s.Uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(s.Bucket),
			Key:    aws.String(destination),
			Body:   f,
		})

		if err != nil {
			return fmt.Errorf("failed to upload file to S3 in Bucket: %s with Error: %s", s.Bucket, err.Error())
		}
	*/
	return nil
}

func (s *S3Repository) Download(ID, fileName string) ([]byte, error) {
	// TODO: path my-bucket/raw/id/dosimage.svg
	// TODO: save my-bucket/thumbnail/id/dosimage.svg
	buf := aws.NewWriteAtBuffer([]byte{})

	downloader := s3manager.NewDownloader(s.sess)

	_, err := downloader.Download(
		buf, &s3.GetObjectInput{
			Bucket: aws.String(""),
			Key:    aws.String(""),
		},
	)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
