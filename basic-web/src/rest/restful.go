/**
 * RESTful的实现
 * @author qinxuewu
 * @create 19/5/16下午10:20
 * @since 1.0.0
 */
package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"

	"log"
)
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "访问index方法")
}

func Hello(w http.ResponseWriter,t *http.Request,ps httprouter.Params)  {
	//获取参数
	fmt.Fprintf(w,"hello, %s!\n",ps.ByName("name"))

}

func getuser(w http.ResponseWriter,r *http.Request,ps httprouter.Params)  {

	uid:=ps.ByName("uid");
	fmt.Fprintf(w,"你传入的UID是：%s",uid)
}

func modifyuser(w http.ResponseWriter,r *http.Request,ps httprouter.Params){
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are modify user %s", uid)
}

func deleteuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are delete user %s", uid)
}

func adduser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// uid := r.FormValue("uid")
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are add user %s", uid)
}

func main()  {
	//c创建路由
	router:=httprouter.New()
	router.GET("/",Index)
	router.GET("/hello/:name",Hello)  // http://localhost:8080/hello/qxw
	router.GET("/user/:uid", getuser)  //http://localhost:8080/user/1

	router.POST("/adduser/:uid", adduser)
	router.DELETE("/deluser/:uid", deleteuser)
	router.PUT("/moduser/:uid", modifyuser)
	log.Fatal(http.ListenAndServe(":8080",router))
}