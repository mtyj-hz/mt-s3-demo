package share

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

type Link struct {
	svc *s3.S3
}

// Perlink 这只桶的ACl可用通过  http(s)://域名/桶名称/图片名称.jpg直接访问object
func (l *Link) Perlink(bucket *string) error {
	// Valid Values: private | public-read | public-read-write | authenticated-read
	acl := "public-read"
	l.svc.PutBucketAcl(&s3.PutBucketAclInput{
		Bucket: bucket,
		ACL:    &acl,
	})
	return nil
}
