package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/app/dao"
	"server/app/domain"
	"server/response"
	"server/verification"
	"strings"
)

// AuthMiddleWare 验证并将userId写入ctx
func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		// 验证token格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, response.ErrUserPermissionDenied)
			ctx.Abort()
			return
		}
		// 删除token中的无效部分
		tokenString = tokenString[7:]

		token, claims, err := verification.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, response.ErrUserPermissionDenied)
			ctx.Abort()
			return
		}

		has, err := dao.UserIsExist(domain.User{
			Account: claims.UserAccount,
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, response.ErrUserPermissionDenied)
			ctx.Abort()
			return
		}

		if !has {

			ctx.JSON(http.StatusUnauthorized, response.ErrUserNil)
			ctx.Abort()
			return
		}

		// 验证通过  写入User
		ctx.Set("userId", claims.Id)
		ctx.Next()
	}
}
