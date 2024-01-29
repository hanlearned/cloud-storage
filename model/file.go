package model

import (
	"cloud-storage/model/mysql"
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
	//CreateTime  datetime.DateTime
}

func CreateFile(
	fileName string, md5 string, savePath string,
	wareHouseId int, folderId int, status bool) (File, error) {
	var queryFile File
	mysql.DB.Where("md5 = ?", md5).Find(&queryFile)
	fmt.Println(queryFile)

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
