package middleware

import (
	"cloud-storage/lib"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckLogin(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	token := authorization[7:]
	is_vaild, myClaims := lib.CheckJwt(token)
	if is_vaild == false {
		c.JSON(401, gin.H{"msg": "请登录"})
		c.Abort()
	}
	var Claims = myClaims
	user := Claims["iss"]
	fmt.Println(user)
	c.Next()
}
