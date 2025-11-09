package fit

import (
	"path/filepath"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FitFileUploadApi struct{}

const FIT_FILE_FIELD = "file"

var fitService = service.ServiceGroupApp.FitServiceGroup

// UploadFitFile 上传并解析 FIT 文件
// @Tags      FIT
// @Summary   上传FIT文件
// @Description 上传FIT文件并解析其中的运动数据(Session/Lap/Record)
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file  true  "FIT运动数据文件"
// @Success   200   {object}  response.Response{msg=string}  "上传成功"
// @Failure   400   {object}  response.Response{msg=string}  "上传失败"
// @Router    /fit/upload [post]
func (f *FitFileUploadApi) UploadFitFile(c *gin.Context) {
	// 1. 获取当前用户信息
	userUUID := utils.GetUserUuid(c)

	// 记录用户操作日志
	global.GVA_LOG.Info("用户上传FIT文件",
		zap.String("userUUID", userUUID.String()),
	)

	// 2. 接收文件
	file, header, err := c.Request.FormFile(FIT_FILE_FIELD)
	if err != nil {
		global.GVA_LOG.Error("接收FIT文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	defer file.Close()

	// 3. 调用 Service 处理文件上传和解析
	parsedData, err := fitService.FITFileUploadService.UploadFitFile(file, header, userUUID)
	if err != nil {
		global.GVA_LOG.Error("解析FIT文件失败!", zap.Error(err))
		response.FailWithMessage("解析文件失败: "+err.Error(), c)
		return
	}

	// 4. 返回解析后的数据给前端
	response.OkWithData(parsedData, c)
}

// GetSessionList 获取会话列表
// @Tags      FIT
// @Summary   获取运动会话列表
// @Description 分页获取用户的运动会话记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     page      query  int  false  "页码"        default(1)
// @Param     pageSize  query  int  false  "每页数量"    default(10)
// @Success   200  {object}  response.Response{msg=string}  "获取成功"
// @Router    /fit/sessions [get]
func (f *FitFileUploadApi) GetSessionList(c *gin.Context) {
	// TODO: 实现获取会话列表逻辑
	response.OkWithMessage("获取会话列表", c)
}

// GetSessionDetail 获取会话详情
// @Tags      FIT
// @Summary   获取单个运动会话详情
// @Description 获取指定会话的完整数据,包括Lap和Record
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  path  int  true  "会话ID"
// @Success   200  {object}  response.Response{msg=string}  "获取成功"
// @Router    /fit/session/{id} [get]
func (f *FitFileUploadApi) GetSessionDetail(c *gin.Context) {
	// TODO: 实现获取会话详情逻辑
	id := c.Param("id")
	response.OkWithMessage("获取会话详情: "+id, c)
}

// ConvertFitToCSV 将 FIT 文件转换为 CSV 格式
// @Tags      FIT
// @Summary   将FIT文件转换为CSV格式
// @Description 上传FIT文件并将其内容转换为CSV格式返回
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   text/csv
// @Param     file  formData  file  true  "FIT运动数据文件"
// @Success   200   {file}    file  "转换成功，返回CSV文件"
// @Failure   400   {object}  response.Response{msg=string}  "转换失败"
// @Router    /fit/convert_to_csv [post]
func (f *FitFileUploadApi) ConvertFitToCSV(c *gin.Context) {
	// 1. 接收文件
	file, header, err := c.Request.FormFile(FIT_FILE_FIELD)
	if err != nil {
		global.GVA_LOG.Error("接收FIT文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	defer file.Close()

	// 2. 调用 DecodeService 转换为 CSV
	csvData, err := fitService.DecodeService.ConvertFitToCSV(file)
	if err != nil {
		global.GVA_LOG.Error("转换FIT文件为CSV失败!", zap.Error(err))
		response.FailWithMessage("转换文件失败: "+err.Error(), c)
		return
	}

	// 3. 设置响应头并返回 CSV 文件
	c.Header("Content-Disposition", "attachment; filename="+filepath.Base(header.Filename)+".csv")
	c.Data(200, "text/csv", csvData)
}
