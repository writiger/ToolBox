package route

import "github.com/gin-gonic/gin"

func CollectRoute(r *gin.Engine) {
	// 按照实体类注册路由 便于管理
	collectUserRoute(r)
}

func collectUserRoute(r *gin.Engine) {
	userApiURI := "/api/user/"

	// 获取用户列表
	r.GET(userApiURI)
}
