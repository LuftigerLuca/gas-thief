package main

import (
	"gas-thief/app/settings"
	"log/slog"
)

func main() {
	slog.Info("gas-thief starting up")
	settings := settings.LoadSettings()
	db := connectToDB(settings)
	slog.Info("starting scheduler", "interval_minutes", settings.LookUpInterval)
	go runScheduler(db, settings)
	select {}
}
