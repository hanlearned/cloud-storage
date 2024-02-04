package service

import (
	"cloud-storage/conf"
	"cloud-storage/lib"
	"cloud-storage/model"
	"cloud-storage/service/schema"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"os"
)

var store = conf.StoreConfig

func UploadFile(c *gin.Context) {
	var fileBind schema.UploadFile
	err := c.ShouldBind(&fileBind)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, fmt.Sprintf("%s", err))
		return
	}
	// 保存文件到缓存目录
	fileName := file.Filename
	filePath := store.CachePath + "/" + fileName
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(400, fmt.Sprintf("%s", err))
		return
	}
	// 计算文件 md5 值
	md5, err := lib.ComputeMd5(filePath)
	if err != nil {
		c.JSON(400, fmt.Sprintf("%s", err))
		return
	}
	// 文件保存到指定目录
	isFileExist := model.IsFileExist(md5)
	newPath := store.UploadPath + "\\" + md5
	if isFileExist == false {
		err = os.Rename(filePath, newPath)
		if err != nil {
			c.JSON(400, gin.H{"msg": fmt.Sprintf("%s", err)})
			return
		}
	}
	// 在缓存中删除文件
	_, err = os.Stat(filePath)
	if err == nil {
		err = os.Remove(filePath)
		if err != nil {
			c.JSON(500, gin.H{"msg": fmt.Sprintf("%s", err)})
			return
		}
	}
	// 创建上传记录
	wareHouseId, _ := c.Get("wareHouseId")
	_, err = model.CreateFile(fileName, md5, newPath, wareHouseId.(int), fileBind.FolderId, true)
	if err != nil {
		c.JSON(400, gin.H{"msg": fmt.Sprintf("%s", err)})
		return
	}
	c.JSON(200, gin.H{"msg": "上传成功"})
}

func DeleteFile(c *gin.Context) {
	var fileBind schema.DeleteFile
	err := c.ShouldBindUri(&fileBind)
	if err != nil {
		c.JSON(400, fmt.Sprintf("%s", err))
		return
	}
	wareHouseId, _ := c.Get("wareHouseId")
	err = model.DeleteFile(fileBind.FileId, wareHouseId.(int))
	if err != nil {
		c.JSON(400, fmt.Sprintf("%s", err))
		return
	}
	c.JSON(200, gin.H{"msg": "删除成功"})
	return
}

func ListFile(c *gin.Context) {
	wareHouseId, _ := c.Get("wareHouseId")
	warehouseId := wareHouseId.(int)
	fileList, err := model.QueryListFile(warehouseId)
	if err != nil {
		c.JSON(400, gin.H{"msg": fmt.Sprintf("%s", err)})
	}
	c.JSON(200, gin.H{"msg": fileList})
}

func DownLoadFile(c *gin.Context) {
	md5 := c.Param("md5")
	wareHouseId, _ := c.Get("wareHouseId")
	warehouseId := wareHouseId.(int)
	// 判断用户是否真拥有此文件
	file, err := model.UserFileExist(warehouseId, md5)
	if err != nil {
		c.JSON(400, gin.H{"msg": fmt.Sprintf("%s", err)})
	}
	filePath := store.UploadPath + "\\" + md5
	// 设置文件名，并解决文件名乱码问题
	c.Header("Content-Disposition",
		fmt.Sprintf("attachment; filename*=utf-8''%s", url.QueryEscape(file.Name)))
	c.File(filePath)
	return
}
