package model

import (
	"cloud-storage/model/mysql"
	"fmt"
)

type UserInfo struct {
	ID       int
	Name     string
	Password string
}

func CreateUser(name string, password string) {
	user := UserInfo{
		Name:     name,
		Password: password,
	}
	result := mysql.DB.Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(user.ID)
}

func QueryUser(name string) bool {
	user := UserInfo{
		Name: name,
	}
	result := mysql.DB.Where("name = ?", name).Find(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	if user.ID == 0 {
		return false
	}
	return true
}
