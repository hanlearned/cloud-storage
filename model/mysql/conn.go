package mysql

import (
	"cloud-storage/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var dbConfig = conf.DBConfig

func InitDB() {
	var err error
	host := dbConfig.Host
	port := dbConfig.Port
	user := dbConfig.User
	password := dbConfig.Password
	db := dbConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, db)
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	//DB.SingularTable(true)
	if err != nil {
		fmt.Printf("数据库连接失败：%s", err)
	}
}
