package controller

import (
	"cloud-storage/controller/schema"
	"cloud-storage/lib"
	"cloud-storage/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userInfo schema.UserInfo
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		c.JSON(200, gin.H{"msg": err})
		return
	}
	name := userInfo.Name
	password := userInfo.Password
	fmt.Println(name, password)
	token, err := lib.GetToken(name)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"msg": ""})
		return
	}
	lib.JwtParse(token)
	c.JSON(200, gin.H{
		"token": token,
	})
}

func Register(c *gin.Context) {
	var userInfo schema.RegisterUserInfo
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		c.JSON(200, gin.H{"msg": "sorry input error"})
		return
	}
	name := userInfo.Name
	password := userInfo.Password
	re_password := userInfo.RePassword
	if password != re_password {
		c.JSON(200, gin.H{"msg": "两次密码不一致"})
		return
	}

	registerRes := model.QueryUser(name)
	if registerRes == false {
		c.JSON(200, gin.H{"msg": "注册失败 用户已存在"})
		return
	}
	model.CreateUser(name, password)
	c.JSON(200, gin.H{"msg": "注册成功"})
	return
}
