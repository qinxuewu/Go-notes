package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "第一个beego程序"
	c.Data["Email"] = "870439570@qq.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.Ctx.WriteString("Post请求直接输出字符串。。。。。")
}

