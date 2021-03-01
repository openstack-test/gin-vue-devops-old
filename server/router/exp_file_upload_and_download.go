package router

import (
	v1 "gin-vue-devops/api/v1"

	"github.com/gin-gonic/gin"
)

func InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	FileUploadAndDownloadGroup := Router.Group("fileUploadAndDownload")
	{
		FileUploadAndDownloadGroup.POST("/upload", v1.UploadFile)       // 上传文件
		FileUploadAndDownloadGroup.POST("/getFileList", v1.GetFileList) // 获取上传文件列表
		FileUploadAndDownloadGroup.POST("/deleteFile", v1.DeleteFile)   // 删除指定文件
	}
}
