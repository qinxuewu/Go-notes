package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"./model"
)



type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	defer model.DB.Close()

	// 检测指定表是否存在
	falg:=model.DB.HasTable("products")
	fmt.Println("指定表表名是否存在: ",falg)

	// 自动迁移模式
	model.DB.AutoMigrate(&Product{})

	model.DB.Create(&Product{Code:"123",Price:1000})

	// 读取
	var product Product
	model.DB.First(&product,1) // 查询ID为1的数据
	fmt.Println("查询ID为1的数据: ",product)

	model.DB.First(&product,"code = ?","123")  // 查询code=123
	fmt.Println("查询code=123的数据: ",product)


	// 更新 上面查询出的数据
	model.DB.Model(&product).Update("Price",2000)
	fmt.Println("更新 上面查询出的数据: ",product)

}


