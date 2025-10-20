package derailleur

import (
	"time"
)

type Derailleur struct {
	DerailleurName string    `gorm:"column:derailleur_name;primaryKey" json:"derailleur_name"`
	ReleaseDate    time.Time `gorm:"column:release_date" json:"release_date"`
	MSRP           float64   `gorm:"column:msrp;not null" json:"msrp"`
	Manufacturer   string    `gorm:"column:manufacturer;not null" json:"manufacturer"`
	Speed          float64   `gorm:"column:speed;not null" json:"speed"`
	Mechanism      string    `gorm:"column:mechanism" json:"mechanism"`
}

// TableName 指定表名
func (Derailleur) TableName() string {
	return "derailleur"
}
