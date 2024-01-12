package service

import (
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "上传文件接口"})
}
