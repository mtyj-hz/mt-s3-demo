package share

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

type Link struct {
	svc *s3.S3
}

func (l *Link) Perlink(bucket *string) error {
	acl := "public-read"
	l.svc.PutBucketAcl(&s3.PutBucketAclInput{
		Bucket: bucket,
		ACL:    &acl,
	})
	return nil
}
