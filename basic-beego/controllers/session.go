/**
 * session 控制
 * @author qinxuewu
 * @create 19/6/1下午6:24
 * @since 1.0.0
 */
package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type SessionController struct {
	beego.Controller
}

func (c *SessionController) Get(){
	v:= c.GetSession("asta")
	if v == nil {
		c.SetSession("asta", int(1))
		c.Data["num"] = 0
	} else {
		c.SetSession("asta", v.(int)+1)
		c.Data["num"] = v.(int)
	}
	fmt.Println(v)

	c.Ctx.WriteString("SessionController。。。。.")
}
