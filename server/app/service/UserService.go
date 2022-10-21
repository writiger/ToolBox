package service

import (
	"server/app/dao"
	"server/verification"
)

func UserEmailVerifyService(account string) (int, error) {
	// 1. 邮箱已存在 返回错误码
	has, err1 := dao.UserIsExist(account)
	if err1 != nil {
		return 500, err1
	}
	if has {
		return 30105, nil
	}
	// 2. 邮箱不存在 生成验证码 保存redis
	code := verification.GenerateRegisterCode()
	err2 := dao.UserEmailRedisSave(account, code)
	if err2 != nil {
		return 500, err2
	}
	// 3. 发送邮件 并返回

	return 200, nil
}
