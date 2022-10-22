package main

import (
	"server/middleware"
	route "server/router"
	utils "server/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.Hello()
	c := utils.GetConf()
	router := gin.Default()

	// 使用中间件配置跨域
	router.Use(middleware.Cors())

	// 注册路由
	route.CollectRoute(router)

	// 开始监听
	err := router.Run(c.IP + ":" + c.Port)
	if err != nil {
		panic(err)
	}

}
