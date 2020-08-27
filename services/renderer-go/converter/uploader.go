package converter

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

type Uploader interface {
	Upload(image image.Image) (string, error)
}

type AWSUploader struct {
	session *session.Session
}

func NewAWSUploder() (*AWSUploader, error) {
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
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, data); err != nil {
		return "", err
	}
	fileName := uuid.New().String() + ".png"
	bucketName := os.Getenv("AWS_BUCKET_NAME")
	if bucketName == "" {
		return "", fmt.Errorf("empty bucket name")
	}
	uploader := s3manager.NewUploader(au.session)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucketName,
		Key:    &fileName,
		Body:   buf,
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.s3-ap-northeast-1.amazonaws.com/%s", bucketName, fileName), nil
}
