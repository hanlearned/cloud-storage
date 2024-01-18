package model

import "cloud-storage/model/mysql"

type WareHouse struct {
	ID           int
	StorageSpace int
	UsedSpace    float32
	UserID       int
}

func CreateWare(userID int) (WareHouse, error) {
	warehouse := WareHouse{
		UserID:       userID,
		StorageSpace: 500,
		UsedSpace:    0,
	}
	result := mysql.DB.Create(&warehouse)
	if result.Error != nil {
		return warehouse, result.Error
	}
	return warehouse, nil
}
