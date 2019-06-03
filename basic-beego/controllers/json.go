package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type JsonController struct {
	beego.Controller
}

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}


//  输出json数据   http://localhost:8080/object/outjson
func (c *JsonController) Outjson(){
	fmt.Println("Outjson.............")
	c.Data["json"]=&Server{"qxw.club","192.168.1.1"}
	c.ServeJSON()
}

//  jsonp调用  http://localhost:8080/object/outjsonp
func (c *JsonController) Outjsonp(){
	fmt.Println("Outjsonp.............")
	c.Data["jsonp"]=&Server{"jsonp","192.168.1.1"}
	c.ServeJSON()
}


