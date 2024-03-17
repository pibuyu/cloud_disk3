package helper

import (
	"cloud_disk3/core/define"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"path"
)

func UploadFile(r *http.Request) (string, error) {

	u, _ := url.Parse(define.TencentCloudURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	file, header, err := r.FormFile("file")
	//key是文件在云上的名称
	key := "cloud_disk/" + GetUUID() + path.Ext(header.Filename)

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		panic(err)
	}
	//返回文件在云上的完整路径
	return define.TencentCloudURL + "/" + key, err
}
