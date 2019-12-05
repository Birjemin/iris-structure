## 简单的go项目结构
其实就是别人的拿过来改一下适合自己的风格

### 改动点
* 加了一个`--port`、`--path`选项，方便加载配置和指定端口
* 加了一个queue的实现
* 加了一个平滑关闭(graceful shutdown)
* 加了mod方式解决依赖

```
#适用于使用supervisor保活
go run main.go --port=8084 --path=/Data/www/Go/src/iris-structure/config.json

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

nohup /data/www/main --port=8084 --path=/Data/www/Go/src/iris-structure/config.json  >> /data/logs/main_8085.log 2>&1 &
```
### 源于
[pppercyWang](https://github.com/pppercyWang/iris-gorm-demo)