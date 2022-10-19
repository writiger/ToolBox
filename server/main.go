package main

import (
	"server/middleware"
	"server/route"
	u "server/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	u.Hello()
	c := u.GetConf()
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
