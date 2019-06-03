package main
import (
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
	"github.com/astaxie/beego/orm"
	"fmt"
)

type User struct {
	Id int
	Name string `orm:"size(100)"`
}
/**
	beego orm是一个Go进行ORM操作的库，它采用了Go style方式对数据库进行操作，
	实现了struct到数据表记录的映射
 */
func init() {

	//设置数据库
	orm.RegisterDataBase("default","mysql","root:870439570@tcp(39.108.144.143:3306)/test2?charset=utf8",30)
	//注册定义的model   可以注册多个
	orm.RegisterModel(new(User))
	//打印调试
	orm.Debug = true
	//创建table
	orm.RunSyncdb("default",false,true)
}

func main()  {
	o:=orm.NewOrm()

	user:=User{Name:"qxwxqxqq"}

	//插入表
	id,err:=o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	//更新表
	user.Name="fssfsfsf"
	num,err:=o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	//读取one

	u:=User{Id:user.Id}
	err=o.Read(&u)
	fmt.Printf("ERR: %v\n", err)



	//删除表
	//num,err=o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)



	// 使用原生sql
	var r orm.RawSeter
	r = o.Raw("UPDATE user SET name = ? WHERE name = ?", "testing", "fssfsfsf")
	r.Exec()
	//fmt.Printf("%d\n", r)

}



