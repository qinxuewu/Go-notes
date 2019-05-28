package controllers

import (
	"github.com/astaxie/beego"
)

type SinatraController struct {
	beego.Controller
}

func (c *SinatraController) FindId() {
	var id=c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("FindId。。。。。"+id)
}


