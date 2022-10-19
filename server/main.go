package main

import (
	"fmt"
	conf "server/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	c := conf.GetConf()

	router := gin.Default()

	err := router.Run(c.IP + ":" + c.Port)
	if err != nil {
		fmt.Println("服务器启动失败")
	}
	fmt.Println("服务端启动")
}
