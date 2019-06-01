package controllers
import (
	"github.com/astaxie/beego"
	"fmt"
)
// CMS API
type CMSController struct {
	beego.Controller
}


func (c *CMSController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
	c.Mapping("Bind", c.Bind)
}



// @router /staticblock/:key [get]
func (this  *CMSController) StaticBlock() {
	var key=this .Ctx.Input.Param(":key")
	this .Ctx.WriteString("StaticBlock。。。。。"+key)
}



// @router /all/:key [get]
func (this  *CMSController) AllBlock() {
	var key=this.Ctx.Input.Param(":key")
	this .Ctx.WriteString("AllBlock。。。。。"+key)
}



// @router /bind/:key [get]
func (this *CMSController) Bind()  {
	fmt.Println("进入bind方法...............")

	/**
	支持从用户请求中直接数据 bind 到指定的对象
	 url:  ?id=123&isok=true&ft=1.2&ol[0]=1&ol[1]=2&ul[]=str&ul[]=array&user.Name=astaxie
 */


	var id int
	this.Ctx.Input.Bind(&id,"id")
	fmt.Println("id:",id)

	var isok bool
	this.Ctx.Input.Bind(&isok,"isok")
	fmt.Println("isok:",isok)

	ol :=make([]int,0,2)
	this.Ctx.Input.Bind(&ol,"ol")
	fmt.Println("ol",ol)

	ul:=make([]string,0,2)
	this.Ctx.Input.Bind(&ul,"ul")
	fmt.Println("ul:",ul)

	this.Ctx.WriteString("支持从用户请求中直接数据 bind 到指定的对象")
}
