package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/**
Go语言里面提供了一个完善的net/http包，通过http包可以很方便的搭建起来一个可以运行的Web服务。
同时使用这个包能很简单地对Web的路由，静态文件，模版，cookie等数据进行设置和操作。
*/
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

/**
http包执行流程:
	1.创建Listen Socket, 监听指定的端口, 等待客户端请求到来。
	2.Listen Socket接受客户端的请求, 得到Client Socket, 接下来通过Client Socket与客户端通信。
	3.处理客户端的请求, 首先从Client Socket读取HTTP请求的协议头, 如果是POST方法, 还可能要读取客户端提交的数据, 然后交给相应的handler处理请求,
	handler处理完毕准备好客户端需要的数据, 通过Client Socket写给客户端。
*/
func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
