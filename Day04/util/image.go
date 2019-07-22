package main

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
	conf "we-blog/conf"
)

// 获取图片完整访问URL
func GetImageFullUrl(name string) string {
	config, _ := conf.GetConfig()
	return config.UploadConfig.ImagePrefixUrl + "/" + GetImagePath() + name
}

// 获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = EncodeMD5(fileName)

	return fileName + ext
}

// 获取图片路径
func GetImagePath() string {
	config, _ := conf.GetConfig()
	return config.UploadConfig.ImageSavePath
}

// 获取图片完整路径
func GetImageFullPath() string {
	config, _ := conf.GetConfig()
	return config.UploadConfig.RuntimeRootPath + GetImagePath()
}

// 检查图片后缀
func CheckImageExt(fileName string) bool {
	config, _ := conf.GetConfig()
	ext := GetExt(fileName)
	for _, allowExt := range config.UploadConfig.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// 检查图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := GetSize(f)
	if err != nil {
		log.Println(err)
		return false
	}
	config, _ := conf.GetConfig()
	return size <= config.UploadConfig.ImageMaxSize*1024*1024
}

// 检查图片
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
