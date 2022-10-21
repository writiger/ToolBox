package verification

import (
	"crypto/rand"
	"math/big"
)

func GenerateRegisterCode() int {
	// 生成6位数验证码
	res, _ := rand.Int(rand.Reader, big.NewInt(900000))
	return int(res.Int64() + 100000)
}
