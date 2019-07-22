# we-blog
gin+grom整合crud实现简单api接口

# 目录结构
```
├── Dockerfile
├── LICENSE
├── README.md
├── blog.sql
├── conf
│   └── config.go
├── config.yaml
├── controller
│   ├── ArticleController.go
│   ├── AuthController.go
│   ├── TagController.go
│   └── UploadController.go
├── cron
│   └── cron.go
├── filter
│   ├── filter.go
│   └── jwt.go
├── gin.log
├── main.go
├── model
│   ├── article.go
│   ├── auth.go
│   ├── category.go
│   ├── db.go
│   ├── tag.go
│   └── user.go
├── router
│   └── router.go
├── static
│   ├── css
│   ├── images
│   └── js
├── test
│   └── test.go
├── upload
│   └── images
├── util
│   ├── code.go
│   ├── excel.go
│   ├── file.go
│   ├── image.go
│   ├── jwt.go
│   ├── logging.go
│   ├── md5.go
│   ├── msg.go
│   ├── pagination.go
├── views
│   └── web
│       ├── about.html
│       ├── daohang.html
│       ├── index.html
│       ├── info.html
│       ├── list.html
│       └── time.html

```