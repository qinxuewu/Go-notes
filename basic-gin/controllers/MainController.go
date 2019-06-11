package controllers

import (
	. "basic-gin/controllers/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func IndexApi(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"第一个gin程序api接口"})
}

// 给请求连接绑定URL参数  请求格式为 端口号/自定义名称/uuid(必须为uuid)  返回200
func BidUri(c *gin.Context)  {
	var person Person
	if err :=c.ShouldBindUri(&person);err != nil{
		fmt.Println("进入400.。。。。。。。。。。。")
		c.JSON(http.StatusBadRequest,gin.H{"msg":err})
		return
	}
	c.JSON(http.StatusOK,gin.H{"name":person.Name,"uuid":person.ID})
}

//渲染模板 index.tmpl
func ShowTmplIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "gin程序加载index.tmpl 模板文件",
	})
}
func ShowHtmlIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "绑定html复选框",
	})
}



// 绑定html复选框
func FormHandler(c *gin.Context) {
	var fakeForm MyForm
	c.ShouldBind(&fakeForm)
	c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
}

// struct绑定请求连接后的参数
func BindParmStr(c *gin.Context)  {
	var s Student
	if c.ShouldBind(&s) == nil{
		log.Println("name:",s.Name)
		log.Println("address:",s.Address)
		log.Println("birthday:",s.Birthday)
	}
}

//  使用AsciiJSON生成具有转义的非ASCII字符的json数据格式
func AscJson(c *gin.Context)  {

	data:=map[string]interface{}{
		"lang":"Go语言",
		"tag":"<br>",
	}
	c.AsciiJSON(http.StatusOK,data)
}


//Goroutines
func LongAsync(c *gin.Context)  {
	cCp:=c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("使Goroutine 异步处理内部方法："+cCp.Request.URL.Path)
	}()

	c.String(http.StatusOK, "使Goroutine 异步处理内部方法："+cCp.Request.URL.Path)
}

func LoandSync(c *gin.Context)  {
	time.Sleep(5 * time.Second)
	log.Println("同步访问："+c.Request.URL.Path)
	c.String(http.StatusOK, "同步访问："+c.Request.URL.Path)
}

func LoginEndpointV1(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{"v1": "路由分组访问v1"})
}
func LoginEndpointV2(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{"v2": "路由分组访问v2"})
}


// jsonp

func ShowHtmlJSONP(c *gin.Context) {
	c.HTML(http.StatusOK, "jsonp.html", gin.H{
		"title": "使用JSONP从不同域中的服务器请求数据。如果查询参数回调存在，则向响应主体添加回调",
	})
}

func JsonP(c *gin.Context)  {
	data:=map[string]interface{}{
		"name":"使用jsonp返回数据实现跨域",
	}
	c.JSONP(http.StatusOK,data)
}

// 映射请求地址后的map类型参数
func GetMap(c *gin.Context)  {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	fmt.Printf("ids: %v; names: %v", ids, names)
	c.JSON(http.StatusOK,gin.H{"code":200,"ids":ids,"names":names})
}

// 字符进行编码
func PureJSON(c *gin.Context)  {
	c.PureJSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}

//  设置并获取cookie
func SetGetCookie(c *gin.Context)  {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		cookie = "没获取到gin_cookie 设置cookie"
		c.SetCookie("gin_cookie", "qinxuewu ", 3600, "/", "localhost", false, true)
	}

	fmt.Printf("Cookie value: %s \n", cookie)
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":cookie})
}