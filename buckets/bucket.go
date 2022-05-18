package buckets

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Service struct {
	svc *s3.S3
}

// CreateBucket 创建桶
func (s *Service) CreateBucket(bucket *string) error {
	_, err := s.svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: bucket,
	})
	if err != nil {
		return err
	}
	err = s.svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: bucket,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) DeleteBucket(bucket *string) error {
	_, err := s.svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: bucket,
	})
	if err != nil {
		return err
	}
	err = s.svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: bucket,
	})
	if err != nil {
		return err
	}
	return nil
}

// BucketExists 存在返回  true  不存在返回false
func (s *Service) BucketExists(bucket *string) (exist bool) {
	bucketInput := &s3.HeadBucketInput{
		Bucket: bucket,
	}
	err := s.svc.WaitUntilBucketExistsWithContext(context.TODO(), bucketInput, request.WithWaiterMaxAttempts(2))
	return err == nil
	//ch := make(chan bool)
	//go func() {
	//	err := s.svc.WaitUntilBucketNotExists(bucketInput)
	//	if err != nil {
	//		ch <- true
	//	}
	//	ch <- false
	//}()
	//go func() {
	//	err := s.svc.WaitUntilBucketExists(bucketInput)
	//	if err != nil {
	//		ch <- false
	//	}
	//	ch <- trues
	//}()
	//return <-ch
}

func (s *Service) GetBucketList() (*s3.ListBucketsOutput, error) {
	result, err := s.svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) GetBucketACL(bucket *string) (*s3.GetBucketAclOutput, error) {
	return s.svc.GetBucketAcl(&s3.GetBucketAclInput{
		Bucket: bucket,
	})
}
