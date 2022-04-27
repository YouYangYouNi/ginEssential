package common

import (
	"essential/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("datasource.username"),
		viper.GetString("datasource.password"),
		viper.GetString("datasource.host"),
		viper.GetString("datasource.port"),
		viper.GetString("datasource.database"),
		viper.GetString("datasource.charset"),
	)
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
