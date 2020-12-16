package services

import (
	"github.com/isongjosiah/lernen-api/config"
	"github.com/pkg/errors"
)

type AWS struct {
	S3 IS3Service
}

//NewAWS is a factory method for the AWS wrapper
func NewAWS(cfg *config.Config) (*AWS, error) {
	s3, err := NewS3Service(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to set up s3 service")
	}
	aws := &AWS{
		S3: s3,
	}
	return aws, nil
}
