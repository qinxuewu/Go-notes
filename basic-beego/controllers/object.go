package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type ObjectController struct {
	beego.Controller
}

//自动路由配置  beego 就会通过反射获取该结构体中所有的实现方法

func (c *ObjectController) Login()  {
	c.Ctx.WriteString("Login............")
}


func (c *ObjectController) Logout()  {
	c.Ctx.WriteString("Logout............")
}

//

/**
	当访问： http://localhost:8080/object/blog/2013/09/12
  除了前缀两个 /:controller/:method 的匹配之外，剩下的 url beego 会帮你自动化解析为参数，保存在 this.Ctx.Input.Params 当中
 */
func (c *ObjectController)  Blog() {
	var mapValue=c.Ctx.Input.Params()
	fmt.Printf("%v\n", mapValue)
	c.Ctx.WriteString("Blog............"+mapValue["0"])
}

