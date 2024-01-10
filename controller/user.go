package controller

import (
	"cloud-storage/controller/schema"
	"cloud-storage/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userInfo schema.UserInfo
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		c.JSON(200, gin.H{"msg": "sorry input error"})
		return
	}
	name := userInfo.Name
	password := userInfo.Password
	fmt.Println(name, password)
	c.JSON(200, userInfo)
	//context.String(http.StatusOK, "hello world")
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
	fmt.Println(name)
	model.QueryUser(name)
	//if res.ID != nil {
	//	c.JSON(200, gin.H{"msg": "注册失败 用户已存在"})
	//}

	//fmt.Println(name, password, re_password)
	//model.CreateUser(name, password)
	c.JSON(200, gin.H{"msg": "注册成功"})
}
