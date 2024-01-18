package model

import "google.golang.org/genproto/googleapis/type/datetime"

type File struct {
	ID         int
	Name       string
	Md5        string
	Path       string
	UploadTime datetime.DateTime
}
