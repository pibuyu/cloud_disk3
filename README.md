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
```