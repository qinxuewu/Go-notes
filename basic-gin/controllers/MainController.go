package controllers
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"log"
	. "basic-gin/controllers/dto"
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