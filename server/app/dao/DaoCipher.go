package dao

import (
	"server/app/domain"
	"server/rsa"
)

func CipherRedisSet(inputUser domain.User) string {
	privateKey, publicKey := rsa.MakeKeys()
	// 保存私人公钥
	rc.Set("pubKey:"+inputUser.Account, publicKey, -1)
	// 返回私人私钥
	return privateKey
}

func CipherRedisGet(account string) (string, error) {
	result, err := rc.Get("pubKey:" + account).Result()
	return result, err
}

func CipherInsert(inputCipher domain.Cipher) error {
	_, err := engine.Insert(&inputCipher)
	return err
}

func CipherFindByOwner(owner string) ([]domain.Cipher, error) {
	ciphers := make([]domain.Cipher, 0)
	err := engine.Where("owner = ?", owner).Find(&ciphers)
	return ciphers, err
}
