package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/app/domain"
	"server/app/service"
	"server/response"
	"strconv"
	"time"
)

// ConsumptionCost 新增消费记录
func ConsumptionCost(ctx *gin.Context) {
	// 1. 读取用户id
	ownerAny, _ := ctx.Get("userId")
	// 断言类型
	owner := ownerAny.(int64)
	// 2. 读取POST参数
	infoStr := ctx.PostForm("info")
	amountStr := ctx.PostForm("amount")
	isFoodStr := ctx.PostForm("isFood")
	// 3. 验证参数
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, response.Err.WithMsg(err.Error()))
		return
	}
	isFood, err := strconv.ParseBool(isFoodStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, response.Err.WithMsg(err.Error()))
		return
	}
	inputConsumption := domain.Consumption{
		Info:   infoStr,
		Owner:  owner,
		Amount: amount,
		IsFood: isFood,
	}
	// 4. 调用服务
	err = service.ConsumptionUpload(&inputConsumption)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, response.Err.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK)
}

// ConsumptionGetMonth 查询用户一个月的消费记录
func ConsumptionGetMonth(ctx *gin.Context) {
	// 1. 读取用户id
	ownerAny, _ := ctx.Get("userId")
	// 断言类型
	owner := ownerAny.(int64)
	// 2. 获取路径中的时间戳
	startStr := ctx.Param("start")
	startInt, err := strconv.ParseInt(startStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, response.Err.WithMsg(err.Error()))
		return
	}
	start := time.Unix(startInt, 0)
	// 3. 调用服务
	res, err := service.ConsumptionGetMonth(owner, start)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, response.Err.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK.WithData(res))
}

// ConsumptionEngel 查询用户的恩格尔系数
func ConsumptionEngel(ctx *gin.Context) {
	// 1. 获取路径中的用户ID
	idStr := ctx.Param("userId")
	userId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, response.Err.WithMsg(err.Error()))
		return
	}
	// 2. 调用服务
	engel, err := service.ConsumptionEngel(userId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, response.Err.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK.WithData(engel))
}
