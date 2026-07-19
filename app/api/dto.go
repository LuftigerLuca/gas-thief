package api

import (
	"fmt"
	"net/url"
	"strconv"
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
	OK       bool      `json:"ok"`
	License  string    `json:"license"`
	Data     string    `json:"data"`
	Status   string    `json:"status"`
	Message  string    `json:"message,omitempty"`
	Stations []Station `json:"stations"`
}

type Station struct {
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
