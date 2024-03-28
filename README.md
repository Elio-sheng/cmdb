# gin-cmdb

### Gin Project Template
> 本项目使用 gin 框架为核心搭建Devops平台


### 运行
执行迁移文件：

安装migrate [查看安装文档](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

使用migrate [更多用法](https://github.com/golang-migrate/migrate)
```shell
# 执行迁移文件创建数据表
migrate -database 'mysql://root:root@tcp(127.0.0.1:3306)/go_layout?charset=utf8mb4&parseTime=True&loc=Local' -path data/migrations up
```
启动服务：
```shell
# 建议开启GO111MODULE
# go env -w GO111MODULE=on

# 下载依赖
go mod download

# 首次运行会自动复制一份示例配置（config/config.example.yaml）文件到config目录(config/config.yaml)
# 启动服务
go run main.go server

# 项目起来后执行下面命令访问示例路由
curl "http://127.0.0.1:9001/ping"
# {"message":"pong"}
curl "http://127.0.0.1:9001/api/v1/hello-world"
# {"code":0,"message":"OK","data":{"result":"hello gin-layout"},"cost":"6.151µs"}
curl "http://127.0.0.1:9001/api/v1/hello-world?name=world"
# {"code":0,"message":"OK","data":{"result":"hello world"},"cost":"6.87µs"}
```
更多用法使用以下命令查看:
```shell
go run main.go -h
```

### 部署
```shell
# 打包项目（如何打包其他os平台的包自行 google）
go build -o cmd/cmdb main.go

# 运行时请配置指定config文件的位置，否则可能会出现找不到配置的情况，修改完配置请重启
cmd/cmdb server -c="指定配置文件位置（/home/cmdb-config.yaml）"

# nginx 反向代理配置示例
server {
    listen 80;
    server_name api.xxx.com;
    location / {
        proxy_set_header Host $host;
        proxy_pass http://172.0.0.1:9001;
    }
}
```

### 生产环境注意事项
> 在构建生产环境时，请配置好 `.yaml` 文件中基础路径 `base_path`，所有的日志记录文件会保存在该目录下的 `{base_path}/cmdb/logs/` 里面，该基础路径默认为执行命令的目录

### 其他说明
##### 项目中使用到的包
- 核心：[gin](https://github.com/gin-gonic/gin)
- CLI命令管理：[github.com/spf13/cobra](https://github.com/spf13/cobra)
- 配置：[spf13/viper](https://github.com/spf13/viper)
- 参数验证：[github.com/go-playground/validator/v10](https://github.com/go-playground/validator)
- 日志：[go.uber.org/zap](https://github.com/uber-go/zap)、[github.com/natefinch/lumberjack](http://github.com/natefinch/lumberjack)、[github.com/lestrrat-go/file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs)
- 数据库：[gorm.io/gorm](https://github.com/go-gorm/gorm)、[go-redis/v8](https://github.com/go-redis/redis)
- 还有其他不一一列举，更多请查看`go.mod`文件
