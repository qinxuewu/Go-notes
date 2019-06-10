package main

import (
	router "basic-gin/routers"
	"net/http"
	"time"
)
 // gin  入门    https://www.cnblogs.com/-beyond/p/9391892.html
func main()  {
	//路由部分
	router:=router.InitRouter()


	//静态资源
	router.Static("/static", "./static")

	//运行的端口
	//router.Run(":8088")

	// 自定义HTTP配置
	//http.ListenAndServe(":8088",router)

	// 或
	s:=&http.Server{
		Addr:":8088",
		Handler:router,
		ReadTimeout:10 * time.Second,  //读取超时时间
		WriteTimeout:10 * time.Second, //写超时时间
		MaxHeaderBytes: 1<< 20,  // 最大字节数 十进制的1048576字节
	}
	s.ListenAndServe()
}

