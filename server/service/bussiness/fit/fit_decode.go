package fit

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/proto"
)

type DecodeService struct{}

var DecodeServiceApp = new(DecodeService)

// RawField 包含原始值及其 Scale/Offset 元数据
type RawField struct {
	Value  any
	Scale  float64
	Offset float64
}

type FitDecoder struct {
	FilePath string
	FIT      *proto.FIT
}

// ReadFromFile 从上传的文件中读取并解析 FIT 数据，返回解析后的 map 数据
// 参数:
//   - file: multipart.File 上传的文件
//   - filename: 文件名（用于记录）
//
// 返回:
//   - map[string]interface{}: 解析后的数据映射
//   - error: 错误信息
func (s *DecodeService) ReadFromFile(file interface{}, filename string) (map[string]interface{}, error) {
	// 创建 buffered reader
	reader := bufio.NewReader(file.(interface{ Read([]byte) (int, error) }))

	// 使用 FIT decoder 解析文件
	fitDecoder := decoder.New(reader)
	fitData, err := fitDecoder.Decode()
	if err != nil {
		log.Printf("Failed to decode FIT file: %v", err)
		return nil, fmt.Errorf("解析FIT文件失败: %w", err)
	}

	// 将解析后的数据转换为 map
	result := make(map[string]interface{})

	// 提取 Session 消息
	sessionData, err := s.extractSessionData(fitData, filename)
	if err != nil {
		return nil, err
	}
	result["session"] = sessionData

	// 提取 Lap 消息列表
	lapsData := s.extractLapsData(fitData)
	result["laps"] = lapsData

	// 提取 Record 消息列表
	recordsData := s.extractRecordsData(fitData)
	result["records"] = recordsData

	// 添加文件元数据
	result["filename"] = filename
	result["messages_count"] = len(fitData.Messages)

	return result, nil
}

// extractSessionData 从 FIT 数据中提取 Session 信息
func (s *DecodeService) extractSessionData(fitData *proto.FIT, filename string) (*Session, error) {
	for _, msg := range fitData.Messages {
		if msg.Num == 18 { // mesgnum.Session = 18
			session, err := ConvertFieldsToSession(msg.Fields, filename)
			if err != nil {
				return nil, fmt.Errorf("转换Session数据失败: %w", err)
			}
			return session, nil
		}
	}
	return nil, fmt.Errorf("未找到Session消息")
}

// extractLapsData 从 FIT 数据中提取所有 Lap 信息
func (s *DecodeService) extractLapsData(fitData *proto.FIT) []map[string]interface{} {
	var laps []map[string]interface{}

	for _, msg := range fitData.Messages {
		if msg.Num == 19 { // mesgnum.Lap = 19
			lapData := make(map[string]interface{})
			for _, field := range msg.Fields {
				lapData[field.Name] = field.Value.Any()
			}
			laps = append(laps, lapData)
		}
	}

	return laps
}

// extractRecordsData 从 FIT 数据中提取所有 Record 信息
func (s *DecodeService) extractRecordsData(fitData *proto.FIT) []map[string]interface{} {
	var records []map[string]interface{}

	for _, msg := range fitData.Messages {
		if msg.Num == 20 { // mesgnum.Record = 20
			recordData := make(map[string]interface{})
			for _, field := range msg.Fields {
				recordData[field.Name] = field.Value.Any()
			}
			records = append(records, recordData)
		}
	}

	return records
}

// FromPath 从指定路径解码 FIT 文件
func (d *FitDecoder) FromPath(path string) error {
	if len(path) == 0 {
		return fmt.Errorf("path is empty")
	}
	d.FilePath = path
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Printf("Failed to open FIT file: %v", err)
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := decoder.New(reader)
	d.FIT, err = decoder.Decode()
	if err != nil {
		log.Printf("Failed to decode FIT file: %v", err)
		return err
	}
	return nil
}

// ReadSession 从已加载的 FIT 数据中读取 Session 信息
func (d *FitDecoder) ReadSession() (*Session, error) {
	if d.FIT == nil {
		return nil, fmt.Errorf("FIT file not loaded")
	}

	// var sessionMsg proto.Message

	// for _, msg := range d.FIT.Messages {
	// 	if msg.Num == mesgnum.Session {
	// 		sessionMsg = msg
	// 		break
	// 	}
	// 	return nil, fmt.Errorf("Session message not found in FIT file")
	// }

	return &Session{}, nil
}

// ConvertDataMapToSession 将从 FIT 消息中解析出的 dataMap 转换为 Session 结构体
// 使用反射动态映射，根据 JSON 标签自动匹配字段
func (d *FitDecoder) ConvertDataMapToSession(dataMap map[string]RawField) (*Session, error) {
	session := &Session{OriginalFilename: filepath.Base(d.FilePath)}
	sessionValue := reflect.ValueOf(session).Elem()
	d.mapEmbeddedStructFields(sessionValue, dataMap)
	return session, nil
}

// mapEmbeddedStructFields 递归遍历结构体（含匿名嵌入）按 json 标签映射赋值
func (d *FitDecoder) mapEmbeddedStructFields(structValue reflect.Value, dataMap map[string]RawField) {
	structType := structValue.Type()
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldValue := structValue.Field(i)

		// 匿名嵌入的结构体，递归展开
		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			d.mapEmbeddedStructFields(fieldValue, dataMap)
			continue
		}

		if !fieldValue.CanSet() {
			continue
		}

		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}
		jsonName := strings.Split(jsonTag, ",")[0]

		// 跳过数据库/关联/元数据字段
		switch jsonName {
		case "id", "created_at", "updated_at", "laps", "records", "user_id", "original_filename", "file_hash":
			continue
		}

		// 特殊映射
		dataKey := jsonName
		switch jsonName {
		case "end_position_lat":
			dataKey = "nec_lat"
		case "end_position_long":
			dataKey = "nec_long"
		}

		rf, ok := dataMap[dataKey]
		if !ok {
			continue
		}

		if err := d.setFieldValue(fieldValue, field, rf, jsonName); err != nil {
			log.Printf("Warning: Failed to set field %s (%s->%s): %v", field.Name, dataKey, jsonName, err)
			continue
		}
	}
}

// setFieldValue 设置字段值，处理各种类型转换
func (d *FitDecoder) setFieldValue(fieldValue reflect.Value, field reflect.StructField, rf RawField, jsonName string) error {
	rawValue := rf.Value
	// 处理 nil 值
	if rawValue == nil {
		return nil
	}

	// 校验 scale/offset
	scale := rf.Scale
	offset := rf.Offset
	if scale == 0 { // 防止除零，按规范默认 1
		log.Printf("Scale=0 detected for field %s, defaulting to 1", jsonName)
		scale = 1
	}
	if scale < 0 { // 不期望出现负 scale
		return fmt.Errorf("invalid negative scale %f", scale)
	}

	fieldType := fieldValue.Type()

	// 特殊处理时间字段
	if fieldType == reflect.TypeOf(time.Time{}) {
		return d.setTimeField(fieldValue, rawValue, jsonName)
	}

	// 处理指针类型
	if fieldType.Kind() == reflect.Ptr {
		return d.setPointerField(fieldValue, fieldType, rawValue, jsonName, scale, offset)
	}

	// 处理基本类型
	return d.setBasicField(fieldValue, fieldType, rawValue, jsonName, scale, offset)
}

// setTimeField 设置时间字段
func (d *FitDecoder) setTimeField(fieldValue reflect.Value, rawValue interface{}, jsonName string) error {
	if timestamp, ok := rawValue.(uint32); ok {
		// FIT 时间戳从 1989-12-31 00:00:00 UTC 开始
		timeValue := time.Unix(int64(timestamp)+631065600, 0)
		fieldValue.Set(reflect.ValueOf(timeValue))
		return nil
	}
	return fmt.Errorf("cannot convert %T to time.Time", rawValue)
}

// setPointerField 设置指针类型字段
func (d *FitDecoder) setPointerField(fieldValue reflect.Value, fieldType reflect.Type, rawValue interface{}, jsonName string, scale, offset float64) error {
	// 获取指针指向的类型
	elemType := fieldType.Elem()

	// 创建新的元素值
	newValue := reflect.New(elemType).Elem()

	// 设置元素值
	if err := d.setBasicField(newValue, elemType, rawValue, jsonName, scale, offset); err != nil {
		return err
	}

	// 创建指针并设置
	ptrValue := reflect.New(elemType)
	ptrValue.Elem().Set(newValue)
	fieldValue.Set(ptrValue)

	return nil
}

// setBasicField 设置基本类型字段
func (d *FitDecoder) setBasicField(fieldValue reflect.Value, fieldType reflect.Type, rawValue interface{}, jsonName string, scale, offset float64) error {
	switch fieldType.Kind() {
	case reflect.Uint32:
		return d.setUint32Field(fieldValue, rawValue, jsonName)
	case reflect.Uint16:
		return d.setUint16Field(fieldValue, rawValue, jsonName)
	case reflect.Uint8:
		return d.setUint8Field(fieldValue, rawValue, jsonName)
	case reflect.Int16:
		return d.setInt16Field(fieldValue, rawValue, jsonName)
	case reflect.Int8:
		return d.setInt8Field(fieldValue, rawValue, jsonName)
	case reflect.Float64:
		// 应用 scale/offset
		scaled := applyScaleOffset(rawValue, scale, offset)
		return d.setFloat64Field(fieldValue, scaled, jsonName)
	case reflect.String:
		if str, ok := rawValue.(string); ok {
			fieldValue.SetString(str)
			return nil
		}
		return fmt.Errorf("cannot convert %T to string", rawValue)
	default:
		return fmt.Errorf("unsupported field type: %s", fieldType.Kind())
	}
}

// applyScaleOffset 根据 FIT 定义执行 (raw/scale) - offset 转换
func applyScaleOffset(raw interface{}, scale, offset float64) interface{} {
	// 无需转换的情况
	if scale == 1 && offset == 0 {
		return raw
	}
	switch v := raw.(type) {
	case uint8:
		return (float64(v) / scale) - offset
	case uint16:
		return (float64(v) / scale) - offset
	case uint32:
		return (float64(v) / scale) - offset
	case int8:
		return (float64(v) / scale) - offset
	case int16:
		return (float64(v) / scale) - offset
	case int32:
		return (float64(v) / scale) - offset
	case float32:
		return (float64(v) / scale) - offset
	case float64:
		return (v / scale) - offset
	default:
		return raw // 非数值类型不处理
	}
}

// setUint32Field 设置 uint32 字段
func (d *FitDecoder) setUint32Field(fieldValue reflect.Value, rawValue interface{}, jsonName string) error {
	switch v := rawValue.(type) {
	case uint32:
		// 特殊处理时间相关字段（毫秒转秒）
		if strings.Contains(jsonName, "time") && !strings.Contains(jsonName, "start_time") && !strings.Contains(jsonName, "timestamp") {
			fieldValue.SetUint(uint64(v / 1000))
		} else {
			fieldValue.SetUint(uint64(v))
		}
		return nil
	case uint16:
		fieldValue.SetUint(uint64(v))
		return nil
	case uint8:
		fieldValue.SetUint(uint64(v))
		return nil
	default:
		return fmt.Errorf("cannot convert %T to uint32", rawValue)
	}
}

// setUint16Field 设置 uint16 字段
func (d *FitDecoder) setUint16Field(fieldValue reflect.Value, rawValue interface{}, jsonName string) error {
	switch v := rawValue.(type) {
	case uint16:
		// 过滤无效值
		if v == 65535 {
			return nil // 不设置无效值
		}
		fieldValue.SetUint(uint64(v))
		return nil
	case uint32:
		if v > 65535 {
			return fmt.Errorf("value %d too large for uint16", v)
		}
		fieldValue.SetUint(uint64(v))
		return nil
	case uint8:
		fieldValue.SetUint(uint64(v))
		return nil
	default:
		return fmt.Errorf("cannot convert %T to uint16", rawValue)
	}
}

// setUint8Field 设置 uint8 字段
func (d *FitDecoder) setUint8Field(fieldValue reflect.Value, rawValue interface{}, jsonName string) error {
	switch v := rawValue.(type) {
	case uint8:
		// 过滤无效值
		if v == 255 {
			return nil // 不设置无效值
		}
		fieldValue.SetUint(uint64(v))
		return nil
	case uint16:
		if v > 255 {
			return fmt.Errorf("value %d too large for uint8", v)
		}
		fieldValue.SetUint(uint64(v))
		return nil
	case uint32:
		if v > 255 {
			return fmt.Errorf("value %d too large for uint8", v)
		}
		fieldValue.SetUint(uint64(v))
		return nil
	default:
		return fmt.Errorf("cannot convert %T to uint8", rawValue)
	}
}

// setInt16Field 设置 int16 字段
func (d *FitDecoder) setInt16Field(fieldValue reflect.Value, rawValue interface{}, jsonName string) error {
	switch v := rawValue.(type) {
	case int16:
		fieldValue.SetInt(int64(v))
		return nil
	case int32:
		if v < -32768 || v > 32767 {
			return fmt.Errorf("value %d out of range for int16", v)
		}
		fieldValue.SetInt(int64(v))
		return nil
	case uint16:
		if v > 32767 {
			return fmt.Errorf("value %d too large for int16", v)
		}
		fieldValue.SetInt(int64(v))
		return nil
	default:
		return fmt.Errorf("cannot convert %T to int16", rawValue)
	}
}

// setInt8Field 设置 int8 字段
func (d *FitDecoder) setInt8Field(fieldValue reflect.Value, rawValue interface{}, jsonName string) error {
	switch v := rawValue.(type) {
	case int8:
		fieldValue.SetInt(int64(v))
		return nil
	case uint8:
		if v > 127 {
			fieldValue.SetInt(int64(int8(v))) // 允许有符号转换
		} else {
			fieldValue.SetInt(int64(v))
		}
		return nil
	case int16:
		if v < -128 || v > 127 {
			return fmt.Errorf("value %d out of range for int8", v)
		}
		fieldValue.SetInt(int64(v))
		return nil
	default:
		return fmt.Errorf("cannot convert %T to int8", rawValue)
	}
}

// setFloat64Field 设置 float64 字段
func (d *FitDecoder) setFloat64Field(fieldValue reflect.Value, rawValue interface{}, jsonName string) error {
	switch v := rawValue.(type) {
	case int32:
		// 特殊处理位置字段（semicircles 转度数）
		if strings.Contains(jsonName, "lat") || strings.Contains(jsonName, "long") {
			degrees := float64(v) * (180.0 / 2147483648.0)
			fieldValue.SetFloat(degrees)
		} else {
			fieldValue.SetFloat(float64(v))
		}
		return nil
	case float64:
		fieldValue.SetFloat(v)
		return nil
	case float32:
		fieldValue.SetFloat(float64(v))
		return nil
	default:
		return fmt.Errorf("cannot convert %T to float64", rawValue)
	}
}

// ConvertFieldsToSession 将 FIT 消息字段直接转换为 Session 结构体（独立函数）
func ConvertFieldsToSession(fields []proto.Field, fileName string) (*Session, error) {
	dataMap := make(map[string]RawField)
	for _, field := range fields {
		// 提取 scale/offset (默认值已在 FieldBase 中给出)
		scale := 1.0
		offset := 0.0
		if field.FieldBase != nil { // 安全检查
			if field.FieldBase.Scale != 0 { // 0 表示默认 1，防止无意除 0
				scale = field.FieldBase.Scale
			}
			offset = field.FieldBase.Offset
		}
		dataMap[field.Name] = RawField{Value: field.Value.Any(), Scale: scale, Offset: offset}
	}
	decoder := &FitDecoder{FilePath: fileName}
	return decoder.ConvertDataMapToSession(dataMap)
}
