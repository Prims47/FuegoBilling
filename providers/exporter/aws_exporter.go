package providersExporter

import (
	"bytes"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSExporter struct{}

func (a *AWSExporter) Save(fileName string, data []byte) error {
	if os.Getenv("AWS_ACCESS_KEY_ID") == "" ||
		os.Getenv("AWS_SECRET_ACCESS_KEY") == "" ||
		os.Getenv("AWS_REGION") == "" ||
		os.Getenv("AWS_BUCKET_NAME") == "" {
		return errors.New("We must provide valid env var like 'AWS_ACCESS_KEY_ID', 'AWS_SECRET_ACCESS_KEY', 'AWS_REGION and AWS_BUCKET_NAME'")
	}

	sess := session.Must(session.NewSession())

	svc := s3.New(sess)

	body := bytes.NewReader(data)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(fileName),
		Body:   body,
	})

	return err
}

func (a *AWSExporter) CanSave(exporterProviderName string) bool {
	if "aws" == exporterProviderName || "AWS" == exporterProviderName {
		return true
	}

	return false
}
