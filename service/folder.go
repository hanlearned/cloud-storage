package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
import "cloud-storage/service/schema"

func CreateFolder(c *gin.Context) {
	var folder schema.Folder
	err := c.ShouldBindJSON(&folder)
	if err != nil {
		c.JSON(200, gin.H{"msg": fmt.Sprintf("%s", err)})
	}
	
}
