package main

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type ListRequest struct {
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Rad    float64 `json:"rad"`
	Type   string  `json:"type"`
	Sort   string  `json:"sort"`
	ApiKey string  `json:"apikey"`
}

type APIResponse struct {
	OK       bool              `json:"ok"`
	License  string            `json:"license"`
	Data     string            `json:"data"`
	Status   string            `json:"status"`
	Message  string            `json:"message,omitempty"`
	Stations []ResponseStation `json:"stations"`
}

type ResponseStation struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Street      string  `json:"street"`
	Place       string  `json:"place"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	Dist        float64 `json:"dist"`
	Diesel      float64 `json:"diesel"`
	E5          float64 `json:"e5"`
	E10         float64 `json:"e10"`
	IsOpen      bool    `json:"isOpen"`
	HouseNumber string  `json:"houseNumber"`
	PostCode    int     `json:"postCode"`
}

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

type HealthStatus struct {
	DBConnected   bool      `json:"db_connected"`
	LastCall      time.Time `json:"last_call"`
	TotalStations int       `json:"total_stations"`
}

func (r ListRequest) URI() string {
	params := url.Values{}
	params.Set("lat", strconv.FormatFloat(r.Lat, 'f', -1, 64))
	params.Set("lng", strconv.FormatFloat(r.Lng, 'f', -1, 64))
	params.Set("rad", strconv.FormatFloat(r.Rad, 'f', -1, 64))
	params.Set("type", r.Type)
	params.Set("sort", r.Sort)
	params.Set("apikey", r.ApiKey)
	q := params.Encode()

	return fmt.Sprintf("https://creativecommons.tankerkoenig.de/json/list.php?%s", q)
}
