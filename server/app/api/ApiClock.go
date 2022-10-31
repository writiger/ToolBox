package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/app/domain"
	"server/app/service"
	"server/response"
	"strconv"
)

func ClockUpload(ctx *gin.Context) {
	// 1. 获取参数
	ownerAny, _ := ctx.Get("userAccount")
	kind, _ := strconv.Atoi(ctx.PostForm("kind"))
	start, _ := strconv.Atoi(ctx.PostForm("start"))
	achieveTarget, _ := strconv.Atoi(ctx.PostForm("achieve_target"))
	interval := ctx.PostForm("interval")
	name := ctx.PostForm("name")
	// 2. 生成Clock
	inputClock := domain.Clock{
		Owner:         ownerAny.(string),
		Kind:          kind,
		Start:         start,
		AchieveTarget: achieveTarget,
		Name:          name,
		Status:        domain.StatusIng,
		Interval:      interval,
	}
	// 3. 调用服务
	err := service.ClockUpload(inputClock)
	// 4. 响应请求
	if err != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK)
}

func ClockQueryByOwner(ctx *gin.Context) {
	// 1. 获取参数
	ownerAny, _ := ctx.Get("userAccount")
	// 2. 调用服务
	clocks, err := service.ClockQueryByOwner(ownerAny.(string))
	// 3. 响应请求
	if err != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK.WithData(clocks))
}

func ClockIn(ctx *gin.Context) {
	// 1. 获取参数
	ownerAny, _ := ctx.Get("userAccount")
	clockIdString := ctx.Param("clockId")
	clock, _ := strconv.ParseInt(clockIdString, 10, 64)
	// 2. 调用服务
	err := service.ClockClockIn(clock, ownerAny.(string))
	// 3. 响应请求
	if err != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK)
}

func ClockDelete(ctx *gin.Context) {
	// 1. 获取参数
	ownerAny, _ := ctx.Get("userAccount")
	clockIdString := ctx.Param("clockId")
	clock, _ := strconv.ParseInt(clockIdString, 10, 64)
	// 2. 调用服务
	err := service.ClockClockDelete(clock, ownerAny.(string))
	// 3. 响应请求
	if err != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK)
}
