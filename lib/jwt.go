package lib

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	//"github.com/golang-jwt/jwt/v5"
)

type JWToken struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetToken(username string) (string, error) {
	var MySecret = []byte("工人阶级才是领导阶级")
	c := JWToken{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * 3).Unix(),
			Issuer:    "hanxuecheng",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	return token.SignedString(MySecret)
}
