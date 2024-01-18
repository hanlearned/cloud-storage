package model

import (
	"cloud-storage/model/mysql"
	"errors"
)

type Folder struct {
	ID          int
	Name        string
	WareHouseId int
	FolderId    int
	Included    int
}

func CreateFolder(name string, folderId int, warehouseId int) (Folder, error) {
	folder := Folder{
		Name:        name,
		FolderId:    folderId,
		WareHouseId: warehouseId,
		Included:    0,
	}
	// 判断父级文件夹是否存在
	folderExist := QueryFolderByID(folderId, warehouseId)
	if folderExist == false {
		return folder, errors.New("父级文件夹不存在")
	}
	result := mysql.DB.Create(&folder)
	if result.Error != nil {
		return folder, nil
	}
	return folder, nil
}

func QueryFolderByID(folderId int, wareHouse int) bool {
	// 判断文件夹是否存在
	if folderId == 0 {
		return true
	}
	var folder Folder
	mysql.DB.Where("id = ? and ware_house_id = ?", folderId, wareHouse).Find(&folder)
	if folder.ID == 0 {
		return false
	}
	return true
}

func DeleteFolder(folderId int, wareHouseId int) bool {
	var folder Folder

	result := mysql.DB.Where("id = ? and ware_house_id = ?", folderId, wareHouseId).Delete(&folder)
	if result.Error != nil {
		return false
	}
	return true
}

func ListFolder(warehouseId int) []Folder {
	var folderList []Folder
	mysql.DB.Where("ware_house_id = ?", warehouseId).Find(&folderList)
	return folderList
}
