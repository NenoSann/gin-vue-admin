package fit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FITUploadRouter struct{}

func (e *FITUploadRouter) InitFITUploadRouter(Router *gin.RouterGroup) {
	fitUploadRouter := Router.Group("fit").Use(middleware.OperationRecord())
	{
		fitUploadRouter.POST("upload", fitFileUploadApi.UploadFitFile)           // 上传FIT文件
		fitUploadRouter.POST("convert_to_csv", fitFileUploadApi.ConvertFitToCSV) // FIT转CSV
		fitUploadRouter.GET("sessions", fitFileUploadApi.GetSessionList)         // 获取会话列表
		fitUploadRouter.GET("session/:id", fitFileUploadApi.GetSessionDetail)    // 获取会话详情
	}
}
