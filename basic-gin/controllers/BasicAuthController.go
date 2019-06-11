package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 对象的账户信息
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func BasicAuth(c *gin.Context)  {
	user := c.MustGet(gin.AuthUserKey).(string)
	if secret, ok := secrets[user]; ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}