package fit

import (
	"time"
)

// ================== 程序类型：仅包含 FIT 原始或计算得到的业务字段 ==================

// RecordData 代表单个记录点的业务数据（无数据库字段）
type RecordData struct {
	Timestamp     time.Time `json:"timestamp"`
	PositionLat   *float64  `json:"position_lat,omitempty"`
	PositionLong  *float64  `json:"position_long,omitempty"`
	Distance      *uint32   `json:"distance,omitempty"`
	Altitude      *uint16   `json:"altitude,omitempty"`
	Speed         *uint16   `json:"speed,omitempty"`
	Power         *uint16   `json:"power,omitempty"`
	Cadence       *uint8    `json:"cadence,omitempty"`
	HeartRate     *uint8    `json:"heart_rate,omitempty"`
	Temperature   *int8     `json:"temperature,omitempty"`
	EnhancedSpeed *uint16   `json:"enhanced_speed,omitempty"`
}

// LapData 代表圈数据的业务字段（无数据库字段）
type LapData struct {
	LapNumber    uint16    `json:"lap_number"`
	Timestamp    time.Time `json:"timestamp"`
	StartTime    time.Time `json:"start_time"`
	MessageIndex uint16    `json:"message_index"`

	StartPositionLat  *float64 `json:"start_position_lat,omitempty"`
	StartPositionLong *float64 `json:"start_position_long,omitempty"`
	EndPositionLat    *float64 `json:"end_position_lat,omitempty"`
	EndPositionLong   *float64 `json:"end_position_long,omitempty"`

	TotalElapsedTime uint32 `json:"total_elapsed_time"`
	TotalTimerTime   uint32 `json:"total_timer_time"`
	TotalMovingTime  uint32 `json:"total_moving_time"`

	TotalDistance    uint32  `json:"total_distance"`
	AvgSpeed         *uint16 `json:"avg_speed,omitempty"`
	MaxSpeed         *uint16 `json:"max_speed,omitempty"`
	EnhancedAvgSpeed *uint32 `json:"enhanced_avg_speed,omitempty"`
	EnhancedMaxSpeed *uint32 `json:"enhanced_max_speed,omitempty"`

	AvgAltitude         *uint16 `json:"avg_altitude,omitempty"`
	MinAltitude         *uint16 `json:"min_altitude,omitempty"`
	MaxAltitude         *uint16 `json:"max_altitude,omitempty"`
	EnhancedAvgAltitude *uint32 `json:"enhanced_avg_altitude,omitempty"`
	EnhancedMinAltitude *uint32 `json:"enhanced_min_altitude,omitempty"`
	EnhancedMaxAltitude *uint32 `json:"enhanced_max_altitude,omitempty"`
	TotalAscent         *uint16 `json:"total_ascent,omitempty"`
	TotalDescent        *uint16 `json:"total_descent,omitempty"`

	AvgGrade            *int16 `json:"avg_grade,omitempty"`
	AvgPosGrade         *int16 `json:"avg_pos_grade,omitempty"`
	AvgNegGrade         *int16 `json:"avg_neg_grade,omitempty"`
	MaxPosGrade         *int16 `json:"max_pos_grade,omitempty"`
	MaxNegGrade         *int16 `json:"max_neg_grade,omitempty"`
	AvgPosVerticalSpeed *int16 `json:"avg_pos_vertical_speed,omitempty"`
	MaxNegVerticalSpeed *int16 `json:"max_neg_vertical_speed,omitempty"`

	AvgPower                    *uint16 `json:"avg_power,omitempty"`
	MaxPower                    *uint16 `json:"max_power,omitempty"`
	NormalizedPower             *uint16 `json:"normalized_power,omitempty"`
	AvgLeftTorqueEffectiveness  *uint8  `json:"avg_left_torque_effectiveness,omitempty"`
	AvgRightTorqueEffectiveness *uint8  `json:"avg_right_torque_effectiveness,omitempty"`
	AvgLeftPedalSmoothness      *uint8  `json:"avg_left_pedal_smoothness,omitempty"`
	AvgRightPedalSmoothness     *uint8  `json:"avg_right_pedal_smoothness,omitempty"`
	LeftRightBalance            *uint16 `json:"left_right_balance,omitempty"`

	AvgHeartRate *uint8 `json:"avg_heart_rate,omitempty"`
	MinHeartRate *uint8 `json:"min_heart_rate,omitempty"`
	MaxHeartRate *uint8 `json:"max_heart_rate,omitempty"`

	AvgCadence *uint8 `json:"avg_cadence,omitempty"`
	MaxCadence *uint8 `json:"max_cadence,omitempty"`

	AvgTemperature *int8 `json:"avg_temperature,omitempty"`
	MaxTemperature *int8 `json:"max_temperature,omitempty"`

	TotalCalories *uint16 `json:"total_calories,omitempty"`

	Event      *int8 `json:"event,omitempty"`
	EventType  *int8 `json:"event_type,omitempty"`
	Sport      *int8 `json:"sport,omitempty"`
	SubSport   *int8 `json:"sub_sport,omitempty"`
	LapTrigger *int8 `json:"lap_trigger,omitempty"`
}

// SessionData 代表完整会话的业务字段（无数据库/元数据）
type SessionData struct {
	Timestamp        time.Time `json:"timestamp"`
	StartTime        time.Time `json:"start_time"`
	TotalElapsedTime uint32    `json:"total_elapsed_time"`
	TotalTimerTime   uint32    `json:"total_timer_time"`
	TotalMovingTime  uint32    `json:"total_moving_time"`

	StartPositionLat  *float64 `json:"start_position_lat,omitempty"`
	StartPositionLong *float64 `json:"start_position_long,omitempty"`
	EndPositionLat    *float64 `json:"end_position_lat,omitempty"`
	EndPositionLong   *float64 `json:"end_position_long,omitempty"`

	TotalDistance    uint32  `json:"total_distance"`
	AvgSpeed         *uint16 `json:"avg_speed,omitempty"`
	MaxSpeed         *uint16 `json:"max_speed,omitempty"`
	EnhancedAvgSpeed *uint32 `json:"enhanced_avg_speed,omitempty"`
	EnhancedMaxSpeed *uint32 `json:"enhanced_max_speed,omitempty"`

	AvgAltitude         *uint16 `json:"avg_altitude,omitempty"`
	MinAltitude         *uint16 `json:"min_altitude,omitempty"`
	MaxAltitude         *uint16 `json:"max_altitude,omitempty"`
	EnhancedAvgAltitude *uint32 `json:"enhanced_avg_altitude,omitempty"`
	EnhancedMinAltitude *uint32 `json:"enhanced_min_altitude,omitempty"`
	EnhancedMaxAltitude *uint32 `json:"enhanced_max_altitude,omitempty"`
	TotalAscent         *uint16 `json:"total_ascent,omitempty"`
	TotalDescent        *uint16 `json:"total_descent,omitempty"`

	AvgGrade            *int16 `json:"avg_grade,omitempty"`
	AvgPosGrade         *int16 `json:"avg_pos_grade,omitempty"`
	AvgNegGrade         *int16 `json:"avg_neg_grade,omitempty"`
	MaxPosGrade         *int16 `json:"max_pos_grade,omitempty"`
	MaxNegGrade         *int16 `json:"max_neg_grade,omitempty"`
	AvgPosVerticalSpeed *int16 `json:"avg_pos_vertical_speed,omitempty"`
	AvgNegVerticalSpeed *int16 `json:"avg_neg_vertical_speed,omitempty"`
	MaxPosVerticalSpeed *int16 `json:"max_pos_vertical_speed,omitempty"`
	MaxNegVerticalSpeed *int16 `json:"max_neg_vertical_speed,omitempty"`

	AvgPower        *uint16 `json:"avg_power,omitempty"`
	MaxPower        *uint16 `json:"max_power,omitempty"`
	NormalizedPower *uint16 `json:"normalized_power,omitempty"`

	AvgHeartRate *uint8 `json:"avg_heart_rate,omitempty"`
	MinHeartRate *uint8 `json:"min_heart_rate,omitempty"`
	MaxHeartRate *uint8 `json:"max_heart_rate,omitempty"`

	AvgCadence *uint8 `json:"avg_cadence,omitempty"`
	MaxCadence *uint8 `json:"max_cadence,omitempty"`

	AvgTemperature *int8 `json:"avg_temperature,omitempty"`
	MaxTemperature *int8 `json:"max_temperature,omitempty"`

	AvgLeftTorqueEffectiveness  *uint8 `json:"avg_left_torque_effectiveness,omitempty"`
	AvgRightTorqueEffectiveness *uint8 `json:"avg_right_torque_effectiveness,omitempty"`
	AvgLeftPedalSmoothness      *uint8 `json:"avg_left_pedal_smoothness,omitempty"`
	AvgRightPedalSmoothness     *uint8 `json:"avg_right_pedal_smoothness,omitempty"`
	LeftRightBalance            *uint8 `json:"left_right_balance,omitempty"`

	TotalCalories *uint16 `json:"total_calories,omitempty"`
	MessageIndex  *uint16 `json:"message_index,omitempty"`

	Event      *int8 `json:"event,omitempty"`
	EventType  *int8 `json:"event_type,omitempty"`
	Sport      *int8 `json:"sport,omitempty"`
	SubSport   *int8 `json:"sub_sport,omitempty"`
	LapTrigger *int8 `json:"lap_trigger,omitempty"`
}

// ================== 数据库类型：在业务类型基础上增加 ID / 外键 / 元数据 ==================

type Record struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	SessionID uint `gorm:"not null;index" json:"session_id"`
	RecordData
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Lap struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	SessionID uint `gorm:"not null;index" json:"session_id"`
	LapData
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Session struct {
	ID               uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID           *uint  `gorm:"index" json:"user_id,omitempty"`
	OriginalFilename string `gorm:"size:255;not null" json:"original_filename"`
	FileHash         string `gorm:"size:64;uniqueIndex" json:"file_hash"`
	SessionData
	// 关联（数据库层使用 DB 类型）
	Laps      []Lap     `gorm:"foreignKey:SessionID" json:"laps,omitempty"`
	Records   []Record  `gorm:"foreignKey:SessionID" json:"records,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (Session) TableName() string { return "sessions" }
func (Record) TableName() string  { return "records" }
func (Lap) TableName() string     { return "laps" }
