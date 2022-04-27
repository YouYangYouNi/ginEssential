package controller

import (
	"essential/common"
	"essential/dto"
	"essential/model"
	"essential/response"
	"essential/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		response.ApiFail(ctx, nil, "手机号不正确")
		return
	}
	db := common.GetDB()
	if model.ExistTelephone(db, telephone) {
		response.ApiFail(ctx, nil, "用户已存在")
		return
	}

	if len(password) < 6 {
		response.ApiFail(ctx, nil, "密码小于6位")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.ApiFail(ctx, nil, "系统错误")
		return
	}

	//创建用户
	newUser := model.Users{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}

	result := model.AddUser(db, newUser)
	if !result {
		response.ApiFail(ctx, nil, "注册失败")
	}
	response.ApiSuccess(ctx, nil, "注册成功"+name)
}

func Login(ctx *gin.Context) {
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		response.ApiFail(ctx, nil, "手机号不正确")
		return
	}
	if len(password) < 6 {
		response.ApiFail(ctx, nil, "密码小于6位")
		return
	}

	db := common.GetDB()
	user := model.Users{}
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.ApiFail(ctx, nil, "用户不存在")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if user.ID == 0 {
			response.ApiFail(ctx, nil, "密码错误")
			return
		}
	}
	token, err := common.ReleaseToke(user)
	if err != nil {
		response.ApiFail(ctx, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}

	response.ApiSuccess(ctx, gin.H{"token": token}, "登录成功")
	return
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.ApiSuccess(ctx, gin.H{"info": dto.ToUserDto(user.(model.Users))}, "")
}
