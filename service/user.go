package service

import (
	"cloud-storage/lib"
	"cloud-storage/model"
	"cloud-storage/service/schema"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userInfo schema.UserInfo
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		c.JSON(200, gin.H{"msg": fmt.Sprintf("%s", err)})
		return
	}
	name := userInfo.Name
	password := userInfo.Password
	userCheck, err := model.CheckoutUserOrPasswd(name, password)
	if err != nil {
		c.JSON(402, gin.H{"msg": fmt.Sprintf("%s", err)})
		return
	}
	token, err := lib.GetToken(userCheck.ID)
	if err != nil {
		c.JSON(200, gin.H{"msg": "token 生成错误"})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

func Register(c *gin.Context) {
	// 校验 json 数据
	var userInfo schema.RegisterUserInfo
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		c.JSON(200, gin.H{"msg": "sorry input error"})
		return
	}
	// 校验密码
	name := userInfo.Name
	password := userInfo.Password
	rePassword := userInfo.RePassword
	if password != rePassword {
		c.JSON(200, gin.H{"msg": "两次密码不一致"})
		return
	}
	// 判断用户是否存在
	registerRes := model.IfUserExist(name)
	if registerRes == false {
		c.JSON(200, gin.H{"msg": "注册失败 用户已存在"})
		return
	}
	// 创建用户
	user, err := model.CreateUser(name, password)
	if err != nil {
		c.JSON(200, gin.H{"msg": fmt.Sprintf("%s", err)})
		return
	}
	// 创建仓库
	warehouse, err := model.CreateWare(user.ID)
	if err != nil {
		c.JSON(200, gin.H{"msg": fmt.Sprintf("%s", err)})
		return
	}
	// 关联仓库
	warehouseUpdateRes := model.SaveUser(user.ID, warehouse.ID)
	if warehouseUpdateRes != nil {
		c.JSON(200, gin.H{"msg": fmt.Sprintf("%s", err)})
		return
	}
	c.JSON(200, gin.H{"msg": "用户注册成功"})
}
