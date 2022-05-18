package buckets

import (
	session2 "aws-sdk/session"
	"fmt"
	"testing"
)

func TestService_BucketExists(t *testing.T) {
	s := &Service{
		svc: session2.GetSession(),
	}
	bucketName := "yzq-test12y"
	exist := s.BucketExists(&bucketName)
	fmt.Println(exist)
}

func Test_All(t *testing.T) {
	s := &Service{
		svc: session2.GetSession(),
	}
	// 判断bucket是否存在
	bucketName := "aws-Sdk"
	exist := s.BucketExists(&bucketName)
	//fmt.Println(exist)
	if !exist {
		// 不存在新建桶
		err := s.CreateBucket(&bucketName)
		if err != nil {
			t.Log(bucketName, "桶创建失败")
			return
		}
	}
	list, err := s.GetBucketList()
	if err != nil {
		t.Log("获取bucket List失败")
		return
	}
	fmt.Println(list)
	err = s.DeleteBucket(&bucketName)
	if err != nil {
		return
	}
	t.Log("ok!")
}

func TestService_GetBucketACL(t *testing.T) {
	s := &Service{
		svc: session2.GetSession(),
	}
	bucketName := "yzq-test"
	acl, err := s.GetBucketACL(&bucketName)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(acl)
}
