package controllers

import (
	"github.com/astaxie/beego"
)

type RestController struct {
	beego.Controller
}

func (c *RestController) ListFood() {
	c.Ctx.WriteString("ListFood。。。。。")
}

func (c *RestController) CreateFood() {
	c.Ctx.WriteString("CreateFood。。。。。")
}

func (c *RestController) UpdateFood() {
	c.Ctx.WriteString("UpdateFood。。。。。")
}


func (c *RestController) DeleteFood() {
	c.Ctx.WriteString("DeleteFood。。。。。")
}

func (c *RestController) ApiFunc() {
	c.Ctx.WriteString("ApiFunc。。。。。")
}


func (c *RestController) GetFunc() {
	c.Ctx.WriteString("GetFunc。。。。。")
}

func (c *RestController) PostFunc() {
	c.Ctx.WriteString("PostFunc。。。。。")
}

