package filter

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	util "we-blog/util"
)

//  JWT 中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = util.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = util.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = util.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = util.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != util.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  util.GetMsg(code),
				"data": data,
			})
			log.Fatal("token验证不通过")
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			return
		}

		c.Next()
	}
}
