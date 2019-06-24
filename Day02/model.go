/**
 *  模型的定义
 * @author qinxuewu
 * @create 19/6/24下午8:39
 * @since 1.0.0
 */
package main

import (
	"github.com/jinzhu/gorm"
	"time"
	"database/sql"
)

// 表名是结构体名称的复数形式    默认表名是`deptnames`
type Deptname struct {
	gorm.Model
	Birthday     time.Time
	Age          int
	Name         string  `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num          int     `gorm:"AUTO_INCREMENT"` // 自增

	CreditCard        CreditCard      // One-To-One (拥有一个 - CreditCard表的UserID作外键)
	Emails            []Email         // One-To-Many (拥有多个 - Email表的UserID作外键)

	BillingAddress    Address         // One-To-One (属于 - 本表的BillingAddressID作外键)
	BillingAddressID  sql.NullInt64

	ShippingAddress   Address         // One-To-One (属于 - 本表的ShippingAddressID作外键)
	ShippingAddressID int

	IgnoreMe          int `gorm:"-"`   // 忽略这个字段
	Languages         []Language `gorm:"many2many:user_languages;"` // Many-To-Many , 'user_languages'是连接表
}


type Email struct {
	ID      int
	UserID  int     `gorm:"index"` // 外键 (属于), tag `index`是为该列创建索引
	Email   string  `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // 设置字段为非空并唯一
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // 创建索引并命名，如果找到其他相同名称的索引则创建组合索引
	Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}

type CreditCard struct {
	gorm.Model
	UserID  uint
	Number  string
}


type Emp struct {
	ID uint             // 列名为 `id`  字段`ID`为默认主键
	Name string         // 列名为 `name`
	Birthday time.Time  // 列名为 `birthday`
	CreatedAt time.Time // 列名为 `created_at`
}


// 重设列名
type Animal struct {
	AnimalId     int64     `gorm:"primary_key"` // 设置AnimalId为主键
	Birthday    time.Time `gorm:"column:day_of_the_beast"` // 设置列名为`day_of_the_beast`
	Age         int64     `gorm:"column:age_of_the_beast"` // 设置列名为`age_of_the_beast`
}


// 自定义表名称

func (Deptname) TableName() string  {

	return "deptname"
}




func main()  {
	modeldDb,_:=gorm.Open("mysql", "root:870439570@tcp(39.108.144.143:3306)/test?charset=utf8&parseTime=True&loc=Local")

	// 全局禁用表名复数
	//modeldDb.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响


	// 自动迁移模式
	modeldDb.CreateTable(&Deptname{},&Email{},&Address{},&Language{},&CreditCard{},&Emp{},&Animal{})
}
