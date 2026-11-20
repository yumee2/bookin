package postgres

type HouseType string

const (
	HouseTypeStandart HouseType = "standart"
	HouseTypeBig      HouseType = "big"
)

type House struct {
	ID       uint      `gorm:"primaryKey"`
	Type     HouseType `gorm:"type:text;not null"`
	Capacity int       `gorm:"not null"`
}
