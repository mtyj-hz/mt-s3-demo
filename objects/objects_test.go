package objects

import (
	session2 "aws-sdk/session"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestObject_GetObjects(t *testing.T) {
	o := &Object{
		svc: session2.GetSession(),
	}
	bucketName := "yzq-test"
	exist, _ := o.GetObjects(&bucketName)
	fmt.Println(exist)
}
func TestObject_UploadObject(t *testing.T) {
	o := &Object{
		svc: session2.GetSession(),
	}
	// snippet-start:[s3.go.create_bucket_and_object.put_args]
	bucket := flag.String("b", "mtyw-sdk-test", "The name of the bucket")
	key := flag.String("k", "_TestFile.txt", "The name of the object (key)")
	flag.Parse()
	err := o.UploadObject(bucket, key, strings.NewReader("Hello World!"))
	if err != nil {
		fmt.Println("err")
		return
	}

	fmt.Println("Successfully created bucket " + *bucket + " and uploaded data with key " + *key)
}

func TestObject_DownloadObject(t *testing.T) {
	o := &Object{
		svc: session2.GetSession(),
	}
	bucket := flag.String("b", "aws-sdk-test", "The name of the bucket")
	key := flag.String("k", "_TestFile.txt", "The name of the object (key)")
	flag.Parse()
	file, err := os.Create("./TestFile.txt")
	err = o.DownloadObject(bucket, key, file)
	if err != nil {
		return
	}
}
