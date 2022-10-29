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

func CipherQueryByOwnerService(owner string) ([]domain.Cipher, error) {
	ciphers, err := dao.CipherFindByOwner(owner)
	return ciphers, err
}

func CipherTranslate(cipherId int64, privateKey string) (string, error) {
	// 查询密码
	cipher, err := dao.CipherFindById(cipherId)
	if err != nil {
		return "", err
	}
	// 翻译密码
	plain, err := rsa.Decrypt(cipher.Info, privateKey)
	return plain, err
}
