package fit

import (
	"time"

	fitService "github.com/flipped-aurora/gin-vue-admin/server/service/bussiness/fit"
)

// ================== 数据库类型:在业务类型基础上增加 ID / 外键 / 元数据 ==================
type Record struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	SessionID uint `gorm:"not null;index" json:"session_id"`
	fitService.RecordData
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Lap struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	SessionID uint `gorm:"not null;index" json:"session_id"`
	fitService.LapData
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Session struct {
	ID               uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID           *uint  `gorm:"index" json:"user_id,omitempty"`
	OriginalFilename string `gorm:"size:255;not null" json:"original_filename"`
	FileHash         string `gorm:"size:64;uniqueIndex" json:"file_hash"`
	fitService.SessionData
	// 关联(数据库层使用 DB 类型)
	Laps      []Lap     `gorm:"foreignKey:SessionID" json:"laps,omitempty"`
	Records   []Record  `gorm:"foreignKey:SessionID" json:"records,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (Session) TableName() string { return "sessions" }
func (Record) TableName() string  { return "records" }
func (Lap) TableName() string     { return "laps" }
