package main

import (
	"essential/common"
	"essential/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	r := gin.Default()
	common.InitDB()
	//注册路由
	r = router.CollectRoute(r)
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
