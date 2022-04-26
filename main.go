package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

type Users struct {
	gorm.Model
	Name      string
	Telephone string
	Password  string
}

func main() {
	r := gin.Default()
	db := InitDB()
	r.POST("/api/user/register", func(ctx *gin.Context) {
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
		if ExistTelephone(db, telephone) {
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
			name = RandomString(10)
		}

		//创建用户
		newUser := Users{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		result := db.Create(&newUser)
		if result.Error != nil {
			ctx.JSON(200, gin.H{
				"msg": "注册失败：" + name,
			})
			return
		}

		ctx.JSON(200, gin.H{
			"msg": "注册成功：" + name,
		})
	})
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}

func RandomString(n int) string {
	letter := []byte("abcdefghijklmnopqrstuvwxyzABCDEFJHJKLLJOOUOYOY")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letter[rand.Intn(len(letter))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/essential?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	db.AutoMigrate(&Users{})

	return db
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
