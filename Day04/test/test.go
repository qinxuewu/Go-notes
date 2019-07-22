package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "we-blog/model"
)

// http://gorm.book.jasperxu.com/advanced.html#sb
// https://eddycjy.gitbook.io/golang/di-3-ke-gin/api-01
func main() {
	db, _ := gorm.Open("mysql", "root:root@tcp(localhost:3306)/blog?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.Debug()
	// 启用Logger，显示详细日志
	db.LogMode(true)
	var art Article

	db.Table("blog_article").Select("id,title").Where("id =?", 1).Find(&art)

	// DB
	var list []Article
	DB.Raw("SELECT * FROM blog_article").Scan(&list)
	fmt.Println(len(list))
}
