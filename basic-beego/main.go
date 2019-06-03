package main

import (
	"basic-beego/controllers"
	_ "basic-beego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	//注册自定义错误处理函数
	beego.ErrorController(&controllers.ErrorController{})

	log := logs.NewLogger(10000)  // 创建一个日志记录器，参数为缓冲区的大小
	log.SetLogger("console", "")  // 设置日志记录方式：控制台记录
	log.SetLevel(logs.LevelDebug) // 设置日志写入缓冲区的等级：Debug级别（最低级别，所以所有log都会输入到缓冲区）
	log.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）

	log.Emergency("Emergency")
	log.Alert("Alert")
	log.Critical("Critical")
	log.Error("Error")
	log.Warning("Warning")
	log.Notice("Notice")
	log.Informational("Informational")
	log.Debug("Debug")

	log.Close()


	beego.Run()
}

