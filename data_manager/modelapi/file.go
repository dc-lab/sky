package modelapi

import "data_manager/modeldb"

type FileRequest struct {
	*modeldb.File

	ProtectedId   string `json:"id,omitempty"`
	ProtectedHash string `json:"hash,omitempty"`
}

type FileResponse struct {
	*modeldb.File

	UploadUrl string `json:"upload_url,omitempty"`
}

type UploadFileResponse struct {
	Status string `json:"status"`
}
