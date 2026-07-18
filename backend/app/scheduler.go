package main

import (
	"log/slog"
	"time"

	"gorm.io/gorm"
)

func run(db *gorm.DB, settings *Settings) {
	for {
		slog.Info("fetching gas stations", "lat", settings.LookUpLat, "lng", settings.LookUpLng, "radius", settings.LoopUpRadius)
		res, err := callApi(settings)
		if err != nil {
			slog.Warn("failed to fetch gas stations from API", "error", err)
		} else {
			slog.Info("received gas station data", "stations", len(res.Stations))
			if err := saveApiResult(db, res.Stations, time.Now()); err != nil {
				slog.Warn("failed to persist gas station data", "stations", len(res.Stations), "error", err)
			}
		}

		time.Sleep(time.Duration(settings.LookUpInterval) * time.Minute)
	}
}
