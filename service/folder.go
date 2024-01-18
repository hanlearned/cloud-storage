package service

import (
	"cloud-storage/model"
	"cloud-storage/service/schema"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateFolder(c *gin.Context) {
	var folder schema.Folder
	err := c.ShouldBindJSON(&folder)
	if err != nil {
		c.JSON(200, gin.H{"msg": fmt.Sprintf("%s", err)})
	}
	user, isExist := c.Get("user")
	if isExist == false {
		c.JSON(400, gin.H{"msg": "用户不存在"})
		return
	}
	userInfo, err := model.QueryUserWare(fmt.Sprintf("%s", user))
	if err != nil {
		c.JSON(200, gin.H{"msg": fmt.Sprintf("%s", err)})
		return
	}
	wareHouseId := userInfo.WareHouseId
	createFolderRes, _ := model.CreateFolder(folder.Name, folder.FolderId, wareHouseId, 0)
	if createFolderRes == false {
		c.JSON(200, gin.H{"msg": "创建失败"})
		return
	}
	c.JSON(200, gin.H{"msg": folder})
	return
}
