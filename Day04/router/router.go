package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "we-blog/controller"
	"we-blog/util"
)

// 路由配置
func InitRouter() *gin.Engine {
	router := gin.Default()
	// 配置上传图片访问路径
	router.StaticFS("/upload/images", http.Dir(util.GetImageFullPath()))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/auth", GetAuth)
	router.POST("/upload", UploadImage)

	//以下的接口，都使用Authorize()中间件身份验证
	//router.Use(JWT())
	apiv1 := router.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", GetTags)
		//新建标签
		apiv1.POST("/tags", AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", GetArticle)
		//新建文章
		apiv1.POST("/articles", AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", DeleteArticle)
	}
	return router
}
