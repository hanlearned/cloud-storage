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
	return router
}
