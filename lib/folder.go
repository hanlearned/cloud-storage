package lib

import (
	"os"
)
import "cloud-storage/conf"

var storeConfig = conf.StoreConfig

func CreateFolder(folderName string) error {
	uploadPath := storeConfig.UploadPath + "/" + folderName
	err := os.Mkdir(uploadPath, 0755)
	return err
}
