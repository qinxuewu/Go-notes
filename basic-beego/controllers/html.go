package controllers

import (
	"github.com/astaxie/beego"
)

type HtmlController struct {
	beego.Controller
}


func (c *HtmlController) Get() {
	c.Data["Title"] = "beego的模板语法学习"
	c.Data["IsLogin"] = true
	c.Data["IsHome"] = true
	c.Data["IsAbout"] = true

	pages := []struct {
		Num int
	}{{10}, {20}, {30}}

	c.Data["Total"] = 100
	c.Data["Pages"] = pages

	c.TplName = "test.html"
}


