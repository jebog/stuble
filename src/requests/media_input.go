package requests

type MediaInput struct {
	ModelId   string `json:"model_id" binding:"required"`
	ModelType string `json:"model_type" binding:"required"`
	FileName  string `json:"file_name" binding:"required"`
	MimeType  string `json:"mime_type" binding:"required"`
}
