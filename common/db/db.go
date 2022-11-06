package db

import (
	"dice/common/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBSetUp(addr string) {
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	global.DBMysql = db
}
