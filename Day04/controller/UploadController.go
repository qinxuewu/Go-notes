package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	util "we-blog/util"
)

func UploadImage(c *gin.Context) {
	code := util.SUCCESS
	data := make(map[string]string)

	// 获取上传的图片（返回提供的表单键的第一个文件）
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		log.Fatal(err)
		code = util.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  util.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = util.INVALID_PARAMS
	} else {
		imageName := util.GetImageName(image.Filename)
		fullPath := util.GetImageFullPath()
		savePath := util.GetImagePath()

		src := fullPath + imageName
		// 检查图片大小，检查图片后缀
		if !util.CheckImageExt(imageName) || !util.CheckImageSize(file) {
			code = util.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			// 检查上传图片所需（权限、文件夹）
			err := util.CheckImage(fullPath)
			if err != nil {
				log.Fatal(err)
				code = util.ERROR_UPLOAD_CHECK_IMAGE_FAIL
				// 保存图片
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				log.Fatal(err)
				code = util.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = util.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  util.GetMsg(code),
		"data": data,
	})
}
