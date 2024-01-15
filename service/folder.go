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
	fmt.Println(user, isExist)
	if isExist == false {
		c.JSON(400, gin.H{"msg": "用户不存在"})
		return
	}
	userInfo := model.QueryUserWare(fmt.Sprintf("%s", user))
	wareHouseId := userInfo.WareHouseId
	createFolderRes, folderObj := model.CreateFolder(folder.Name, folder.FolderId, wareHouseId, 0)
	fmt.Println(folderObj)
	if createFolderRes == false {
		c.JSON(200, gin.H{"msg": "创建失败"})
		return
	}
	c.JSON(200, gin.H{"msg": folder})
	return
}
