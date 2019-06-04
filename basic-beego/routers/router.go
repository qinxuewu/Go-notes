package routers

import (
	"basic-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 固定路由也就是全匹配的路由
    beego.Router("/", &controllers.MainController{})

	beego.Router("/html", &controllers.HtmlController{})

    //固定路由 指定访问的请求类型  http://localhost:8080/admin
	beego.Router("/admin", &controllers.MainController{},"get:Get;post:Post")

	//固定路由 只允许post请求   http://localhost:8080/admin/post
	beego.Router("/admin/post", &controllers.MainController{},"post:Post")

    // 正则路由

    // 匹配 http://localhost:8080/api/123 此时变量”:id”值为”123”
    beego.Router("/api/?:id",&controllers.SinatraController{},"get:FindId")

    //  匹配  http://localhost:8080/api/123    此时变量”:id”值为”123”    但URL”/api/“匹配时  没有ID
	beego.Router("/api/:id",&controllers.SinatraController{},"get:FindId")


    /**

    beego.Router(“/api/:id([0-9]+)“, &controllers.RController{})
	自定义正则匹配 //例如对于URL”/api/123”可以匹配成功，此时变量”:id”值为”123”


    beego.Router(“/user/:username([\\w]+)“, &controllers.RController{})
	正则字符串匹配 //例如对于URL”/user/astaxie”可以匹配成功，此时变量”:username”值为”astaxie”


    beego.Router(“/download/*.*”, &controllers.RController{})

    匹配方式 //例如对于URL”/download/file/api.xml”可以匹配成功，此时变量”:path”值为”file/api”， “:ext”值为”xml”


    beego.Router(“/download/ceshi/*“, &controllers.RController{})
    全匹配方式 //例如对于URL”/download/ceshi/file/api.json”可以匹配成功，此时变量”:splat”值为”file/api.json”


    beego.Router(“/:id:int”, &controllers.RController{})
    int 类型设置方式，匹配 :id为int 类型，框架帮你实现了正则 ([0-9]+)

	beego.Router(“/:hi:string”, &controllers.RController{})
	string 类型设置方式，匹配 :hi 为 string 类型。框架帮你实现了正则 ([\w]+)

	beego.Router(“/cms_:id([0-9]+).html”, &controllers.CmsController{})
	带有前缀的自定义正则 //匹配 :id 为正则类型。匹配 cms_123.html 这样的 url :id = 123


     */




     /**

     	自定义方法及 RESTful 规则
     	beego.Router("/",&IndexController{},"*:Index")
     		* 表示任意的 method 都执行该函数
              使用 httpmethod:funcname 格式来展示
              多个不同的格式使用 ; 分割
              多个 method 对应同一个 funcname，method 之间通过 , 来分割
      */

    //  http://localhost:8080/api/list
	beego.Router("/list",&controllers.RestController{},"*:ListFood")
	beego.Router("/rest/create",&controllers.RestController{},"post:CreateFood")
	beego.Router("/rest/update",&controllers.RestController{},"put:UpdateFood")
	beego.Router("/rest/delete",&controllers.RestController{},"delete:DeleteFood")


	//  多个 HTTP Method 指向同一个函数的示例
	beego.Router("/api/getpost",&controllers.RestController{},"get,post:ApiFunc")

	// 不同的 method 对应不同的函数，通过 ; 进行分割的示例：
	beego.Router("/simple",&controllers.RestController{},"get:GetFunc;post:PostFunc")

	beego.Router("/simplePost",&controllers.RestController{},"post:PostFunc")


	//自动路由配置  beego 就会通过反射获取该结构体中所有的实现方法
	/**
			http://localhost:8080/object/login
	        http://localhost:8080/object/logout
	 */
	beego.AutoRouter(&controllers.ObjectController{})

	beego.AutoRouter(&controllers.JsonController{})
	// 注册注解路由
	/**
	  http://localhost:8080/staticblock/key13
	  http://localhost:8080/all/123
	 */
	beego.Include(&controllers.CMSController{})


	// 文件上传
	//  http://localhost:8080/file
	beego.Router("/file", &controllers.FileController{},"get:Get;post:Uplod")


	// http://localhost:8080/session
	beego.Router("/session",&controllers.SessionController{})


}
