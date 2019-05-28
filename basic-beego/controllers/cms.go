package controllers
import (
	"github.com/astaxie/beego"
)
// CMS API
type CMSController struct {
	beego.Controller
}


func (c *CMSController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
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

