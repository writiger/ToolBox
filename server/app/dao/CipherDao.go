package dao

import (
	"server/app/domain"
	"server/rsa"
	"strconv"
)

func CipherSet(inputUser domain.User) string {
	privateKey, publicKey := rsa.MakeKeys()
	// 保存私人公钥
	rc.Set("pubKey:"+strconv.FormatInt(inputUser.Id, 10), publicKey, -1)
	return privateKey
}
