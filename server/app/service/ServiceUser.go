package service

import (
	"server/app/dao"
	"server/app/domain"
	"server/error_code"
	"server/verification"
)

func UserEmailVerifyService(account string) error {
	inputUser := domain.User{
		Account: account,
	}
	// 1. 邮箱已存在 返回错误码
	has, err1 := dao.UserIsExist(inputUser)
	if err1 != nil {
		return err1
	}
	if has {
		return errorcode.GetErr(errorcode.ErrUserAlreadyExist)
	}
	// 2. 邮箱不存在 生成验证码 保存redis
	code := verification.GenerateRegisterCode()
	err2 := dao.UserEmailRedisSet(inputUser, code)
	if err2 != nil {
		return err2
	}
	// 3. 发送邮件 并返回
	err3 := verification.SendMail(account, code)
	return err3
}

func UserEmailVerifyCodeService(account string, code string) bool {
	codeExist := dao.UserEmailCodeGet(domain.User{
		Account: account,
	})
	return code == codeExist
}

func UserAddService(inputUser domain.User) (string, error) {
	err := dao.UserInsert(inputUser)
	if err != nil {
		return "", err
	}
	// 获取私人密钥
	privateKey := dao.CipherRedisSet(inputUser)
	return privateKey, err
}

// UserLoginService
// @return string "由查询到的用户生成token 返回空则登陆失败"
func UserLoginService(inputUser domain.User) (string, error) {
	has, err := dao.UserIsExist(inputUser)
	if err != nil {
		return "", err
	}
	if !has {
		// 用户不存在
		return "", errorcode.GetErr(errorcode.ErrUserNil)
	} else {
		// 用户存在
		login, err := dao.UserLogin(&inputUser)
		if err != nil {
			return "", err
		}
		if login.Id != 0 {
			// 登陆成功
			login.Password = ""
			token, err2 := verification.ReleaseToken(login)
			if err2 != nil {
				return "", err2
			}
			return token, nil
		} else {
			return "", errorcode.GetErr(errorcode.ErrUserWrongPassword)
		}
	}
}