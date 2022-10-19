package middleware

import (
	"net/http"
	"server/response"
	conf "server/utils"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	c := conf.GetConf()
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", c.DomainName) // 设置允许访问所有域
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Set("content-type", "application/json")
		// 前端向后端发送的OPTIONS测试请求 服务器只需回应200
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, response.OK)
		}

		//处理请求
		context.Next()
	}
}
