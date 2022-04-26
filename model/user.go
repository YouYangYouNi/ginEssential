package model

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name      string
	Telephone string
	Password  string
}

func ExistTelephone(db *gorm.DB, telephone string) bool {
	var user Users
	result := db.Where("telephone = ?", telephone).First(&user)
	if result.RowsAffected < 0 {
		return false
	}
	if user.ID > 0 {
		return true
	}
	return false
}

func AddUser(db *gorm.DB, users Users) bool {
	result := db.Create(&users)
	if result.Error != nil {
		return false
	}
	return true
}
