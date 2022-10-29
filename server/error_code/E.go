package errorcode

import (
	"errors"
)

// 错误码规则:
// (1) 错误码需为 > 0 的数;
//
// (2) 错误码为 5 位数:
//              ----------------------------------------------------------
//                  第1位               2、3位                  4、5位
//              ----------------------------------------------------------
//                  A                    B                      C
//              ----------------------------------------------------------
//                服务级错误码           函数位置    	         具体错误
//              ----------------------------------------------------------
//
// A表示模块报错位置，1代表接口层（api，1-0000），2代表service层（2-0000），3表示model层（3-0000），4表示定时任务错误（4-0000），5表示初始化错误（5-0000）
// B模块，两位数字代表函数位置（接口函数，service函数，model函数），01～99，例如10101，表示第一个接口函数的错误
// C模块，表示函数内错误位置/内容，普通报错01~98，panic错误99，自己定义一些常用的错误码

func GetErr(code int) error {
	return errors.New("错误信息:" + ErrMap[code])
}

var ErrMap map[int]string

const (
	ErrParam                = 10001
	ErrUserService          = 30100
	ErrUserWrongPassword    = 30101
	ErrUserNil              = 30102
	ErrUserAlreadyExist     = 30103
	ErrUserEmailFail        = 30104
	ErrUserPermissionDenied = 30105
	ErrCipherNil            = 30301
	ErrRsaDecrypt           = 30401
)

func init() {
	ErrMap = make(map[int]string)
	ErrMap[10001] = "参数有误"
	ErrMap[30100] = "用户服务错误"
	ErrMap[30101] = "用户密码错误"
	ErrMap[30102] = "用户不存在"
	ErrMap[30103] = "用户已存在"
	ErrMap[30104] = "验证码失效或错误"
	ErrMap[30301] = "密码不存在"
	ErrMap[30401] = "私钥错误"
}
