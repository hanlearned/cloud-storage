package schema

type Folder struct {
	Name     string `json:"name" binding:"required"`
	FolderId int    `json:"folder_id" binding:"required"`
}
