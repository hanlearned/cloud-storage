package service

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func UploadFile(c *gin.Context) {
	// 1. 创建一条数据库上传记录
	// 2. 上传文件到指定目录
	// 3. 计算文件 md5 值
	file, _ := c.FormFile("file")
	file_name := file.Filename
	file_path := "G:\\GoProject\\src\\cloud-storage\\upload\\" + file_name
	c.SaveUploadedFile(file, file_path)

	f, err := os.Open(file_path)
	if nil != err {
		fmt.Println(err)
		c.Status(400)
		return
	}
	md5Handle := md5.New()
	_, err = io.Copy(md5Handle, f)
	if nil != err {
		fmt.Println(err)
		return
	}
	md := md5Handle.Sum(nil)
	md5str := fmt.Sprintf("%x", md)
	c.String(200, md5str)
}
