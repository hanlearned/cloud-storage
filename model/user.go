package model

import (
	"cloud-storage/model/mysql"
	"errors"
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
	var user UserInfo
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

func CheckoutUserOrPasswd(name string, passwd string) (UserInfo, error) {
	var user UserInfo
	result := mysql.DB.Where("name = ? and password = ?", name, passwd).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	var err error
	if user.ID == 0 {
		err = errors.New("用户不存在")
		return user, err
	}
	return user, nil
}

func QueryUserWare(userID interface{}) (UserInfo, error) {
	var user UserInfo
	uid := int(userID.(float64))
	result := mysql.DB.Where("id", uid).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
