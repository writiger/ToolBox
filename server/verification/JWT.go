package verification

import (
	"github.com/dgrijalva/jwt-go"
	"server/app/domain"
	"server/utils"
	"time"
)

var jwtKey = []byte(utils.GetConf().JWTKey)

type Claims struct {
	UserAccount string
	jwt.StandardClaims
	Id int64
}

func ReleaseToken(user domain.User) (string, error) {
	expirationTime := time.Now().Add(time.Duration(utils.GetConf().ContinuousTime) * time.Hour)
	claims := &Claims{
		UserAccount: user.Account,
		Id:          user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "ToolBoxW",
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
