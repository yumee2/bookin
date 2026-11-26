package booking

import "fmt"

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

func ParseHouseType(value string) (HouseType, error) {
	switch HouseType(value) {
	case HouseTypeStandart, HouseTypeBig:
		return HouseType(value), nil
	default:
		return "", fmt.Errorf("unknown house type: %q", value)
	}
}
