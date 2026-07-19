package api

import (
	"gas-thief/app/domain"
	"time"
)

func MapResponse(response APIResponse, timestamp time.Time) ([]domain.Station, []domain.Call) {
	count := len(response.Stations)
	stations := make([]domain.Station, count)
	calls := make([]domain.Call, count)

	for i, s := range response.Stations {

		station := domain.Station{
			ID:          s.ID,
			Name:        s.Name,
			Brand:       s.Brand,
			Street:      s.Street,
			Place:       s.Place,
			Lat:         s.Lat,
			Lng:         s.Lng,
			Dist:        s.Dist,
			HouseNumber: s.HouseNumber,
			PostCode:    s.PostCode,
		}
		stations[i] = station

		call := domain.Call{
			StationID: s.ID,
			Timestamp: timestamp,
			Diesel:    s.Diesel,
			E5:        s.E5,
			E10:       s.E10,
			IsOpen:    s.IsOpen,
			Dist:      s.Dist,
		}
		calls[i] = call
	}

	return stations, calls
}
