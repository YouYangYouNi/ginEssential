package middleware

import (
	"essential/common"
	"essential/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware
/**
c.Next() 之前的操作是在 Handler 执行之前就执行；
c.Next() 之后的操作是在 Handler 执行之后再执行；

之前的操作一般用来做验证处理，访问是否允许之类的。
之后的操作一般是用来做总结处理，比如格式化输出、响应结束时间，响应时长计算之类的。
*/
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParesToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//验证通过后获取claims中的userId
		userId := claims.UserId
		db := common.GetDB()
		var user model.Users
		db.First(&user, userId)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
