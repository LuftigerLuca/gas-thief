package main

import (
	"fmt"
	"log/slog"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func connectToDB(settings *Settings) *gorm.DB {
	URI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", settings.DBUser, settings.DBPassword, settings.DBHost, settings.DBPort, settings.DBName)
	DB, err := gorm.Open(mysql.Open(URI), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect to database", "host", settings.DBHost, "port", settings.DBPort, "database", settings.DBName, "error", err)
	}

	if err = DB.AutoMigrate(&Station{}, &Call{}); err != nil {
		slog.Error("database migration failed", "database", settings.DBName, "error", err)
	}

	slog.Info("database connection established", "host", settings.DBHost, "database", settings.DBName)
	return DB
}

func getHealthStatus(db *gorm.DB) (*HealthStatus, error) {
	var stationCount int64
	if err := db.Model(&Station{}).Count(&stationCount).Error; err != nil {
		return nil, fmt.Errorf("failed to count stations: %w", err)
	}

	var lastCall Call
	err := db.Order("timestamp DESC").First(&lastCall).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("failed to query last call: %w", err)
	}

	return &HealthStatus{
		DBConnected:   true,
		LastCall:      lastCall.Timestamp,
		TotalStations: int(stationCount),
	}, nil
}

func saveApiResult(db *gorm.DB, stations []ResponseStation, time time.Time) error {
	slog.Info("saving gas station data", "count", len(stations))
	for _, s := range stations {

		station := Station{
			ID:          s.ID,
			Name:        s.Name,
			Brand:       s.Brand,
			Street:      s.Street,
			Place:       s.Place,
			Lat:         s.Lat,
			Lng:         s.Lat,
			Dist:        s.Dist,
			HouseNumber: s.HouseNumber,
			PostCode:    s.PostCode,
		}

		err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "brand", "street", "house_number", "place", "post_code", "lat", "lng"}),
		}).Create(&station).Error
		if err != nil {
			return err
		}

		call := Call{
			StationID: s.ID,
			Timestamp: time,
			Diesel:    s.Diesel,
			E5:        s.E5,
			E10:       s.E10,
			IsOpen:    s.IsOpen,
			Dist:      s.Dist,
		}
		err = db.Create(&call).Error
		if err != nil {
			return err
		}
	}
	return nil
}
