package lib

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func ComputeMd5(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	md5Handle := md5.New()
	_, err = io.Copy(md5Handle, f)
	if err != nil {
		return "", err
	}
	md := md5Handle.Sum(nil)
	md5str := fmt.Sprintf("%x", md)
	return md5str, nil
}
