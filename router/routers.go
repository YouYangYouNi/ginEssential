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

	categoryRouters := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRouters.POST("", categoryController.Create)
	categoryRouters.PUT("/:id", categoryController.Update)
	categoryRouters.GET("/:id", categoryController.Show)
	categoryRouters.DELETE("/:id", categoryController.Delete)
	return r
}
