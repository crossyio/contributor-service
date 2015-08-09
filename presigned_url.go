package main

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type PresignedURL struct {
	URL     string `json:"url"`
	Timeout int    `json:"timeout"`
}

func GetS3Presigned(bucket, key string, timeout int) PresignedURL {
	svc := s3.New(nil)
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	url, err := req.Presign(time.Duration(timeout) * time.Second)
	if err != nil {
		panic(err)
	}

	p := PresignedURL{URL: url, Timeout: timeout}
	return p
}
