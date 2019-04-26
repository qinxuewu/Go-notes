package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	//fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//for k, v := range r.Form {
	//	fmt.Println("key:", k)
	//	fmt.Println("val:", strings.Join(v, ""))
	//}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("F:/ideWork/githubWork/Go-notes/basic-web/src/form/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//默认情况下，Handler里面是不会自动解析form的，必须显式的调用r.ParseForm()
		r.ParseForm()

		//验证表单的输入  必填字段
		if len(r.Form["username"][0]) == 0 {
			//为空的处理
			fmt.Println("username为空")
			fmt.Fprintf(w, "username为空") //这个写入到w的是输出到客户端的
			return
		}

		//判断正整数，那么我们先转化成int类型，然后进行处理
		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			//数字转化出错了，那么可能就不是数字
			return
		}

		//接下来就可以判断这个数字的大小范围了
		if getint > 100 {
			//太大了
			return
		}

		//还有一种方式就是正则匹配的方式
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
			return
		}

		//中文
		if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {
			return
		}

		//英文 我们期望通过表单元素获取一个英文值，例如我们想知道一个用户的英文名，应该是astaxie，而不是asta谢。
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
			return
		}

		// 电子邮件地址  你想知道用户输入的一个Email地址是否正确，通过如下这个方式可以验证：
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Println("no")
		} else {
			fmt.Println("yes")
		}

		//手机号码
		if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
			return
		}

		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9091", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
