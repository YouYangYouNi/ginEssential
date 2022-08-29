package main

import (
	"essential/common"
	"essential/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:8005", nil))
	}()
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
