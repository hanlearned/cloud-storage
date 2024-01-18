package service

import (
	"cloud-storage/model"
	"cloud-storage/service/schema"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ListFolder(c *gin.Context) {
	listFolder := model.ListFolder(1)
	c.JSON(200, gin.H{"msg": listFolder})
}

func CreateFolder(c *gin.Context) {
	var folder schema.Folder
	err := c.ShouldBindJSON(&folder)
	if err != nil {
		c.JSON(200, gin.H{"msg": fmt.Sprintf("%s", err)})
		return
	}
	user, isExist := c.Get("user")
	if isExist == false {
		c.JSON(400, gin.H{"msg": "用户不存在"})
		return
	}
	userInfo, err := model.QueryUserWare(user)
	if err != nil {
		c.JSON(200, gin.H{"msg": fmt.Sprintf("%s", err)})
		return
	}
	wareHouseId := userInfo.WareHouseId
	_, err = model.CreateFolder(folder.Name, folder.FolderId, wareHouseId)
	if err != nil {
		c.JSON(200, gin.H{"msg": fmt.Sprintf("%s", err)})
		return
	}
	c.JSON(200, gin.H{"msg": folder})
	return
}

func DeleteFolder(c *gin.Context) {
	var folder schema.DeleteFolder
	err := c.ShouldBindUri(&folder)
	if err != nil {
		c.JSON(200, gin.H{"folderId": fmt.Sprintf("%s", folder)})
		return
	}
	res := model.DeleteFolder(folder.FolderId)
	if res == true {
		c.JSON(200, gin.H{"msg": "删除成功"})
		return
	}
	c.JSON(200, gin.H{"msg": "删除失败"})
	return
}
