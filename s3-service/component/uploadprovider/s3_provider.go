package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
	"s3-service/common"
)

type s3Provider struct {
	buckerName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(
	bucketName string,
	region string,
	apiKey string,
	secret string,
	domain string,
) *s3Provider {
	provider := &s3Provider{
		buckerName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}
	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey,
			provider.secret,
			"",
		),
	})
	if err != nil {
		log.Fatalln(err)
	}
	provider.session = s3Session
	return provider
}

func (provider s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.File, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.buckerName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})
	if err != nil {
		return nil, err
	}
	img := &common.File{
		Url: fmt.Sprintf("%s/%s", "https://"+provider.domain, dst),
	}
	return img, nil
}
