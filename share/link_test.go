package share

import (
	session2 "aws-sdk/session"
	"fmt"
	"testing"
)

func TestLink_Perlink(t *testing.T) {
	o := &Link{
		svc: session2.GetSession(),
	}
	bucketName := "y1211"
	err := o.Perlink(&bucketName)
	fmt.Println(err)
}
