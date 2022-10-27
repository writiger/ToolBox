package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

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

func Encrypt(plain string, publicByte []byte) ([]byte, error) {
	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicByte)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plain))
	if err != nil {
		return nil, err
	}
	return encryptedBytes, nil
}

func Decrypt(cipher []byte, privateByte []byte) []byte {
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(privateByte)
	if err != nil {
		panic(err)
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipher)
	return plainText
}
