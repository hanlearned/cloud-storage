package service

import (
	"cloud-storage/conf"
	"cloud-storage/lib"
	"cloud-storage/model"
	"cloud-storage/service/schema"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var store = conf.StoreConfig

func UploadFile(c *gin.Context) {
	var fileBind schema.File
	err := c.ShouldBind(&fileBind)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, fmt.Sprintf("%s", err))
		return
	}

	fileName := file.Filename
	filePath := store.CachePath + "/" + fileName
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(400, fmt.Sprintf("%s", err))
		return
	}

	md5, err := lib.ComputeMd5(filePath)
	if err != nil {
		c.JSON(400, fmt.Sprintf("%s", err))
	}

	isFileExist := model.IsFileExist(md5)
	newPath := store.UploadPath + "\\" + md5
	if isFileExist == false {
		err = os.Rename(filePath, newPath)
		if err != nil {
			c.JSON(400, gin.H{"msg": fmt.Sprintf("%s", err)})
		}
	}
	wareHouseId, _ := c.Get("wareHouseId")
	_, err = model.CreateFile(fileName, md5, newPath, wareHouseId.(int), fileBind.FolderId, true)
	if err != nil {
		c.JSON(400, gin.H{"msg": fmt.Sprintf("%s", err)})
	}
}
