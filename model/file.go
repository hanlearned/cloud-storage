package model

import (
	"cloud-storage/model/mysql"
	"errors"
)

type File struct {
	ID          int
	Name        string
	Md5         string
	Path        string
	WareHouseId int
	FolderId    int
	Status      bool
	//CreateTime  datetime.DateTime
}

func CreateFile(
	fileName string, md5 string, savePath string,
	wareHouseId int, folderId int, status bool) (interface{}, error) {
	/*
		1. 判断 md5 是否存在
		2. 判断用户是否有此文件夹
	*/
	var queryFolder Folder
	mysql.DB.Where("id = ?", folderId).Find(&queryFolder)
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
