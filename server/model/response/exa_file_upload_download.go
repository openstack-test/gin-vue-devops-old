package response

import "gin-vue-devops/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
