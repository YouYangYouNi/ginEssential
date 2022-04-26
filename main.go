package main

import (
	"essential/common"
	"essential/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//注册路由
	r = router.CollectRoute(r)
	common.InitDB()
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}
