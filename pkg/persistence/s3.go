package persistence

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/idasilva/chdb-dump/pkg/logger"
	"go.uber.org/zap"
)

type S3 struct {
	session  interface{}
	logger   *zap.Logger
	name     string
	filename string
}

func (s3 *S3) Store() {
	uploader := s3manager.NewUploader(s3.session)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3.name),
		Key:    aws.String(s3.filename),
		Body:   "",
	})

	if err != nil {
		s3.logger.Fatal("FATAL")
	}
}

func NewS3() (*S3, error) {
	logger := logger.New()
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	if err != nil {
		log.Fatal("FATAL")
		return err
	}

	return &S3{
		session: s3.New(sess),
		logger:  logger,
	}, nil

}
