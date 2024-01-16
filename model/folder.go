package model

import (
	"cloud-storage/model/mysql"
)

type Folder struct {
	ID          int
	Name        string
	WareHouseId int
	FolderId    int
	Included    int
}

func CreateFolder(name string, folderId int, warehouseId int, included int) (bool, Folder) {
	folder := Folder{
		Name:        name,
		FolderId:    folderId,
		WareHouseId: warehouseId,
		Included:    included,
	}
	// 判断父级文件夹是否存在
	folderExist := QueryFolderByID(folderId)
	if folderExist == false {
		return false, folder
	}
	result := mysql.DB.Create(&folder)
	if result.Error != nil {
		return false, folder
	}
	return true, folder
}

func QueryFolderByID(folderId int) bool {
	// 判断文件夹是否存在
	if folderId == 0 {
		return true
	}
	folder := Folder{}
	mysql.DB.Where("id = ?", folderId).Find(&folder)
	if folder.ID == 0 {
		return false
	}
	return true
}

func DeleteFolder(folderId int) bool {
	folder := Folder{}
	result := mysql.DB.Where("id = ?", folderId).Delete(&folder)
	if result.Error != nil {
		return false
	}
	return true
}
