package main

import (
	"gas-thief/app/api"
	"gas-thief/app/settings"
	"log/slog"
	"time"

	"gorm.io/gorm"
)

func runScheduler(db *gorm.DB, settings *settings.Settings) {
	for {
		slog.Info("fetching gas stations", "lat", settings.LookUpLat, "lng", settings.LookUpLng, "radius", settings.LoopUpRadius)
		res, err := api.Call(settings)

		if err != nil {
			slog.Warn("failed to fetch gas stations from API", "error", err)
			continue
		}

		slog.Info("received gas station data", "stations", len(res.Stations))
		stations, calls := api.MapResponse(res, time.Now())
		slog.Info("saving data", "stations", len(stations), "calls", len(calls))
		if err := saveStations(db, stations); err != nil {
			slog.Error("failed to save stations", "count", len(stations), "error", err)
		}
		if err := saveCalls(db, calls); err != nil {
			slog.Error("failed to save calls", "count", len(calls), "error", err)
		}

		time.Sleep(time.Duration(settings.LookUpInterval) * time.Minute)
	}
}
