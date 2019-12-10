## 简单的iris项目结构
其实就是别人的拿过来改一下适合自己的风格

### 改动点
* 加了一个`--port`选项，方便加载配置和指定端口
* 加了一个queue的实现
* 加了一个平滑关闭(graceful shutdown)
* 加了mod方式解决依赖

```
// 运行
go run main.go --port=8081
// 编译
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

### 部署方式

参考：[https://beego.me/docs/deploy/beego.md](https://beego.me/docs/deploy/beego.md)

> 请注意config.json配置文件的位置！！！（比如supervisor中的directory配置）

* 独立部署

``
nohup /data/www/main --port=8081  >> /data/logs/main_8081.log 2>&1 &
``

* Supervisor部署

```
directory=/data/www/Go/src/iris-structure
command=/data/www/main --port=8081
numprocs=1
autostart=true
autorestart=true
startretries=3
user=www
redirect_stderr=true
stdout_logfile=/data/logs/supervisor_xxx_8081.log
```

* Nginx负载均衡
...

### 源于
[pppercyWang](https://github.com/pppercyWang/iris-gorm-demo)