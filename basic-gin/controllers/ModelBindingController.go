package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type LoginForm2 struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 通过接收json参数   模型绑定参数  模拟登录
func LoginJSON(c *gin.Context)  {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User != "admin" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "登录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "登录成功"})
}

//  接收xml数据格式 模拟登录
func  LoginXML(c *gin.Context)  {
	var xml Login
	if err := c.ShouldBindXML(&xml); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if xml.User != "admin" || xml.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "登录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "登录成功"})
}

// 通过接收表单参数 模拟登录
func LoginForm(c *gin.Context)  {
	var form Login
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.User != "admin" || form.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "登录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "登录成功"})
}


//  获取请求路径中的参数
func GetParam(c *gin.Context)  {
	name:=c.Param("name")
	c.String(http.StatusOK,"Hello %s",name)
}

func GetParam2(c *gin.Context)  {
	name := c.Param("name")
	action := c.Param("action")
	c.JSON(http.StatusOK,gin.H{"code":200,"name":name,"action":action})
}