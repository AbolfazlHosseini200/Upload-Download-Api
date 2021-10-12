package Models
type Files struct {
	FileId string `json:"fileId" binding:"required"`
	Username string `json:"username" binding:"required"`
}