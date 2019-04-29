package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)




//预防跨站脚本
//func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
//func HTMLEscapeString(s string) string //转义s之后返回结果字符串
//func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("/Users/qinxuewu/Desktop/GiWork/Go-notes/basic-web/src/escape/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//默认情况下，Handler里面是不会自动解析form的，必须显式的调用r.ParseForm()
		r.ParseForm()




		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端

		//fmt.Println("username:", r.Form["username"])
		//fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9091", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
