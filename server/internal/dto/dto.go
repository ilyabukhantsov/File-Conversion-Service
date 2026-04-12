package dto

type UploadResponse struct {
	FileId string `json:"fileid"`
}

type ConvertResponse struct {
	ConvertedFileId string `json:"convertedFileId"`
	Status          string `json:"status"`
}

type Error struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}
