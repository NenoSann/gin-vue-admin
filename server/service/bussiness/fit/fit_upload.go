package fit

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type FITFileUploadService struct{}

var FITFileUploadServiceApp = new(FITFileUploadService)

// UploadFitFile 上传FIT文件并解析，返回解析后的数据
// 参数:
//   - file: multipart.File 上传的文件
//   - header: *multipart.FileHeader 文件头信息
//   - userUUID: uuid.UUID 用户UUID
//
// 返回:
//   - map[string]interface{}: 解析后的数据
//   - error: 错误信息
func (s *FITFileUploadService) UploadFitFile(
	file multipart.File,
	header *multipart.FileHeader,
	userUUID uuid.UUID,
) (map[string]interface{}, error) {
	// 调用 DecodeService 解析文件
	parsedData, err := DecodeServiceApp.ReadFromFile(file, header.Filename)
	if err != nil {
		return nil, err
	}

	// 添加用户信息到返回数据中
	parsedData["user_uuid"] = userUUID.String()
	parsedData["uploaded_filename"] = header.Filename
	parsedData["file_size"] = header.Size

	// TODO: 这里可以添加保存到数据库的逻辑
	// 例如：保存 Session、Laps、Records 到数据库

	return parsedData, nil
}
