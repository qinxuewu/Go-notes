package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Category struct {
	Model
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:20"`
}

// 添加分类
func SaveCategory(c Category) bool {
	DB.Create(c)
	return true
}

// 删除分类
func DelCategory(id uint) bool {
	DB.Where("id = ?", id).Delete(&Category{})
	return true
}

// 硬删除代码  要使用 Unscoped()，这是 GORM 的约定
func CleanAllCategory() bool {
	DB.Unscoped().Where("deleted_on != ? ", 0).Delete(&Category{})
	return true
}
