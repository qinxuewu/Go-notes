package main


import (
	"github.com/astaxie/beego/orm"
)

type UserInfo struct {
	Id          int
	Name        string
	Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
	Post    	[]*Post `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id          int
	Age         int16
	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`	//设置一对多关系

}

func init() {
	//设置数据库
	// 参数1        数据库的别名，用来在ORM中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	orm.RegisterDataBase("default","mysql","root:870439570@tcp(39.108.144.143:3306)/test2?charset=utf8",30)
	// 需要在init中注册定义的model
	orm.RegisterModel(new(UserInfo), new(Profile),new(Post))
	//打印调试
	orm.Debug = true
	//创建table
	orm.RunSyncdb("default",false,true)
}

func main()  {
	
}