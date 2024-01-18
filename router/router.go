package router

import (
	"cloud-storage/middleware"
	"cloud-storage/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//router := gin.Default()
	// 登录
	router := gin.New()
	user := router.Group("user")
	{
		user.GET("/login", service.Login)
		// 注册
		user.GET("/register", service.Register)
	}

	folder := router.Group("folder")
	folder.Use(middleware.CheckLogin)
	{
		folder.POST("/create_folder", service.CreateFolder)
		folder.DELETE("/delete_folder/:folder_id/", service.DeleteFolder)
	}

	file := router.Group("file")
	file.Use(middleware.CheckLogin)
	{
		file.POST("/upload", service.Upload)
		// 注册
	}
	return router
}
