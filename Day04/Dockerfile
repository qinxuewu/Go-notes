#基础镜像
FROM golang:latest

#工作目录
WORKDIR $GOPATH/src/github.com/qinxuewu/we-blo
#将当前上下文目录的内容复制到
COPY . $GOPATH/src/github.com/qinxuewu/we-blo
#编译
RUN go build .
# 指定端口号
EXPOSE 8080
ENTRYPOINT ["./we-blog"]