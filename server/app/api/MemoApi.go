package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/app/service"
	"server/response"
)

func MemoQueryByOwner(ctx *gin.Context) {
	// 1. 读取用户id
	ownerAny, _ := ctx.Get("userId")
	owner := fmt.Sprintf("%v", ownerAny)
	// 2. 调用service
	res, err := service.MemoQueryByOwnerService(owner)
	// 3. 响应请求
	if err != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK.WithData(res))
}

func MemoUpload(ctx *gin.Context) {
	fmt.Println(ctx.Get("userId"))
}
