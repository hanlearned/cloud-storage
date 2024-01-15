package model

import (
	"cloud-storage/model/mysql"
	"fmt"
)

type UserInfo struct {
	ID          int
	Name        string
	Password    string
	WareHouseId int
}

func CreateUser(name string, password string) (bool, UserInfo) {
	user := UserInfo{
		Name:     name,
		Password: password,
	}
	result := mysql.DB.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false, user
	}
	return true, user
}

func SaveUser(userId int, warehouseId int) bool {
	user := UserInfo{
		WareHouseId: warehouseId,
	}
	result := mysql.DB.Where("id = ?", userId).Find(&user).Update("WareHouseId", warehouseId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	}
	return true
}

func IfUserExist(name string) bool {
	user := UserInfo{
		Name: name,
	}
	result := mysql.DB.Where("name = ?", name).Find(&user)
	fmt.Println(user)
	if result.Error != nil {
		fmt.Println(result.Error)
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
		fmt.Println(result.Error)
	}
	if user.ID == 0 {
		return false
	}
	return true
}

func QueryUserWare(name string) UserInfo {
	user := UserInfo{
		Name: name,
	}
	result := mysql.DB.Where("name", name).Find(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return user
}
