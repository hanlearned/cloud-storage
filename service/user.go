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
	res := model.CheckoutUserOrPasswd(name, password)
	if res == false {
		c.JSON(402, gin.H{"msg": "用户或密码错误"})
		return
	}
	token, err := lib.GetToken(name)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"msg": "token 生成错误"})
		return
	}
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

	user_create_res, user := model.CreateUser(name, password)
	if user_create_res == false {
		c.JSON(200, gin.H{"msg": "用户注册失败"})
		return
	}

	warehouse_create_res, warehouse := model.CreateWare(user.ID)
	if warehouse_create_res == false {
		c.JSON(200, gin.H{"msg": "用户仓库创建失败"})
		return
	}

	warehouse_update_res := model.SaveUser(user.ID, warehouse.ID)
	if warehouse_update_res == false {
		c.JSON(200, gin.H{"msg": "用户仓库跟新失败"})
		return
	}
	c.JSON(200, gin.H{"msg": "用户注册成功"})
}
