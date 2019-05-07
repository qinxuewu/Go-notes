
package main

import (
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Userinfo struct {
	Id            string `orm:"column(Uid);pk"` // 设置主键
	Username    string
	Departname  string
	Created     time.Time
}

type User struct {
	Id            string `orm:"column(Uid);pk"` // 设置主键
	Name        string
	Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
	Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id          int
	Age         int16
	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
	Tags  []*Tag `orm:"rel(m2m)"`    //设置一对多关系
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

/**
  beego orm针对驼峰命名会自动帮你转化成下划线字段，例如你定义了Struct名字为UserInfo，
  那么转化成底层实现的时候是user_info，字段命名也遵循该规则
 */

func init() {

	//设置数据库
	orm.RegisterDataBase("default","mysql","root:870439570@tcp(39.108.144.143:3306)/test?charset=utf8",30)
	//注册定义的model   可以注册多个
	orm.RegisterModel(new(Userinfo),new(User), new(Profile), new(Tag))

	//创建table
	orm.RunSyncdb("default",false,true)

}

func main()  {
	insert()
}
func insert()  {
	o := orm.NewOrm()
	var user User
	user.Name = "zxxx"

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}
	
}
