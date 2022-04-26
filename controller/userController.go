package controller

import (
	"essential/common"
	"essential/model"
	"essential/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "系统错误",
		})
		return
	}

	//创建用户
	newUser := model.Users{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
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

func Login(ctx *gin.Context) {
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "手机号不正确",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "密码小于6位",
		})
		return
	}

	db := common.GetDB()
	user := model.Users{}
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "用户不存在",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if user.ID == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": "密码错误",
			})
			return
		}
	}
	token := 111
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})
	return
}
