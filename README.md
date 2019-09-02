# skill_blog

> 这项目只是介绍怎么使用iris,及相关的框架集成使用

iris(mvc) + mysql + go-jwt + redis + yaag(文档) blog demo 

## 目录内容

- [背景](#背景)
- [拉取项目](#拉取项目)
- [使用](#使用)
- [API](#api)
- [Contributing](#contributing)
- [License](#license)

## 背景
为了供他人更加快速的上手iris框架 提供相关demo，供参数
   - 采用 iris 框架
   - 采用 gorm 数据库模块 和 jwt 的单点登陆认证方式
   - 采用 redis 缓存集成
   - 测试默认使用了 mysql 数据库
   
## 拉取项目

该项目是使用mod,需要go版本1.12.7 以上
- 还需要配置下代理https://goproxy.io 如果使用的goland File->setting->Go->go modules set proxy

```
go install github.com/bobacsmall/skill_blog
```

## 使用

```
go build main.go
```

Note: 需要安装mysql redis 具体安装参照[redis](https://redis.io/)  [mysql](https://www.mysql.com/downloads/) 

### 目录结构
```
com.github.bobacsmall
    ├─config  配置
    ├─datamodels 实体
    ├─datasource 数据源
    ├─pkg        工具类
    │  ├─redis
    │  ├─response
    │  └─setting
    ├─repositories 数据操作层
    ├─route        路由
    ├─services     数据服务
    └─web
        ├─controllers controller
        ├─middleware  中间件
        └─viewmodels

```
### api文档
自动生成文档 (访问过接口就会自动成功) 因为原生的 jquery.min.js 里面的 cdn 是使用国外的，访问很慢。 有条件的可以开个 vpn ,如果没有可以根据下面的方法修改一下，访问就很快了
```
1.打开index.html文件
https://ajax.googleapis.com/ajax/libs/jquery/2.1.3/jquery.min.js
国内的 cdn
https://cdn.bootcss.com/jquery/2.1.3/jquery.min.js

2.直接修改yaag->yaag/base.go 模板文件后续生成的不用老是替换

```

## API
```
GET   path：/article/:param1
GET   path：/article/delete/:param1
POST  path：/article/add
POST  path：/article/edit
POST  path：/user/login
POST  path：/user/register
GET   path：/tag/delete/:param1
POST  path：/tag/add
POST  path：/tag/edit
```

### 说明
由于网上iris相关框架集成项目比较少，而且有些demo不完善，利用空闲时间写了demo供他人学习使用，没有提供前端的demo ,
后续有时间在加上，有本人时间有限，如有人的不足的，请指出，大家共同交流学习

如果感觉对你有帮助麻烦start下Thanks♪(･ω･)ﾉ 以是鼓励支持
有问题欢迎留言

## Contributing

参考 [IrisApiProject](https://github.com/snowlyg/IrisApiProject) 感谢.


## License
供参考学习使用