package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"we-blog/conf"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

var DB *gorm.DB

func init() {
	config, _ := conf.GetConfig()
	c := config.DatasourceConfig
	DB = Open(c.Host, c.Port, c.User, c.Pwd, c.Dbname)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return c.TablePrefix + defaultTableName
	}

	// 启用Logger，显示详细日志
	DB.LogMode(true)
	//  连接池
	DB.DB().SetMaxIdleConns(c.Idles)
	DB.DB().SetMaxOpenConns(c.Connections)

	// 新增的回调
	DB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// 修改的回调
	DB.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// 删除的回调
	DB.Callback().Delete().Replace("gorm:delete", deleteCallback)

	//// 自动迁移模式
	//DB.AutoMigrate(
	//	&SysUsers{},
	//	&Category{},
	//)

	// 禁用表明复数
	DB.SingularTable(true)
}

func CloseDB() {
	defer DB.Close()
}

func Open(host, port, username, password, name string) *gorm.DB {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		name,
	)
	db, err := gorm.Open("mysql", dbConfig)
	if err != nil {
		log.Fatalf("【初始数据库连接失败..】 %v ", err)
	}
	return db
}

// 回调钩子
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// 回调钩子
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// 根据入参获取设置了字面值的参数
	// 假设没有指定 update_column 的字段，我们默认在更新回调设置 ModifiedOn 的值
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// 软删除
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		// 检查是否手动指定了delete_option
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		//  获取我们约定的删除字段，若存在则 UPDATE 软删除，若不存在则 DELETE 硬删除
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				// 该方法可以添加值作为SQL的参数，也可用于防范SQL注入
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(), // 返回引用的表名
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
