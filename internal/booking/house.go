package booking

import "fmt"

type HouseType string

const (
	HouseTypeStandart HouseType = "standard"
	HouseTypeBig      HouseType = "big"
)

type House struct {
	ID       int       `json:"id"`
	Type     HouseType `json:"type"`
	Capacity int       `json:"capacity"`
}

func ParseHouseType(value string) (HouseType, error) {
	switch HouseType(value) {
	case HouseTypeStandart, HouseTypeBig:
		return HouseType(value), nil
	default:
		return "", fmt.Errorf("unknown house type: %q", value)
	}
}
