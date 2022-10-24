package route

import (
	"server/app/api"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) {
	// 按照实体类注册路由 便于管理
	collectUserRoute(r)
	collectMemoRoute(r)
}

func collectUserRoute(r *gin.Engine) {
	userApiURI := "/api/user"

	// 获取用户列表
	r.GET(userApiURI)
	// 请求发送验证码
	r.POST(userApiURI+"/mail", api.UserEmailVerify)
	// 注册用户
	r.POST(userApiURI, api.UserRegister)
	// 登录
	r.POST(userApiURI+"/login", api.UserLogin)

}

func collectMemoRoute(r *gin.Engine) {
	memoApiURI := "/api/memo"

	// 根据owner查询memo
	r.GET(memoApiURI+"/:owner", api.MemoQueryByOwner)
}
