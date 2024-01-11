package router

import (
	"cloud-storage/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//router := gin.Default()

	// 登录
	router := gin.New()
	user := router.Group("user")
	{
		user.GET("/login", controller.Login)
		// 注册
		user.GET("/register", controller.Register)
	}
	file := router.Group("file")
	{
		file.POST("/upload", controller.Upload)
		// 注册
	}
	return router
}
