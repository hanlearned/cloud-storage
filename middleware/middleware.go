package middleware

import (
	"cloud-storage/lib"
	"cloud-storage/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckLogin(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if len(authorization) == 0 {
		c.JSON(401, gin.H{"msg": "token 不存在"})
		c.Abort()
		return
	}
	token := authorization[7:]
	isValid, myClaims := lib.CheckJwt(token)
	if isValid == false {
		c.JSON(401, gin.H{"msg": "请登录"})
		c.Abort()
		return
	}
	var Claims = myClaims
	userId := Claims["iss"]
	userInfo, err := model.QueryUserWare(userId)
	if err != nil {
		c.JSON(200, gin.H{"msg": fmt.Sprintf("%s", err)})
		c.Abort()
		return
	}
	wareHouseId := userInfo.WareHouseId
	c.Set("user", userId)
	c.Set("wareHouseId", wareHouseId)
	c.Next()
}
