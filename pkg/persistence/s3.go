package persistence

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/idasilva/chdb-dump/pkg/logger"
	"go.uber.org/zap"
)

type S3 struct {
	session  *session.Session
	logger   *zap.Logger
	name     string
	filename string
}

func (s3 *S3) Store(database, docs string) error {
	uploader := s3manager.NewUploader(s3.session)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3.name),
		Key:    aws.String(s3.filename),
		Body:   nil,
	})

	if err != nil {
		s3.logger.Fatal("FATAL")
	}

	return nil
}

func (s3 *S3) Get(database, docs string) error {
	s3.logger.Info("getting all docs to restore...")
	return nil
}

func NewS3() Storage {
	logger := logger.New()
	session := session.Must(session.NewSession())
	return &S3{
		session: session,
		logger:  logger,
	}

}
