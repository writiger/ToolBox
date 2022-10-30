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
	name := ctx.PostForm("name")
	// 2. 生成Clock
	inputClock := domain.Clock{
		Owner:         ownerAny.(string),
		Kind:          kind,
		Start:         start,
		AchieveTarget: achieveTarget,
		Name:          name,
		Status:        domain.StatusIng,
		Interval:      -1,
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
