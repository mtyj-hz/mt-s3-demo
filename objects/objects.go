package objects

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"os"
)

type Object struct {
	svc *s3.S3
}

// GetObjects 获取对象列表
func (o *Object) GetObjects(bucket *string) (*s3.ListObjectsV2Output, error) {
	resp, err := o.svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: bucket})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (o *Object) UploadObject(bucket, objectName *string, read io.ReadSeeker) error {
	_, err := o.svc.PutObject(&s3.PutObjectInput{
		Body:   read,
		Bucket: bucket,
		Key:    objectName,
	})
	if err != nil {
		fmt.Println("Got error putting object in bucket:")
		return err
	}
	return nil
}
func (o *Object) DownloadObject(bucket, objectName *string, w *os.File) (err error) {
	//downloader := s3manager.NewDownloader(sess)
	downloader := s3manager.NewDownloaderWithClient(o.svc)
	_, err = downloader.Download(w,
		&s3.GetObjectInput{
			Bucket: bucket,
			Key:    objectName,
		})
	return err
}
func (o *Object) DeleteObject(bucket, objectName *string) error {
	_, err := o.svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: bucket,
		Key:    objectName,
	})
	if err != nil {
		return err
	}
	return o.svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: bucket,
		Key:    objectName,
	})
}
func (o *Object) UploadObjectS() {

}
