package domain

import (
	"time"
)

type Station struct {
	ID          string  `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Street      string  `json:"street"`
	Place       string  `json:"place"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	Dist        float64 `json:"dist"`
	HouseNumber string  `json:"houseNumber"`
	PostCode    int     `json:"postCode"`
}

type Call struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	StationID string    `json:"stationId" gorm:"index;not null"`
	Timestamp time.Time `json:"timestamp" gorm:"index"`

	Diesel float64 `json:"diesel"`
	E5     float64 `json:"e5"`
	E10    float64 `json:"e10"`
	IsOpen bool    `json:"isOpen"`
	Dist   float64 `json:"dist"`
}
