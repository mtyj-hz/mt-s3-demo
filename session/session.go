package session

import (
	"github.com/aws/aws-sdk-go/aws"
	awsAls "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetSession() *s3.S3 {
	//accessKey := "6GrE30h7e25I3VnwKWp3"
	//secretKey := "vExxPVZSFwP11nEPaqDe4ehBYLNLGa"
	//endPoint := "http://192.168.1.68:9000" //endpoint设置
	accessKey := ""
	secretKey := ""
	endPoint := "https://oss.mty.wang" //endpoint设置
	sess, err := session.NewSession(&aws.Config{
		Credentials:      awsAls.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(endPoint),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true), //virtual-host style方式
	})
	if err != nil {
		return nil
	}
	svc := s3.New(sess)
	return svc
}

//
//func GetMinioClient() *minio.Client{
//	endpoint := "192.168.1.68:9000"
//	accessKeyID := "YQsyBLy2E6upkYQWyjs3"
//	secretAccessKey := "RdKbtpu5JHu0yU8qlwR6xVP45O2WMV"
//	//accessKeyID := "4xE7yaFk2KQpq7fgBoWs"
//	//secretAccessKey := "xE0RNDXkr15buJYE8nAdJqu7v327u8"
//	// Initialize minio client object.
//	minioClient, err := minio.New(endpoint, &minio.Options{
//		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
//		Secure: false,
//	})
//	if err != nil {
//		return nil
//	}
//	return minioClient
//}
