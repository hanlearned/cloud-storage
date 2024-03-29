package model

import (
	"cloud-storage/model/mysql"
	"errors"
	"fmt"
	"time"
)

type File struct {
	ID          int
	Name        string
	Md5         string
	Path        string
	WareHouseId int
	FolderId    int
	Status      bool
	CreateTime  time.Time
}

func CreateFile(
	fileName string, md5 string, savePath string,
	wareHouseId int, folderId int, status bool) (interface{}, error) {
	var queryFolder Folder
	mysql.DB.Where("id = ?", folderId).Find(&queryFolder)
	fmt.Println(queryFolder)
	if queryFolder.ID == 0 {
		var err = errors.New("文件夹 ID 不存在")
		return queryFolder, err
	}

	file := File{
		Name:        fileName,
		Md5:         md5,
		Path:        savePath,
		WareHouseId: wareHouseId,
		FolderId:    folderId,
		Status:      status,
		CreateTime:  time.Now(),
	}
	result := mysql.DB.Create(&file)
	if result.Error != nil {
		return file, result.Error
	}
	return file, nil
}

func IsFileExist(md5 string) bool {
	var queryFile File
	mysql.DB.Where("md5 = ?", md5).Find(&queryFile)
	if queryFile.ID == 0 {
		return false
	}
	return true
}

func DeleteFile(fileId int, wareHouseId int) error {
	var queryFile File
	res := mysql.DB.Model(&queryFile).Where("id = ? and ware_house_id = ?",
		fileId, wareHouseId).Update("is_delete", 1)
	return res.Error
}

func QueryListFile(wareHouseId int) ([]File, error) {
	var fileList []File
	res := mysql.DB.Where("ware_house_id = ? and is_delete = 0", wareHouseId).Find(&fileList)
	return fileList, res.Error
}

func UserFileExist(wareHouseId int, md5 string) (File, error) {
	var file File
	res := mysql.DB.Where("ware_house_id = ? and md5 = ?", wareHouseId, md5).Find(&file)
	if res.Error != nil {
		return file, res.Error
	}
	return file, nil
}
