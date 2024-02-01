package model

import (
	"cloud-storage/model/mysql"
	"errors"
	"fmt"
)

type File struct {
	ID          int
	Name        string
	Md5         string
	Path        string
	WareHouseId int
	FolderId    int
	Status      bool
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
	fmt.Println(queryFile)
	if queryFile.FolderId == 0 {
		return false
	}
	return true
}
