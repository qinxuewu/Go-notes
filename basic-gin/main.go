package main

import "github.com/gin-gonic/gin"

 // gin  入门
func main()  {
	// http://localhost:8080/ping  默认端口8080
	r:=gin.Default()
	r.GET("/ping", func(c *gin.Context) {
			c.JSON(200,gin.H{"code":200,"msg":"gin实现http接口"})
	})
	r.Run()
}
