package router

import (
	"essential/controller"
	"essential/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/user/register", controller.Register)
	r.POST("/api/user/login", controller.Login)
	r.POST("/api/user/info", middleware.AuthMiddleware(), controller.Info)

	return r
}
