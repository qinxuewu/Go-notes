package main
import (
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
	"github.com/astaxie/beego/orm"
	"fmt"
	"log"
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
	// 参数1        数据库的别名，用来在ORM中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	orm.RegisterDataBase("default","mysql","root:870439570@tcp(39.108.144.143:3306)/test2?charset=utf8",30)
	//注册定义的model   可以注册多个
	orm.RegisterModel(new(User))
	//打印调试
	orm.Debug = true
	//创建table
	orm.RunSyncdb("default",false,true)
}




func main()  {
	//o:=orm.NewOrm()
	//
	//user:=User{Name:"qxwxqxqq"}
	//
	////插入表
	//id,err:=o.Insert(&user)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)
	//
	////更新整个表
	//user.Name="fssfsfsf"
	//num,err:=o.Update(&user)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// // 指定字段更新   o.Update(&user, "Field1", "Field2", ...)


	//
	////读取one
	//u:=User{Id:user.Id}
	//err=o.Read(&u)
	//fmt.Printf("ERR: %v\n", err)
	//
	//
	//
	////删除表
	////num,err=o.Delete(&u)
	////fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//
	//
	//// 使用原生sql
	//var r orm.RawSeter
	//r = o.Raw("UPDATE user SET name = ? WHERE name = ?", "testing", "fssfsfsf")
	//r.Exec()
	////fmt.Printf("%d\n", r)

	//ReadOrCreate()

	InsertMulti();

}

func ReadOrCreate()  {
	//	 尝试从数据库读取，不存在的话就创建一个
	o:=orm.NewOrm()
	userSaveUpdae:=User{Name:"slene"}
	if created,id,err:=o.ReadOrCreate(&userSaveUpdae,"Name");err==nil{
		if created{
			fmt.Println("不存在新增一条数据",id)
		}else{
			fmt.Println("存在 直接查询返回",id)

		}
	}
}

// 批量插入
func InsertMulti()  {
	o:=orm.NewOrm()
	users := []User{
		{Name: "slene"},
		{Name: "astaxie"},
		{Name: "unknown"},
	}
	successNums, err := o.InsertMulti(100, users)
	if err !=nil{
		log.Println(err)
	}
	fmt.Println(successNums)
}


