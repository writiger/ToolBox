package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/app/domain"
	"server/app/service"
	"server/error_code"
	"server/response"
	"strconv"
)

func CipherUpload(ctx *gin.Context) {
	// 1. 读取用户account
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
	if err == errorcode.GetErr(errorcode.ErrRsaDecrypt) {
		ctx.JSON(http.StatusOK, response.ErrRsaDecrypt)
		return
	} else if err != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK)
}

func CipherQueryByOwner(ctx *gin.Context) {
	// 1. 读取用户account
	ownerAny, _ := ctx.Get("userAccount")
	// 断言类型
	owner := ownerAny.(string)
	// 2. 调用服务
	ciphers, err := service.CipherQueryByOwnerService(owner)
	// 3. 响应请求
	if err != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err.Error()))
	}
	ctx.JSON(http.StatusOK, response.OK.WithData(ciphers))
}

func CipherTranslate(ctx *gin.Context) {
	// 1. 读取参数
	cipherIdString := ctx.PostForm("cipher_id")
	privateKey := ctx.PostForm("private_key")
	cipherId, err := strconv.ParseInt(cipherIdString, 10, 64)
	// 2. 调用服务

	plain, err := service.CipherTranslate(cipherId, privateKey)
	// 3. 相应请求
	if err != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.OK.WithData(plain))

}
