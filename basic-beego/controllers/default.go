package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
)

type MainController struct {
	beego.Controller
}

// 表单中的值直接解析到 struct
type user struct {
	 // 如果要忽略一个字段，有两种办法，一是：字段名小写开头，二是：form 标签的值设置为 -
 	Id    int         `form:"-"`
	// 定义 struct 时，字段名后如果有 form 这个 tag，则会以把 form 表单里的 name 和 tag 的名称一样的字段赋值给这个字段，
	// 否则就会把 form 表单里与字段名一样的表单内容赋值给这个字段
	Messages  interface{} `form:"message"`
	Nums   int         `form:"nums"`
}



// 如果开启了 XSRF，那么 beego 的 Web 应用将对所有用户设置一个 _xsrf 的 cookie 值（默认过期 1 小时），
// 如果 POST PUT DELET 请求中没有这个 cookie 值，那么这个请求会被直接拒绝
// 。如果你开启了这个机制，那么在所有被提交的表单中，你都需要加上一个域来提供这个值。你可以通过在模板中使用 专门的函数 XSRFFormHTML()
func (c *MainController) Get() {
	c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.Data["Website"] = "第一个beego程序"
	c.Data["Email"] = "870439570@qq.com"
	c.TplName = "index.tpl"
}


func (c *MainController) Post() {
	msg:=c.GetString("message");
	nums:=c.Input().Get("nums");
	fmt.Println("msg",msg)
	fmt.Println("nums",nums)

	u := user{}
	// 调用  ParseForm 这个方法的时候，传入的参数必须为一个 struct 的指针
	if err := c.ParseForm(&u); err != nil {
		//handle error
		fmt.Println("err",err)
	}

	fmt.Println("u:",u)

	c.Ctx.WriteString("Post请求直接输出字符串。。。。。"+msg+","+nums)
}


