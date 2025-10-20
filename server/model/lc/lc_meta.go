package models

import "time"

type Meta struct {
	MetaName      string    `gorm:"column:meta_name;not null" json:"meta_name"`
	MetaCode      string    `gorm:"column:meta_code;primaryKey" json:"meta_code"`
	MetaType      string    `gorm:"column:meta_type" json:"meta_type"`
	Visible       bool      `gorm:"column:visible;default:true" json:"visible"`
	MetaTableName string    `gorm:"column:table_name;not null" json:"table_name"`
	Description   string    `gorm:"column:description" json:"description"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (Meta) TableName() string {
	return "meta"
}
