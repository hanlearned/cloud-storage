package controller

import (
	"cloud-storage/lib"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	token := authorization[7:]
	fmt.Println(token)
	lib.CheckJwt(token)
	c.JSON(200, gin.H{})
}
