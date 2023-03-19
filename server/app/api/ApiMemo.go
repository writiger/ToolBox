package api

import (
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
	// 断言类型
	owner := ownerAny.(int64)
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
	info := ctx.PostForm("info")
	// 断言类型
	owner := ownerAny.(int64)
	endTime, err := strconv.ParseInt(ctx.PostForm("end_time"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ErrParam)
		return
	}
	// 2. 生成Memo
	memoInput := domain.Memo{
		Owner:   owner,
		Info:    info,
		Status:  0,
		EndTime: endTime,
	}
	// 3. 调用服务
	err3 := service.MemoAdd(memoInput)
	// 4. 响应请求
	if err3 != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err3.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK)
}

func MemoDelete(ctx *gin.Context) {
	// 1. 读取参数
	ownerAny, _ := ctx.Get("userId")
	// 断言类型
	ownerString := strconv.FormatInt(ownerAny.(int64), 10)
	owner, _ := strconv.ParseInt(ownerString, 10, 64)
	memoIdString := ctx.Param("memoId")
	memoId, err := strconv.ParseInt(memoIdString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ErrParam)
		return
	}
	// 2. 生成Memo
	memoInput := domain.Memo{
		Owner: owner,
		Id:    memoId,
	}
	// 3. 调用服务
	err2 := service.MemoDelete(memoInput)
	// 4. 返回请求
	if err2 != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err2.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK)
}

func MemoUpdate(ctx *gin.Context) {
	// 1. 读取参数
	ownerAny, _ := ctx.Get("userId")
	// 断言类型
	owner, _ := ownerAny.(int64)
	// 2. owner绑定
	inputMemo := domain.Memo{
		Owner: owner,
	}
	// 3. 补充参数
	err := ctx.Bind(&inputMemo)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ErrParam)
		return
	}
	// 4. 响应请求
	err2 := service.MemoPut(inputMemo)
	if err2 != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err2.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK)
}
