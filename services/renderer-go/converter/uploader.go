package converter

import (
	"image"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Uploader interface {
	Upload(image image.Image) (string, error)
}

type AWSUploader struct {
	session *session.Session
}

func (au *AWSUploader) NewAWSUploder() (*AWSUploader, error) {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	if err != nil {
		return nil, err
	}
	return &AWSUploader{
		session: s,
	}, nil
}

func (au *AWSUploader) Upload(data image.Image) (string, error) {
	return "http://locahost:", nil
}
