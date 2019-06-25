/**
 *
 * @author qinxuewu
 * @create 19/6/25下午8:28
 * @since 1.0.0
 */
package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Users struct {
	ID        uint `gorm:"primary_key"`
	Name string
	Age int
}



func main()  {
	dbs,_:=gorm.Open("mysql", "root:870439570@tcp(39.108.144.143:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer dbs.Close()
	// 检测指定表是否存在
	falg:=dbs.HasTable(&Users{})
	fmt.Println("指定表表名是否存在: ",falg)

	// 自动迁移模式
	dbs.AutoMigrate(&Users{})

	user := Users{Name: "Jinzhu", Age: 18}
	dbs.Create(&user)

	// 获取第一条记录，按主键排序
	dbs.First(&user)
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	fmt.Println("获取第一条记录，按主键排序:",user)

	// 获取最后一条记录，按主键排序
	dbs.Last(&user)
	fmt.Println("获取最后一条记录，按主键排序:",user)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// 获取所有记录
	dbs.Find(&user)
	//// SELECT * FROM users;




	// 使用主键获取记录
	dbs.First(&user, 1)
	fmt.Println("使用主键获取记录:",user)

	//// SELECT * FROM users WHERE id = 1;

	where(dbs)
	
}

func where(db *gorm.DB)  {
	user := Users{}
	// 获取第一个匹配记录
	db.Where("name = ?", "jinzhu").First(&user)
	//// SELECT * FROM users WHERE name = 'jinzhu' limit 1;

	// 获取所有匹配记录
	db.Where("name = ?", "jinzhu").Find(&user)
	fmt.Println("获取所有匹配记录  :",user)
	//// SELECT * FROM users WHERE name = 'jinzhu';

	db.Where("name <> ?", "jinzhu").Find(&user)

	// IN
	db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&user)

	// LIKE
	db.Where("name LIKE ?", "%jin%").Find(&user)
	fmt.Println("LIKE:",user)
	// AND
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&user)

	// Time
	db.Where("age > ?").Find(&user)

}


