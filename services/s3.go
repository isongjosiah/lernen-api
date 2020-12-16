package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/isongjosiah/lernen-api/config"
	log "github.com/sirupsen/logrus"
	"mime/multipart"
)

const (
	//name of the s3 bucket
	S3Bucket = "lernen"
)

type IS3Service interface {
	Upload(filename string, body multipart.File) (*s3manager.UploadOutput, error)
	Download(link string) error
	Delete(key string) (*s3.DeleteObjectOutput, error)
}

type S3Service struct {
	config *config.Config
	//TODO: I am using storing the session variable now because I haven't decided which method to use in uploading the file to s3
	sess *session.Session
	s3   *s3.S3
}

// NewS3Service creates a new instance of S3 service
func NewS3Service(cfg *config.Config) (IS3Service, error) {
	conf := &aws.Config{Region: aws.String(cfg.AWSRegion)}
	sess, err := session.NewSession(conf)
	if err != nil {
		return nil, err
	}
	s3manager.NewUploader(sess)
	svc := S3Service{
		config: cfg,
		sess:   sess,
		s3:     s3.New(sess),
	}

	//possible debugging point.
	return &svc, nil
}

func (s *S3Service) Upload(filename string, body multipart.File) (*s3manager.UploadOutput, error) {
	uploader := s3manager.NewUploader(s.sess)
	//The upload output that has been ignored here might actually be useful for you in the future.
	ret, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3Bucket),
		Key:    aws.String(filename),
		Body:   body,
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *S3Service) Download(link string) error {
	return nil
}

func (s *S3Service) Delete(key string) (*s3.DeleteObjectOutput, error) {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(S3Bucket),
		Key:    aws.String(key),
	}
	res, err := s.s3.DeleteObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Error(aerr.Error())
				return nil, aerr
			}
		} else {
			log.Error(err.Error())
			return res, err
		}
	}
	return nil, nil
}
