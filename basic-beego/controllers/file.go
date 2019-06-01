/**
 * 文件上传
    配置文件设置文件大小: beego.MaxMemory = 1<<22
 * @author qinxuewu
 * @create 19/5/31下午10:08
 * @since 1.0.0
 */
package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"fmt"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) Get()  {
	c.Data["Website"] = "beego实现文件上传"
	c.TplName = "file.tpl"
}
/**
Beego 提供了两个很方便的方法来处理文件上传：

   GetFile(key string) (multipart.File, *multipart.FileHeader, error)
   该方法主要用于用户读取表单中的文件名 the_file，然后返回相应的信息，用户根据这些变量来处理文件上传：过滤、保存文件等。

   SaveToFile(fromfile, tofile string) error
   该方法是在 GetFile 的基础上实现了快速保存的功能 fromfile 是提交时候的 html 表单中的 name
 */
func (c *FileController) Uplod()  {
	fmt.Println("进入上传文件方法。。。。。。。。")

	f,h,err:=c.GetFile("uploadname") ////获取上传的文件
	if err !=nil{
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	// 保存位置在 static/upload, 没有文件夹要先创建
	c.SaveToFile("uploadname", "static/upload/" + h.Filename)

	c.Ctx.WriteString("文件上传结束。。。。")
}