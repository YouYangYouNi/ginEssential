package controller

import (
	"essential/common"
	"essential/model"
	"essential/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ICategory interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategory {
	db := common.GetDB()
	err := db.AutoMigrate(&model.Category{})
	if err != nil {
		return nil
	}
	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	err := ctx.Bind(&requestCategory)
	if err != nil {
		panic(err)
	}
	if requestCategory.Name == "" {
		response.ApiFail(ctx, gin.H{}, "分类名称必填，请修改")
	}
	c.DB.Create(&requestCategory)
	response.ApiSuccess(ctx, gin.H{"category": requestCategory}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	panic("implement me ，请修改")
}

func (c CategoryController) Show(ctx *gin.Context) {
	panic("implement me")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	panic("implement me")
}
