package api

import (
	"net/http"
	"server/app/domain"
	"server/app/service"
	"server/error_code"
	"server/response"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	// 1. 读取POST参数
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")
	code := ctx.PostForm("code")
	// 2. redis中验证验证码是否符合邮箱
	conform := service.UserEmailVerifyCodeService(account, code)
	if !conform {
		ctx.JSON(http.StatusOK, response.ErrUserEmailFail)
		return
	}
	// 3. 生成User
	newUser := domain.User{
		Name:     "临时用户名",
		Avatar:   1,
		Account:  account,
		Password: utils.MD5(password),
	}
	// 4. 提交服务
	privateKey, err := service.UserAddService(newUser)
	// 5. 响应请求
	if err != nil {
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, response.OK.WithData(privateKey))
	}
}

func UserEmailVerify(ctx *gin.Context) {
	// 1. 读取POST参数
	account := ctx.PostForm("account")
	// 2. 验证邮箱是否存在
	err := service.UserEmailVerifyService(account)
	//// 3. 判断结果
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, response.OK)
	case errorcode.GetErr(errorcode.ErrUserAlreadyExist):
		ctx.JSON(http.StatusOK, response.ErrUserAlreadyExist)
	default:
		ctx.JSON(http.StatusOK, response.Err.WithMsg(err.Error()))
	}
}

func UserLogin(ctx *gin.Context) {
	// 1. 读取POST参数
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")
	// 2. 提交登录服务
	token, err := service.UserLoginService(domain.User{
		Account:  account,
		Password: utils.MD5(password),
	})
	// 3. 登陆成功 返回session
	if token != "" {
		ctx.JSON(http.StatusOK, response.OK.WithData(gin.H{
			"token": token,
		}))
		return
	}
	// 4. 登陆失败 返回失败原因
	ctx.JSON(http.StatusOK, response.Err.WithMsg(err.Error()))
}
