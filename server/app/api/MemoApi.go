package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/app/domain"
	"server/app/service"
	"server/response"
	"strconv"
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
	// 1. 读取参数
	ownerAny, _ := ctx.Get("userId")
	ownerString := fmt.Sprintf("%v", ownerAny)
	owner, _ := strconv.ParseInt(ownerString, 10, 64)
	info := ctx.PostForm("info")
	status, err := strconv.Atoi(ctx.PostForm("status"))
	if err != nil {
		ctx.JSON(http.StatusOK, response.ErrParam)
	}
	endTime, err2 := strconv.ParseInt(ctx.PostForm("end_time"), 10, 64)
	if err2 != nil {
		ctx.JSON(http.StatusOK, response.ErrParam)
	}
	// 2. 生成Memo
	memoInput := domain.Memo{
		Owner:   owner,
		Info:    info,
		Status:  status,
		EndTime: endTime,
	}
	// 3. 调用服务
	err3 := service.MemoAdd(memoInput)
	// 4. 响应请求
	if err3 != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err3.Error()))
	}
	ctx.JSON(http.StatusOK, response.OK)
}
