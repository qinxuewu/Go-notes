package main

import (
	"net/http"
	"html/template"
	"log"
	"fmt"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
)


func upload(w http.ResponseWriter, r *http.Request)  {
	//获取请求的方法
	fmt.Println("method",r.Method)

	if r.Method == "GET" {


		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("/Users/qinxuewu/Desktop/GiWork/Go-notes/basic-web/src/upload/upload.gtpl")

		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)  // 此处假设当前目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}

}


func main() {
	http.HandleFunc("/upload", upload)         //设置访问的路由
	err := http.ListenAndServe(":9091", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
