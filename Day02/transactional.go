package main


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)
type Emoinfo struct {
	ID        uint `gorm:"primary_key"`
	Name string
	Age int
}

func main()  {
	dbs,_:=gorm.Open("mysql", "root:870439570@tcp(39.108.144.143:3306)/test?charset=utf8&parseTime=True&loc=Local")
	// 启用Logger，显示详细日志
	dbs.LogMode(true)
	// Ping
	dbs.DB().Ping()
	//  连接池
	dbs.DB().SetMaxIdleConns(10)
	dbs.DB().SetMaxOpenConns(100)

	defer dbs.Close()
	// 检测指定表是否存在
	falg:=dbs.HasTable(&Emoinfo{})
	fmt.Println("指定表表名是否存在: ",falg)

	// 自动迁移模式
	dbs.AutoMigrate(&Emoinfo{})

	emp := Emoinfo{Name: "qxw", Age: 18}
	dbs.Create(&emp)

	//test(dbs);
	CreateAnimals(dbs)
}

func test(db *gorm.DB)  {
	// 开始事务
	tx := db.Begin()

	// 在事务中做一些数据库操作（从这一点使用'tx'，而不是'db'）
	db.Save(&Emoinfo{Name:"qxw",Age:26})
	db.Save(&Emoinfo{Name:"qxw888",Age:28})

	// 发生错误时回滚事务
	tx.Rollback()

	// 或提交事务
	tx.Commit()
}

func CreateAnimals(db *gorm.DB) error {
	tx := db.Begin()
	// 注意，一旦你在一个事务中，使用tx作为数据库句柄

	if err := tx.Create(&Emoinfo{Name: "Giraffe",Age:18}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Emoinfo{Name: "Lion",Age:20}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}


//  原生sql构建
func SqlTest(db *gorm.DB)  {

	db.Exec("drop table emoinfos;")
	db.Exec("update emoinfos set name='sqlname' where id in (?)","666",[]int64{11,22,33})

	db.Raw("select id,name,age from emoinfos where id = ?",3).Scan(&Emoinfo{})




	rows, _ := db.Model(&Emoinfo{}).Where("name = ?", "jinzhu").Select("name, age").Rows()
	for rows.Next() {

	}
	defer rows.Close()

	// 调试单个操作，显示此操作的详细日志
	db.Debug().Where("name = ?", "jinzhu").First(&Emoinfo{})
}
