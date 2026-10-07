package dto

type UploadResponse struct {
	FileID   string `json:"fileId"`
	Filename string `json:"filename"`
}

type ConvertResponse struct {
	ConvertedFileID string `json:"convertedFileId"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
