package api

import (
	"net/http"
	"server/app/service"
	"server/response"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	// 1. 读取POST参数
	// 2. redis中验证验证码是否符合邮箱
	// 3. 生成User
	// 4. 提交服务
	// 5. 响应请求
}

func UserEmailVerify(ctx *gin.Context) {
	// 1. 读取POST参数
	account := ctx.PostForm("account")
	// 2. 验证邮箱是否存在
	res, err := service.UserEmailVerifyService(account)
	// 3. 判断结果
	switch res {
	case 30105:
		ctx.JSON(http.StatusOK, response.ErrUserAlreadyExist)
	case 200:
		ctx.JSON(http.StatusOK, nil)
	case 500:
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err.Error()))
	}
}
