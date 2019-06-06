package routers

import (
	"github.com/gin-gonic/gin"
	. "basic-gin/controllers"
	"io"
	"os"
)

//
func InitRouter() *gin.Engine{

	//gin.DisableConsoleColor()    //禁用日志的颜色


	//gin.ForceConsoleColor()    //  Force log's color

	//创建日志文件
	f,_:=os.Create("gin.log")
	gin.DefaultWriter=io.MultiWriter(f)


	router := gin.Default()
	router.LoadHTMLGlob("views/*")


	//Hello World
	router.GET("/", IndexApi)

	// http://localhost:8088/bindurl/qxw/987fbc97-4bed-5078-9f07-9141ba07c9f3
	// http://localhost:8088/bindurl/qxw/not-uuid  进入400
	router.GET("/bindurl/:name/:id",BidUri)

	// http://localhost:8088/getb?a=hello&b=world
	router.GET("/getb", GetDataB)

	//  http://localhost:8088/getc?a=hello&c=world
	router.GET("/getc", GetDataC)

	//   http://localhost:8088/getd?x=hello&d=world
	router.GET("/getd", GetDataD)

	// http://localhost:8088/tmpl
	router.GET("/tmpl", ShowTmplIndex)

	// http://localhost:8088/html
	router.GET("/html", ShowHtmlIndex)

	router.POST("/bindForm", FormHandler)

	// http://localhost:8088/someJSON
	router.GET("/someJSON",AscJson)
	return router
}