package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

// MakeKeys 生成基于X509加密的私钥和公钥
func MakeKeys() (X509PrivateKey []byte, X509PublicKey []byte) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey = x509.MarshalPKCS1PrivateKey(privateKey)
	//X509对公钥编码
	X509PublicKey, err = x509.MarshalPKIXPublicKey(&publicKey)
	return
}

// Encrypt
// @Param plain string 需要加密的信息
// @Param publicByte []byte 公钥
func Encrypt(plain string, publicByte []byte) ([]byte, error) {
	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicByte)
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
func Decrypt(cipher []byte, privateByte []byte) ([]byte, error) {
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(privateByte)
	if err != nil {
		return nil, err
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipher)
	return plainText, nil
}
