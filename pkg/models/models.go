package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Cars struct {
	CarID     uuid.UUID       `json:"carId"`
	Mark      string          `json:"mark"`
	Model     string          `json:"model"`
	Kilometre uint64          `json:"kilometre"`
	Price     uint64          `json:"price"`
	Images    json.RawMessage `json:"images"`
}
type CarInfo struct {
	CarID         uuid.UUID       `json:"carId"`
	SellerID      uuid.UUID       `json:"sellerID"`
	Mark          *string         `json:"mark"`
	Model         *string         `json:"model"`
	BuildYear     int16           `json:"buildYear"`
	Power         int32           `json:"power"`
	TypePower     string          `json:"typePower"`
	Owners        *int32          `json:"owners"`
	Condition     *string         `json:"condition"`
	SteeringWheel *string         `json:"steeringWheel"`
	Color         *string         `json:"color"`
	Descriptions  string          `json:"descriptions"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	Release       bool            `json:"release"`
	Images        json.RawMessage `json:"images"`
	VIN           string          `json:"vin"`
	Kilometre     *uint64         `json:"kilometre"`
	Price         *uint64         `json:"price"`
	IsUserCar     *bool           `json:"isUserCar"`
	IsFavorite    *bool           `json:"isFavorite"`
}
