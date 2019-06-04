package main


import (
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
	"github.com/astaxie/beego/orm"
	"fmt"

)
type Student struct {
	Id int
	Name string `orm:"size(100)"`
	Age         int16
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
	orm.RegisterModel(new(Student))
	//打印调试
	orm.Debug = true
	//创建table
	orm.RunSyncdb("default",false,true)
}

func main()  {
	//o:=orm.NewOrm()
	//// 获取 QuerySeter 对象，user 为表名
	//qs:=o.QueryTable("student")
	//
	////返回总数
	//cnt,err:=qs.Count()
	//checkErr(err)
	//fmt.Printf("Count Num: %d",cnt)

	//findOne()
	//findMaps()

	//tes1()
	test2()

}

func findOne()  {
	var student Student
	o:=orm.NewOrm()
	err := o.QueryTable("student").Filter("name", "slene").One(&student)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Printf("Not row found")
	}
}
func findMaps()  {
	o:=orm.NewOrm()
	var maps []orm.Params
	num, err := o.QueryTable("student").Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range maps {
			fmt.Println(m["Id"], m["Name"])
		}
	}

}



// 使用SQL语句进行查询 返回 sql.Result 对象
func tes1()  {
	o:=orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("UPDATE student SET name = ? WHERE id = ?", "sdsdsfsf",1)
	res,err:=r.Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}


}
// QueryRow 和 QueryRows 提供高级 sql mapper 功能
func test2()  {
	var student Student
	o:=orm.NewOrm()
	err:=o.Raw("select id,name,age from student where id=?",1).QueryRow(&student)
	checkErr(err)

	fmt.Println("student:",student)

//	 QueryRows 支持的对象还有 map 规则是和 QueryRow 一样的，但都是 slice

	var students []Student
	num,err:=o.Raw("select id,name,age from student").QueryRows(&students)
	fmt.Println("user nums: ", num)
	fmt.Println("students: ", students)

}

func test3()  {
	o:=orm.NewOrm()
	err :=o.Begin()
	// 事务处理过程

	students:=Student{Name:"qxwxqxqq",Age:18}
	//插入表
	id,err:=o.Insert(&students)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	////更新整个表
	students.Name="fssfsfsf"
	num,err:=o.Update(&students)
	if num==0 {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
}




func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}