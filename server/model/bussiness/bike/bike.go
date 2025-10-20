package bike

import "time"

type Bike struct {
	BikeName      string    `gorm:"column:bike_name;primaryKey" json:"bike_name"`
	ReleaseDate   time.Time `gorm:"column:release_date" json:"release_date"`
	MSRP          float64   `gorm:"column:msrp;not null" json:"msrp"`
	Manufacturer  string    `gorm:"column:manufacturer;not null" json:"manufacturer"`
	Mechanism     string    `gorm:"column:mechanism" json:"mechanism"`
	FrameMaterial string    `gorm:"column:frame_material" json:"frame_material"`
	BikeType      string    `gorm:"column:bike_type" json:"bike_type"`
}
