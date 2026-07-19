package main

import (
	"fmt"
	"gas-thief/app/domain"
	"gas-thief/app/settings"
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func connectToDB(settings *settings.Settings) *gorm.DB {
	URI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", settings.DBUser, settings.DBPassword, settings.DBHost, settings.DBPort, settings.DBName)
	DB, err := gorm.Open(mysql.Open(URI), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect to database", "host", settings.DBHost, "port", settings.DBPort, "database", settings.DBName, "error", err)
	}

	if err = DB.AutoMigrate(&domain.Station{}, &domain.Call{}); err != nil {
		slog.Error("database migration failed", "database", settings.DBName, "error", err)
	}

	slog.Info("database connection established", "host", settings.DBHost, "database", settings.DBName)
	return DB
}

func saveStations(db *gorm.DB, stations []domain.Station) error {
	if len(stations) == 0 {
		slog.Debug("no stations to save")
		return nil
	}
	slog.Debug("upserting stations", "count", len(stations))
	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"name", "brand", "street", "house_number", "place", "post_code", "lat", "lng",
		}),
	}).Create(&stations).Error
}

func saveCalls(db *gorm.DB, calls []domain.Call) error {
	if len(calls) == 0 {
		slog.Debug("no calls to save")
		return nil
	}
	slog.Debug("inserting calls", "count", len(calls))
	return db.Create(&calls).Error
}
