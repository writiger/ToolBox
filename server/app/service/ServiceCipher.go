package service

import (
	"server/app/dao"
	"server/app/domain"
	"server/rsa"
)

func CipherAdd(inputCipher domain.Cipher) error {
	publicKey, err := dao.CipherRedisGet(inputCipher.Owner)
	if err != nil {
		return err
	}
	cipher, err := rsa.Encrypt(inputCipher.Info, publicKey)
	inputCipher.Info = cipher
	err = dao.CipherInsert(inputCipher)
	return err
}
