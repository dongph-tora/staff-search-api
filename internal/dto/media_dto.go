package dto

type UploadURLRequest struct {
	FileName    string `json:"file_name"`
	ContentType string `json:"content_type"`
	Folder      string `json:"folder"`
}

type UploadURLResponse struct {
	UploadURL string `json:"upload_url"`
	FileKey   string `json:"file_key"`
	PublicURL string `json:"public_url"`
	ExpiresIn int    `json:"expires_in"`
}

type DeleteMediaRequest struct {
	FileKey string `json:"file_key"`
}
