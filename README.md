# api-gateway-golang
基于go语言gin框架的接口网关服务模版


## 环境
- `go`1.23.6
- 使用vscode，安装插件`Go`。

## 开发运行
- 安装`air`：`go install github.com/air-verse/air@latest`
- 检查`air`是否安装成功：`air -v`
- `go mod tidy`
- 运行：`air`

## 打包
- `go build -o tmp/api-gateway`
- 运行：`./tmp/api-gateway`
