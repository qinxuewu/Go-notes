/**
 *
 * @author qinxuewu
 * @create 19/6/24下午8:12
 * @since 1.0.0
 */
package main


import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

type User struct {
	gorm.Model
	Name string
	Age int
	Description string
} 
func main()  {
	mybdb,_:=gorm.Open("mysql", "root:870439570@tcp(39.108.144.143:3306)/test?charset=utf8&parseTime=True&loc=Local")

	// 检查表是否存在
	falg:=mybdb.HasTable(&User{})
	fmt.Println(falg)

	if falg == false {
		fmt.Println("创建表逻辑")
		// 为模型`User`创建表
		//mybdb.CreateTable(&User{})

		// 创建表`users'时将“ENGINE = InnoDB”附加到SQL语句
		mybdb.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})

		updateColumn(mybdb)

		deleteColumn(mybdb)

		addIndex(mybdb)
	}else {

		fmt.Println("删除表逻辑")
		// 删除模型`User`的表
		mybdb.DropTable(&User{})

		// 删除表`users`
		//mybdb.DropTable("users")

		// 删除模型`User`的表和表`products`
		mybdb.DropTableIfExists(&User{}, "products")
	}


}

// 删除列
func deleteColumn(db *gorm.DB)  {
	// 删除指定模型的列
	db.Model(&User{}).DropColumn("description")
}

// 修改列

func updateColumn(db *gorm.DB)  {
	// 修改模型`User`的description列的数据类型为`text`
	db.Model(&User{}).ModifyColumn("description", "text")
}


// 添加外建

func addForeigKey(db *gorm.DB)  {
	// 添加主键
	// 1st param : 外键字段
	// 2nd param : 外键表(字段)
	// 3rd param : ONDELETE
	// 4th param : ONUPDATE
	db.Model(&User{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")

}

func  addIndex(db *gorm.DB)  {

	// 为`age`列添加索引
	db.Model(&User{}).AddIndex("idex_age","age")

	// 添加唯一索引
	db.Model(&User{}).AddUniqueIndex("idx_user_name","name")


	// 为多列添加唯一索引
	db.Model(&User{}).AddUniqueIndex("idx_user_name_age", "name", "age")

	// 删除索引
	db.Model(&User{}).RemoveIndex("idx_user_name")
}
