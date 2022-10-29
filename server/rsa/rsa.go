package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

// MakeKeys 生成基于X509加密Base64编码的私钥和公钥的字符串
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

// Encrypt 加密
// @Param plain string 需要加密的信息
// @Param publicByte string 公钥
// @return encryptedPlain string 加密后的密文
func Encrypt(plain string, publicK string) (string, error) {
	// Base64解码
	keyGet, _ := base64.StdEncoding.DecodeString(publicK)
	// X509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(keyGet)
	if err != nil {
		return "", err
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plain))
	if err != nil {
		return "", err
	}
	encryptedPlain := base64.StdEncoding.EncodeToString(encryptedBytes)
	return encryptedPlain, nil
}

// Decrypt 解密
// @param cipher string 加密后的密文
// @param privateK string base64编码的私钥
// @return plainText string 解密后的明文
func Decrypt(cipher string, privateK string) (string, error) {
	cipherByte, _ := base64.StdEncoding.DecodeString(cipher)
	// Base64解码
	keyGet, _ := base64.StdEncoding.DecodeString(privateK)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(keyGet)
	if err != nil {
		return "", err
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherByte)
	return string(plainText), nil
}
