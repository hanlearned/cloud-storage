package main

import "cloud-storage/router"
import "cloud-storage/model/mysql"
import "cloud-storage/conf"

func main() {
	mysql.InitDB()
	conf.InitConfig()
	r := router.SetupRouter()
	r.Run(":80")
}
