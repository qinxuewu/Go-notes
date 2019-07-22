package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
	"we-blog/conf"
	router "we-blog/router"
	util "we-blog/util"
)

func main() {
	router := router.InitRouter()
	router.LoadHTMLGlob("views/**/*")

	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 定义路由日志的格式
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())

	config, _ := conf.GetConfig()

	//静态资源
	router.Static("/static", "./static")

	port := config.ServerConfig.Port
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second, //读取超时时间
		WriteTimeout:   10 * time.Second, //写超时时间
		MaxHeaderBytes: 1 << 20,          // 最大字节数 十进制的1048576字节
	}
	s.ListenAndServe()
	//// 优雅的重启
	//go func() {
	//	if err := s.ListenAndServe(); err != nil {
	//		log.Printf("Listen: %s\n", err)
	//	}
	//}()
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<- quit
	//
	//log.Println("Shutdown Server ...")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	//defer cancel()
	//if err := s.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}
	//
	//log.Println("Server exiting")

}
