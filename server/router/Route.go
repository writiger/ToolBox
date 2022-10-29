package route

import (
	"server/app/api"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) {
	// 按照实体类注册路由 便于管理
	collectUserRoute(r)
	collectMemoRoute(r)
	collectCipherRoute(r)
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
	r.GET(memoApiURI, middleware.AuthMiddleWare(), api.MemoQueryByOwner)
	// 新增memo
	r.POST(memoApiURI, middleware.AuthMiddleWare(), api.MemoUpload)
	// 删除memo
	r.DELETE(memoApiURI+"/:memoId", middleware.AuthMiddleWare(), api.MemoDelete)
	// 更新memo
	r.PUT(memoApiURI, middleware.AuthMiddleWare(), api.MemoUpdate)
}

func collectCipherRoute(r *gin.Engine) {
	cipherApiURI := "/api/cipher"

	// 新增cipher
	r.POST(cipherApiURI, middleware.AuthMiddleWare(), api.CipherUpload)
	// 根据owner查询
	r.GET(cipherApiURI, middleware.AuthMiddleWare(), api.CipherQueryByOwner)
	// 翻译密码
	r.POST(cipherApiURI+"/translation", middleware.AuthMiddleWare(), api.CipherTranslate)
}
