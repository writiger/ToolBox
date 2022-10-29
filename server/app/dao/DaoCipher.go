package dao

import (
	"server/app/domain"
	"server/error_code"
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

func CipherFindById(id int64) (domain.Cipher, error) {
	cipher := new(domain.Cipher)
	has, err := engine.ID(id).Get(cipher)
	if !has {
		return domain.Cipher{}, errorcode.GetErr(errorcode.ErrCipherNil)
	}
	return *cipher, err
}

func CipherFindByOwner(owner string) ([]domain.Cipher, error) {
	ciphers := make([]domain.Cipher, 0)
	err := engine.Where("owner = ?", owner).Find(&ciphers)
	for _, item := range ciphers {
		// 去除密码  暂时没找到更好的方法
		item.Info = ""
	}
	return ciphers, err
}
