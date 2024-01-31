package service

import (
	"cloud-storage/lib"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	lib.CreateFolder("1")
	c.JSON(200, gin.H{})
}
