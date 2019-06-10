package routers

import (
	. "basic-gin/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//
func InitRouter() *gin.Engine{

	//gin.DisableConsoleColor()    //禁用日志的颜色
	//gin.ForceConsoleColor()    //  Force log's color

	//
	////如何写日志文件
	//f,_:=os.Create("gin.log")
	//gin.DefaultWriter=io.MultiWriter(f)



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


	// 定义路由日志的格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	// http://localhost:8088/longAsync
	router.GET("/longAsync", LongAsync)
	// http://localhost:8088/loandSync
	router.GET("/loandSync", LoandSync)

	// 分组路由

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", LoginEndpointV1)
		v1.GET("/longAsync", LongAsync)

	}
	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/login", LoginEndpointV2)
		v2.GET("/loandSync", LoandSync)

	}

	router.GET("/showJsonp",ShowHtmlJSONP)
	router.GET("/JSONP",JsonP)

	//  请求地址： http://localhost:8088/post?ids[a]=1234&ids[b]=hello
	//  参数 ：names[first]=thinkerou&names[second]=tianou
	router.POST("/post",GetMap)


	// 获取请求路径中的参数
	// http://localhost:8088/user/qxw, http://localhost:8088/user/shasha
	router.GET("/user/:name",GetParam)

	// http://localhost:8088/user/john/ddd
	router.GET("/user/:name/*action",GetParam2)
	return router
}