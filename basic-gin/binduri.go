package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Person  struct {
	ID string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}
// gin  入门
func main()  {

	// http://localhost:8088/qxw/987fbc97-4bed-5078-9f07-9141ba07c9f3
	// http://localhost:8088/qxw/not-uuid  进入400
	r:=gin.Default()

	//给请求连接绑定URL参数  请求格式为 端口号/自定义名称/uuid(必须为uuid)  返回200
	r.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err :=c.ShouldBindUri(&person);err != nil{
			fmt.Println("进入400.。。。。。。。。。。。")
			c.JSON(400,gin.H{"msg":err})
			return
		}
		c.JSON(200,gin.H{"name":person.Name,"uuid":person.ID})
	})
	r.Run(":8088")
}
