package common

import (
	"essential/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/essential?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	err = db.AutoMigrate(&model.Users{})
	if err != nil {
		return nil
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
