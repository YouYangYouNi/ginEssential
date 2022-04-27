package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string
}
type Tabler interface {
	TableName() string
}

// TableName 会将 User 的表名重写为 `profiles`
func (Category) TableName() string {
	return "category"
}