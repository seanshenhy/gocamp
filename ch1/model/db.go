package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var DB *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"123456",
		"localhost",
		"3306",
		"test",
	)), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	DB = db
	// defer sqlDB.Close()
}
