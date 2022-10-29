package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

// MakeKeys 生成基于X509加密的私钥和公钥
func MakeKeys() (X509PrivateKey string, X509PublicKey string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey = base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey))
	//X509对公钥编码
	X509PublicKeyByte, _ := x509.MarshalPKIXPublicKey(&publicKey)
	X509PublicKey = base64.StdEncoding.EncodeToString(X509PublicKeyByte)
	return
}

// Encrypt
// @Param plain string 需要加密的信息
// @Param publicByte []byte 公钥
func Encrypt(plain string, publicK string) ([]byte, error) {
	// Base64解码
	keyGet, _ := base64.StdEncoding.DecodeString(publicK)
	// X509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(keyGet)
	if err != nil {
		return nil, err
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plain))
	if err != nil {
		return nil, err
	}
	return encryptedBytes, nil
}

// Decrypt
// @param
func Decrypt(cipher []byte, privateK string) ([]byte, error) {
	// Base64解码
	keyGet, _ := base64.StdEncoding.DecodeString(privateK)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(keyGet)
	if err != nil {
		return nil, err
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipher)
	return plainText, nil
}
