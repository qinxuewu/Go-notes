package model

import (
	"crypto/md5"
	"encoding/hex"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type SysUsers struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"size:20"`
	Pwd       string `gorm:"size:100"`
	CreatedAt time.Time
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 登录验证
func CheckLogin(name, pwd string) SysUsers {
	var user SysUsers
	DB.Where("name = ? and pwd=? ", name, md5V(pwd)).First(&user)
	return user
}

//  更新密码
func UpdatePwd(id uint, pwd string) int64 {
	user := SysUsers{}
	// 主键查询
	DB.First(&user, id)
	count := DB.Model(&user).Update("pwd", md5V(pwd)).RowsAffected
	return count
}
