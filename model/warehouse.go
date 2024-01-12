package model

import "fmt"
import "cloud-storage/model/mysql"

type WareHouse struct {
	ID           int
	StorageSpace int
	UsedSpace    float32
	UserID       int
}

func CreateWare(userID int) (bool, WareHouse) {
	warehouse := WareHouse{
		UserID:       userID,
		StorageSpace: 500,
		UsedSpace:    0,
	}
	result := mysql.DB.Create(&warehouse)

	if result.Error != nil {
		fmt.Println(result.Error)
		return false, warehouse
	}
	return true, warehouse
}
