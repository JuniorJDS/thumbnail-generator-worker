package infra

import (
	"log"
	"thumbnail-generator-worker/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess *session.Session

var settings = config.GetSettings()

func Session() *session.Session {
	endpoint := aws.String("http://localhost:4566/")
	s3ForcePathStyle := aws.Bool(true)

	sess, err := session.NewSession(&aws.Config{
		Endpoint: endpoint,
		Region:   aws.String(settings["REGION"]),
		Credentials: credentials.NewStaticCredentials(
			settings["AWS_ACCESS_KEY_ID"],
			settings["AWS_SECRET_ACCESS_KEY"],
			"",
		),
		S3ForcePathStyle: s3ForcePathStyle,
	})

	if err != nil {
		log.Printf("Error to connect to aws: %s\n", err.Error())
	}

	return sess
}

func GetSessionAWS() *session.Session {
	if sess == nil {
		sess = Session()
	}
	return sess
}
