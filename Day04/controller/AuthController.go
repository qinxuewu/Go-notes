package controller

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	models "we-blog/model"
	"we-blog/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := util.INVALID_PARAMS
	if ok {
		// 验证账号密码是否存在
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = util.ERROR_AUTH_TOKEN
			} else {
				// 返回 token
				data["token"] = token
				code = util.SUCCESS
			}

		} else {
			code = util.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			// 变量异常
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  util.GetMsg(code),
		"data": data,
	})
}
