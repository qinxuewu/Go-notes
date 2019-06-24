package main


import (
"fmt"
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义全局db变量
var db *gorm.DB

// 初始化db
func init()  {
	db, _=gorm.Open("mysql", "root:870439570@tcp(39.108.144.143:3306)/test?charset=utf8&parseTime=True&loc=Local")
}


type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	defer db.Close()

	// 检测指定表是否存在
	falg:=db.HasTable("products")
	fmt.Println("指定表表名是否存在: ",falg)

	// 自动迁移模式
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code:"123",Price:1000})


	// 读取
	var product Product
	db.First(&product,1) // 查询ID为1的数据
	fmt.Println("查询ID为1的数据: ",product)

	db.First(&product,"code = ?","123")  // 查询code=123
	fmt.Println("查询code=123的数据: ",product)


	// 更新 上面查询出的数据
	db.Model(&product).Update("Price",2000)
	fmt.Println("更新 上面查询出的数据: ",product)

	// 删除
	db.Delete(&product)

}



func check(err error)  {
	if err !=nil{
		panic(err)
	}
}
