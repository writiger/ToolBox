package api

import "github.com/gin-gonic/gin"

func UserRegister(ctx *gin.Context){
	// 1. 读取POST参数
	// 2. redis中验证验证码是否符合邮箱
	// 3. 生成User
	// 4. 提交服务
	// 5. 响应请求
}