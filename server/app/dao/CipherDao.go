package dao

import (
	"server/app/domain"
	"server/rsa"
)

func CipherSet(inputUser domain.User) string {
	privateKey, publicKey := rsa.MakeKeys()
	// 保存私人公钥
	rc.Set("pubKey:"+inputUser.Account, publicKey, -1)
	return privateKey
}
