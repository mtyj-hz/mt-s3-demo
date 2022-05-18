package main

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"net/http"
	"net/url"
	"time"
)

var expiry = time.Second * 24 * 60 * 60 * 7 // 7 day.
func postObject(client *minio.Client) {
	preassignedURL, err := client.Presign(context.Background(), http.MethodPost, "yzq-test", "JSP任意文件下载漏洞拿下webshell.pdf", expiry, nil)
	if err != nil {
		return
	}
	fmt.Println(preassignedURL)
}
func getObject(client *minio.Client) {
	params := url.Values{}
	//params.Set("versionId", "123")
	preassignedURL, err := client.PresignedGetObject(context.Background(), "encrypto", "Downloads.zip", expiry, params)
	if err != nil {
		return
	}
	fmt.Println(preassignedURL)
}
func FromContentMD5() (string, error) {
	v := "/z3gX1MZrDyoGZMbOvo+Sw=="
	le, err := hex.DecodeString(v)
	fmt.Println(len(v), le)
	b, err := base64.StdEncoding.Strict().DecodeString(v)
	if err != nil {
		return "", err
	}
	l := len(b)
	if l != md5.Size {
		return "", errors.New("etag: invalid content-md5")
	}
	return "", nil
}

func main() {
	//client := session.GetMinioClient()
	////postObject(client)
	////FromContentMD5()
	//getObject(client)
	//var sm sync.Map
	//sm.Store("", 123)
	//sm.Store("", 1234)
	//sm.Range(func(key, value interface{}) bool {
	//	fmt.Println(key, value)
	//	return true
	//})
	//return
	//md5c := fmt.Sprintf("%x", md5.Sum([]byte("QmVtkos5UrNu66G1xSrjne3UtQWKuuqWheNTjjdmbxc5Mk")))
	//fmt.Println(md5c)
	for i := 0; i < -1; i++ {
		fmt.Println("aaaaaaaaaaa")
	}
	//hex.EncodeToString()
	return
}
