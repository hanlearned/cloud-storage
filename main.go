package main

import "cloud-storage/router"
import "cloud-storage/model/mysql"

func main() {
	r := router.SetupRouter()
	mysql.InitDB()
	r.Run(":80")
}
