package lib

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWToken struct {
	Username string `json:"username"`
}

var letters = []byte("www.baidu.com")

func GetToken(username_id int) (string, error) {
	// 生成 token
	c := jwt.MapClaims{
		"iss": username_id,                                           // 设置签发者
		"sub": "john",                                                // 设置主题
		"exp": time.Now().Add(time.Second * 60 * 60 * 24 * 7).Unix(), // 设置过期时间 时间校验为false
	}
	// 固定密钥需要使用：SigningMethodHS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(letters)
}

func CheckJwt(jwtString string) (bool, map[string]interface{}) {
	// 解析 token, 判断token是否超时
	myClaims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(jwtString, myClaims,
		func(token *jwt.Token) (interface{}, error) {
			return letters, nil
		},
	)
	if err != nil {
		return false, *myClaims
	}
	if token.Valid {
		return true, *myClaims
	}
	return false, *myClaims
}
