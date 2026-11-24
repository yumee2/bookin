package booking

type HouseType string

const (
	HouseTypeStandart HouseType = "standart"
	HouseTypeBig      HouseType = "big"
)

type House struct {
	ID       uint
	Type     HouseType
	Capacity int
}
