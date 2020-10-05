package modelapi

import "github.com/dc-lab/sky/internal/data_manager/master/modeldb"

type FileRequest struct {
	*modeldb.File

	ProtectedId   string `json:"-" example:""`
	ProtectedHash string `json:"-"`
}

type FileResponse struct {
	*modeldb.File

	UploadUrl string `json:"upload_url,omitempty" example:"https://sky.sskvor.dev/v1/files/6d83a3d2-16a6-486a-91a2-5d44ba74e326/data?token=9ee2e81f-16a3-46e1-b794-9e8364a90128"`
}

type UploadFileResponse struct {
	Status string `json:"status" example:"ok"`
}
