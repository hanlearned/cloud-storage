package lib

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWToken struct {
	Username string `json:"username"`
}

var letters = []byte("www.baidu.com")

func GetToken(username string) (string, error) {

	c := jwt.MapClaims{
		"iss": username,                             // 设置签发者
		"sub": "john",                               // 设置主题
		"exp": time.Now().Add(time.Hour * 1).Unix(), // 设置过期时间
	}
	// 固定密钥需要使用：SigningMethodHS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(letters)
}

func JwtParse(jwt_string string) {
	myClaims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(jwt_string, myClaims,
		func(token *jwt.Token) (interface{}, error) {
			return letters, nil
		},
	)

	//token, err = jwt.Parse(jwt_string,
	//	func(token *jwt.Token) (interface{}, error) {
	//		return letters, nil
	//	},
	//)
	if err != nil {
		fmt.Println("解析 JWT 失败:", err)
		return
	}
	var exp = *myClaims
	var exp1 = exp["exp"]
	fmt.Println(exp1)
	fmt.Println(token)
}
