package gateways

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/ghmeier/bloodlines/config"
)

type S3 interface {
	Upload(bucket, name string, body io.Reader) (string, error)
}

type s3Client struct {
	config config.S3
	client *s3.S3
}

func NewS3(config config.S3) S3 {

	// Create S3 service client
	s := &s3Client{config: config}
	if config.Region != "" {
		os.Setenv("AWS_REGION", config.Region)
	}
	if config.AccessKey != "" {
		os.Setenv("AWS_ACCESS_KEY_ID", config.AccessKey)
	}
	if config.Secret != "" {
		os.Setenv("AWS_SECRET_ACCESS_KEY", config.Secret)
	}

	s.connect()
	return s
}

func (s *s3Client) connect() {
	if s.client != nil {
		return
	}

	sess := session.Must(session.NewSession(aws.NewConfig()))
	s.client = s3.New(sess)
}

func (s *s3Client) Upload(location, name string, body io.Reader) (string, error) {
	s.connect()

	err := s.client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(s.bucket(location)),
	})
	if err != nil {
		return "", err
	}

	uploader := s3manager.NewUploaderWithClient(s.client)
	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucket(location)),
		Key:    aws.String(name),
		Body:   body,
	})
	if err != nil {
		return "", err
	}

	fmt.Println(res.Location)
	return res.Location, nil
}

func (s *s3Client) bucket(name string) string {
	return fmt.Sprintf("%s-%s", s.config.Bucket, name)
}
