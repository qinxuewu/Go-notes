package controllers
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"fmt"
)

func FileHtml(c *gin.Context)  {
	c.HTML(http.StatusOK, "file.html", gin.H{
		"title": "文件上传",
	})
}


// 上传多个文件
func UploadFiles(c *gin.Context)  {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		log.Println(file.Filename)
		//保存文件到指定目录
		c.SaveUploadedFile(file, "static/upload/"+file.Filename)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

// 上传单个文件
func UploadFile(c *gin.Context)  {
	log.Println("进入上传文件方法..................")


	name:=c.PostForm("name")
	email:=c.PostForm("email")

	log.Println("name:",name)
	log.Println("email:",email)

	file, _ := c.FormFile("file")
	log.Println("file: " +
		"",file.Filename)

	c.SaveUploadedFile(file, "static/upload/"+file.Filename)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", "static/upload/"+file.Filename))
}