```text
//生成core文件夹
goctl api new core 

cd core


//直接运行默认服务
go run core.go -f etc/core-api.yaml

//启动服务
go run core.go -f etc/core-api.yaml 

//使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero


//腾讯云首页
https://console.cloud.tencent.com/cos/bucket
//腾讯云对象存储文档
https://cloud.tencent.com/document/product/436/31215
```