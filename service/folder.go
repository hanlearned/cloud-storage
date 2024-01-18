package service

import (
	"cloud-storage/model"
	"cloud-storage/service/schema"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ListFolder(c *gin.Context) {
	wareHouseId, isExist := c.Get("wareHouseId")
	if isExist == false {
		c.JSON(400, gin.H{"msg": "仓库不存在"})
		return
	}
	listFolder := model.ListFolder(wareHouseId.(int))
	c.JSON(200, gin.H{"msg": listFolder})
}

func CreateFolder(c *gin.Context) {
	var folder schema.Folder
	err := c.ShouldBindJSON(&folder)
	if err != nil {
		c.JSON(200, gin.H{"msg": fmt.Sprintf("%s", err)})
		return
	}
	wareHouseId, isExist := c.Get("wareHouseId")
	if isExist == false {
		c.JSON(400, gin.H{"msg": "仓库不存在"})
		return
	}
	_, err = model.CreateFolder(folder.Name, folder.FolderId, wareHouseId.(int))
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
	wareHouseId, isExist := c.Get("wareHouseId")
	if isExist == false {
		c.JSON(400, gin.H{"msg": "仓库不存在"})
		return
	}
	res := model.DeleteFolder(folder.FolderId, wareHouseId.(int))
	if res == true {
		c.JSON(200, gin.H{"msg": "删除成功"})
		return
	}
	c.JSON(200, gin.H{"msg": "删除失败"})
	return
}
