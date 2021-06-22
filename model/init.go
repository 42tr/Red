package model

import (
	"red/util"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库连接单例
var DB *gorm.DB

// Database 在中间件中初始化 mysql 连接
func Database(connStr string) {
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		util.Log().Panic("数据库连接失败", err)
	}
	db.LogMode(true)
	// set connection pool
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}
