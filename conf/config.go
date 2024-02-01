package conf

import (
	"github.com/go-ini/ini"
)

var DBConfig = new(MySQL)

var StoreConfig = new(StoreConf)

type MySQL struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type StoreConf struct {
	UploadPath string `ini:"upload_path"`
	CachePath  string `ini:"cache_path"`
}

func LoadConfig() *ini.File {
	var configPath = "G:\\GoProject\\src\\cloud-storage\\conf\\conf.ini"
	cfg, err := ini.Load(configPath)
	if err != nil {
		panic(err)
	}
	return cfg
}

func MySQLConfig() {
	cfg := LoadConfig()
	err := cfg.Section("MySQL").MapTo(DBConfig)
	if err != nil {
		panic(err)
	}
}

func Store() {
	cfg := LoadConfig()
	err := cfg.Section("Store").MapTo(StoreConfig)
	if err != nil {
		panic(err)
	}
}

func InitConfig() {
	MySQLConfig()
	Store()
}
