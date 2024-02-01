package main

import (
	"cloud-storage/conf"
	"cloud-storage/model/mysql"
	"cloud-storage/router"
)

func main() {
	conf.InitConfig()
	mysql.InitDB()
	r := router.SetupRouter()
	r.Run(":80")
}
