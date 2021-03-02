# GoAdmin

## 特征
- 基于 gin + layuiadmin
- 开箱即用的rbac认证系统
- opentelemetry,jaeger实现golang链路追踪

![](https://github.com/bensema/goadmin/blob/main/run.png)


## 使用

先建立数据库`goadmin`

```shell
$ git clone https://github.com/bensema/goadmin.git
$ cd goadmin
```

导入目录下的`goadmin.sql`到数据库`goadmin`中

```shell
$ cd cmd
$ go run main.go -conf=example.toml
```
访问：[http://127.0.0.1:8890](http://127.0.0.1:8890)

启动jaeger
```
    docker run \
    -p 5775:5775/udp \
    -p 16686:16686 \
    -p 6831:6831/udp \
    -p 6832:6832/udp \
    -p 5778:5778 \
    -p 14268:14268 \
    jaegertracing/all-in-one:latest
```

