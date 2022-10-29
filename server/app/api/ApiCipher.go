package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/app/domain"
	"server/app/service"
	"server/response"
)

func CipherUpload(ctx *gin.Context) {
	// 1. 读取用户id
	ownerAny, _ := ctx.Get("userAccount")
	// 断言类型
	owner := ownerAny.(string)
	// 2. 读取参数
	use := ctx.PostForm("use")
	info := ctx.PostForm("info")
	// 3. 生成cipher类
	inputCipher := domain.Cipher{
		Owner: owner,
		Info:  info,
		Use:   use,
	}
	// 4. 调用服务
	err := service.CipherAdd(inputCipher)
	// 5. 响应请求
	if err != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK)
}
