package schema

type UploadFile struct {
	FolderId int `form:"folder_id" binding:"required"`
}

type DeleteFile struct {
	FileId int `uri:"file_id" binding:"required"`
}
