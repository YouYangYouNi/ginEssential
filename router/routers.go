package router

import (
	"essential/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/user/register", controller.Register)
	r.POST("/api/user/login", controller.Login)

	return r
}
