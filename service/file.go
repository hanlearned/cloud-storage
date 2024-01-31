package service

import (
	"cloud-storage/conf"
	"cloud-storage/lib"
	"cloud-storage/model"
	"cloud-storage/service/schema"
	"fmt"
	"github.com/gin-gonic/gin"
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
	filePath := store.UploadPath + fileName
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(400, fmt.Sprintf("%s", err))
		return
	}

	md5, err := lib.ComputeMd5(filePath)
	if err != nil {
		c.JSON(400, fmt.Sprintf("%s", err))
	}

	wareHouseId, _ := c.Get("wareHouseId")
	_, err = model.CreateFile(fileName, md5, filePath, wareHouseId.(int), fileBind.FolderId, true)
	if err != nil {
		c.JSON(400, gin.H{"msg": fmt.Sprintf("%s", err)})
	}
}
