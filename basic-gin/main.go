package main

import (
	router "basic-gin/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)
 // gin  入门    https://www.cnblogs.com/-beyond/p/9391892.html
func main()  {
	//路由部分
	router:=router.InitRouter()
	router.LoadHTMLGlob("views/*")
	//使用中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//如何写日志文件

	//f,_:=os.Create("gin.log")
	//gin.DefaultWriter=io.MultiWriter(f)

	// 定义路由日志的格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

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

