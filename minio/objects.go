package main

import (
	"aws-sdk/session"
	"context"
	"fmt"
	mtoss "github.com/mtyj-hz/mtoss-go-sdk"
)

func getObjectCid() {
	m := session.GetMinioClient()
	ctx := context.Background()
	opts := mtoss.GetObjectOptions{}
	objectInfo, err := m.StatObject(ctx, "y1211", "app.zip", opts)
	if err != nil {
		return
	}
	fmt.Println(objectInfo.Cid)
}
