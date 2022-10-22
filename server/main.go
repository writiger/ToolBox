package main

import (
	"server/middleware"
	"server/router"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.Hello()
	c := utils.GetConf()
	r := gin.Default()

	// 使用中间件配置跨域
	r.Use(middleware.Cors())

	// 注册路由
	route.CollectRoute(r)

	// 开始监听
	err := r.Run(c.IP + ":" + c.Port)
	if err != nil {
		panic(err)
	}

}
