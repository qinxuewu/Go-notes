package routers

import (
	. "basic-gin/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

//
func InitRouter() *gin.Engine{

	router := gin.Default()


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

	//http://localhost:8088/pureJson
	router.GET("/pureJson",PureJSON)


	// http://localhost:8088/getKeyValue?id=1234&page=1
	router.POST("/getKeyValue",GetKeyValue)

	// http://localhost:8088/welCome?firstname=qxw&lastname=aaaaa
	router.GET("/welCome",WelCome)


	//重定向  http://localhost:8088/redirect?status=0
	router.GET("/redirect", func(c *gin.Context) {
		status:=c.DefaultQuery("status","0")
		if status=="0" {
			c.Redirect(http.StatusMovedPermanently, "https://blog.qinxuewu.club/")
		}else {
			c.Request.URL.Path = "/test2"
			router.HandleContext(c)
		}
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})


	// 设置并获取cookie  http://localhost:8088/cookie
	router.GET("/cookie",SetGetCookie)

	// 文件上传  http://localhost:8088/fileHtml
	router.GET("/fileHtml",FileHtml)
	router.POST("/upload",UploadFile)


	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))


	//访问链接弹出账号密码登录框   http://localhost:8088/admin/secrets
	authorized.GET("/secrets",BasicAuth)
	return router
}