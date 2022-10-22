package verification

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"server/app/domain"
	"server/utils"
	"time"
)

var jwtKey = []byte(utils.GetConf().JWTKey)

type Claims struct {
	UserAccount string
	jwt.StandardClaims
}

func ReleaseToken(user domain.User) (string, error) {
	fmt.Println("这是密钥", jwtKey)
	expirationTime := time.Now().Add(time.Duration(utils.GetConf().ContinuousTime) * time.Hour)
	claims := &Claims{
		UserAccount: user.Account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "paperIndexSystem",
			Subject:   "User Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
