package test

import (
	"bytes"
	"cloud_disk3/core/define"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"sync"

	"testing"
)

func TestFileUploadByPath(t *testing.T) {

	u, _ := url.Parse(define.TencentCloudURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			//SecretId:AKID2QrIVaVQEwTnt592z3wgRIOTTMbZQ6aF
			//SecretKey:wooiR3meTGwiV9g5Qix6Tn3EbvLi0Kcl
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: "AKID2QrIVaVQEwTnt592z3wgRIOTTMbZQ6aF", // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: "wooiR3meTGwiV9g5Qix6Tn3EbvLi0Kcl", // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})

	//SecretId:AKID2QrIVaVQEwTnt592z3wgRIOTTMbZQ6aF
	//SecretKey:wooiR3meTGwiV9g5Qix6Tn3EbvLi0Kcl
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	key := "cloud_disk/2.png"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./1.png", nil,
	)
	if err != nil {
		panic(err)
	}

}

func TestFileUploadByReader(t *testing.T) {
	u, _ := url.Parse(define.TencentCloudURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			//SecretId:AKID2QrIVaVQEwTnt592z3wgRIOTTMbZQ6aF
			//SecretKey:wooiR3meTGwiV9g5Qix6Tn3EbvLi0Kcl
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: "AKID2QrIVaVQEwTnt592z3wgRIOTTMbZQ6aF", // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: "wooiR3meTGwiV9g5Qix6Tn3EbvLi0Kcl", // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})

	//SecretId:AKID2QrIVaVQEwTnt592z3wgRIOTTMbZQ6aF
	//SecretKey:wooiR3meTGwiV9g5Qix6Tn3EbvLi0Kcl
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	key := "cloud_disk/1.png"

	file, _ := os.ReadFile("./1.png")
	_, err := client.Object.Put(
		context.Background(), key, bytes.NewReader(file), nil,
	)
	if err != nil {
		panic(err)
	}

}

func upload(wg *sync.WaitGroup, c *cos.Client, files <-chan string) {
	defer wg.Done()
	for file := range files {
		name := "cloud_disk/" + file
		fd, err := os.Open(file)
		if err != nil {
			//ERROR
			continue
		}
		_, err = c.Object.Put(context.Background(), name, fd, nil)
		if err != nil {
			//ERROR
		}
	}
}

func TestBatchUpload(t *testing.T) {
	u, _ := url.Parse(define.TencentCloudURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.TencentSecretID),  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			SecretKey: os.Getenv(define.TencentSecretKey), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})
	// 多线程批量上传文件
	filesCh := make(chan string, 2)
	filePaths := []string{"./3.png", "./4.png", "./5.png"}
	var wg sync.WaitGroup
	threadpool := 2
	for i := 0; i < threadpool; i++ {
		wg.Add(1)
		go upload(&wg, c, filesCh)
	}
	for _, filePath := range filePaths {
		filesCh <- filePath
	}
	close(filesCh)
	wg.Wait()
}
