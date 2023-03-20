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
	collectClockRoute(r)
	collectConsumption(r)
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
	// 修改余额信息
	r.PUT(userApiURI+"/pay", middleware.AuthMiddleWare(), api.UserChangePay)
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

func collectClockRoute(r *gin.Engine) {
	clockApiURI := "/api/clock"

	// 新增clock
	r.POST(clockApiURI, middleware.AuthMiddleWare(), api.ClockUpload)
	// 查询clock
	r.GET(clockApiURI, middleware.AuthMiddleWare(), api.ClockQueryByOwner)
	// 打卡
	r.PUT(clockApiURI+"/:clockId", middleware.AuthMiddleWare(), api.ClockIn)
	// 删除打卡
	r.DELETE(clockApiURI+"/:clockId", middleware.AuthMiddleWare(), api.ClockDelete)
}

func collectConsumption(r *gin.Engine) {
	consumptionApiURI := "/api/consumption"

	// 新增消费记录
	r.POST(consumptionApiURI, middleware.AuthMiddleWare(), api.ConsumptionCost)
	// 查询一个月的消费记录
	r.GET(consumptionApiURI+"/month/:start", middleware.AuthMiddleWare(), api.ConsumptionGetMonth)
	// 查询用户的恩格尔系数
	r.GET(consumptionApiURI+"/engel/:userId", api.ConsumptionEngel)
}
