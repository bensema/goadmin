# GoAdmin

## 特征
- 基于 golang + react 构建管理后台

![](https://github.com/bensema/goadmin/blob/main/run.png)


## 使用

先建立数据库`goadmin`

```shell
$ git clone https://github.com/bensema/goadmin.git
$ cd goadmin
```

导入目录下的`goadmin.sql`到数据库`goadmin`中

### Backend
```shell
$ cd cmd
$ go run main.go
```

### Frontend 
```shell
$ cd web
$ npm start
```

访问：[http://127.0.0.1:3000](http://127.0.0.1:3000) 账户:root 密码:mozi123

