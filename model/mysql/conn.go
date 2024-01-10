package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "root:qwer1234@tcp(47.93.116.220:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	//DB.SingularTable(true)
	if err != nil {
		fmt.Printf("数据库连接失败：%s", err)
	}
}
