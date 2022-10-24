package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/app/service"
)

func MemoQueryByOwner(ctx *gin.Context) {
	// 1. 读取用户id
	owner := ctx.Param("owner")
	// 2. 调用service
	res, err := service.MemoQueryByOwnerService(owner)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
	// 3. 响应请求
}
