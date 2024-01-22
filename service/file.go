package service

import (
	"cloud-storage/lib"
	"cloud-storage/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	// 1. 创建一条数据库上传记录
	// 2. 上传文件到指定目录
	// 3. 计算文件 md5 值
	file, _ := c.FormFile("file")
	fileName := file.Filename
	filePath := "D:\\GoProject\\src\\cloud-storage\\upload\\" + fileName
	err := c.SaveUploadedFile(file, filePath)
	if nil != err {
		c.JSON(400, fmt.Sprintf("%s", err))
		return
	}
	md5, err := lib.ComputeMd5(filePath)
	if err != nil {
		c.JSON(400, fmt.Sprintf("%s", err))
	}
	wareHouseId, _ := c.Get("wareHouseId")

	model.CreateFile(fileName, md5, filePath, wareHouseId.(int), 1, true)

}
