package model

import (
	"cloud-storage/model/mysql"
	"fmt"
)

type Folder struct {
	ID          int
	Name        string
	WarehouseId int
	FolderId    int
	Included    int
}

func CreateFolder(name string, folderId int, warehouseId int, included int) (bool, Folder) {
	folder := Folder{
		Name:        name,
		FolderId:    folderId,
		WarehouseId: warehouseId,
		Included:    included,
	}
	result := mysql.DB.Create(&folder)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false, folder
	}
	return true, folder
}
