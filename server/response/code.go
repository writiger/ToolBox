package response

import "server/error_code"

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
// A表示模块报错位置，1代表接口层（controller，1-0000），2代表service层（2-0000），3表示model层（3-0000），4表示定时任务错误（4-0000），5表示初始化错误（5-0000）
// B模块，两位数字代表函数位置（接口函数，service函数，model函数），01～99，例如10101，表示第一个接口函数的错误
// C模块，表示函数内错误位置/内容，普通报错01~98，panic错误99，自己定义一些常用的错误码
var (
	OK  = response(200, "ok")    // 通用成功
	Err = response(500, "wrong") // 通用错误

	// 服务级错误码

	ErrParam = responseErr(errorcode.ErrParam)

	//用户模块错误

	ErrUserService          = responseErr(errorcode.ErrUserService)
	ErrUserWrongPassword    = responseErr(errorcode.ErrUserWrongPassword)
	ErrUserNil              = responseErr(errorcode.ErrUserNil)
	ErrUserAlreadyExist     = responseErr(errorcode.ErrUserAlreadyExist)
	ErrUserEmailFail        = responseErr(errorcode.ErrUserEmailFail)
	ErrUserPermissionDenied = responseErr(errorcode.ErrUserPermissionDenied)
	ErrCipherNil            = responseErr(errorcode.ErrCipherNil)
	ErrRsaDecrypt           = responseErr(errorcode.ErrRsaDecrypt)
)
