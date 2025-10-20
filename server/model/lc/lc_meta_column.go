package models

import "time"

type MetaColumn struct {
	ID             uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MetaCode       string    `gorm:"column:meta_code;not null" json:"meta_code"`
	ColumnName     string    `gorm:"column:column_name;not null" json:"column_name"`
	ColumnCode     string    `gorm:"column:column_code;not null" json:"column_code"`
	ColumnType     string    `gorm:"column:column_type;not null" json:"column_type"`
	DisplayType    string    `gorm:"column:display_type;not null" json:"display_type"`
	IsRequired     bool      `gorm:"column:is_required;default:false" json:"is_required"`
	IsSearchable   bool      `gorm:"column:is_searchable;default:true" json:"is_searchable"`
	IsSortable     bool      `gorm:"column:is_sortable;default:true" json:"is_sortable"`
	IsVisible      bool      `gorm:"column:is_visible;default:true" json:"is_visible"`
	DefaultValue   string    `gorm:"column:default_value" json:"default_value"`
	Placeholder    string    `gorm:"column:placeholder" json:"placeholder"`
	DictCode       string    `gorm:"column:dict_code" json:"dict_code"`
	ColumnOrder    int       `gorm:"column:column_order;default:0" json:"column_order"`
	MaxLength      *int      `gorm:"column:max_length" json:"max_length"`
	MinValue       *float64  `gorm:"column:min_value" json:"min_value"`
	MaxValue       *float64  `gorm:"column:max_value" json:"max_value"`
	ValidationRule string    `gorm:"column:validation_rule" json:"validation_rule"`
	Description    string    `gorm:"column:description" json:"description"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (MetaColumn) TableName() string {
	return "meta_column"
}
