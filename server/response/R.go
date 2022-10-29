package response

import (
	"encoding/json"
	"server/error_code"
)

// R 统一返回格式
type R struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 返回信息
	Data interface{} `json:"data"` // 响应数据
}

func (r *R) WithMsg(message string) R {
	return R{
		Code: r.Code,
		Msg:  message,
		Data: r.Data,
	}
}

func (r *R) WithData(data interface{}) R {
	return R{
		Code: r.Code,
		Msg:  r.Msg,
		Data: data,
	}
}

func (r *R) ToString() string {
	resp := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: r.Code,
		Msg:  r.Msg,
		Data: r.Data,
	}
	res, _ := json.Marshal(resp)
	return string(res)
}

// 伪构造函数
func response(code int, msg string) *R {
	return &R{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func responseErr(code int) *R {
	return &R{
		Code: code,
		Msg:  errorcode.GetErr(code).Error(),
		Data: nil,
	}
}
