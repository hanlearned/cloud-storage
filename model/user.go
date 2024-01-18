package model

import (
	"cloud-storage/model/mysql"
)

type UserInfo struct {
	ID          int
	Name        string
	Password    string
	WareHouseId int
}

func CreateUser(name string, password string) (UserInfo, error) {
	user := UserInfo{
		Name:     name,
		Password: password,
	}
	result := mysql.DB.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func SaveUser(userId int, warehouseId int) error {
	user := UserInfo{
		WareHouseId: warehouseId,
	}
	result := mysql.DB.Where("id = ?", userId).Find(&user).Update("WareHouseId", warehouseId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func IfUserExist(name string) bool {
	user := UserInfo{
		Name: name,
	}
	result := mysql.DB.Where("name = ?", name).Find(&user)
	if result.Error != nil {
		return false
	}
	if user.ID == 0 {
		return true
	}
	return false
}

func CheckoutUserOrPasswd(name string, passwd string) bool {
	user := UserInfo{
		Name:     name,
		Password: passwd,
	}
	result := mysql.DB.Where("name = ? and password = ?", name, passwd).Find(&user)
	if result.Error != nil {
		return false
	}
	if user.ID == 0 {
		return false
	}
	return true
}

func QueryUserWare(name string) (UserInfo, error) {
	user := UserInfo{
		Name: name,
	}
	result := mysql.DB.Where("name", name).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
