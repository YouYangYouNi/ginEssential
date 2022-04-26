package controller

import (
	"essential/common"
	"essential/model"
	"essential/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号不正确",
		})
		return
	}
	db := common.GetDB()
	if model.ExistTelephone(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户已存在",
		})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码小于6位",
		})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	//创建用户
	newUser := model.Users{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}

	result := model.AddUser(db, newUser)
	code := 200
	msg := "注册成功：" + name
	if !result {
		code = http.StatusUnprocessableEntity
		msg = "注册失败：" + name
	}
	ctx.JSON(code, gin.H{
		"msg": msg,
	})

}
