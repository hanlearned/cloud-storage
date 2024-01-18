package middleware

import (
	"cloud-storage/lib"
	"github.com/gin-gonic/gin"
)

func CheckLogin(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	token := authorization[7:]
	isValid, myClaims := lib.CheckJwt(token)
	if isValid == false {
		c.JSON(401, gin.H{"msg": "请登录"})
		c.Abort()
	}
	var Claims = myClaims
	userId := Claims["iss"]
	c.Set("user", userId)
	c.Next()
}
