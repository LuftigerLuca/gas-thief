package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Settings struct {
	APIKey         string  `json:"api_key"`
	LookUpInterval uint    `json:"look_up_interval"`
	LookUpLat      float64 `json:"look_up_lat"`
	LookUpLng      float64 `json:"look_up_lng"`
	LoopUpRadius   float64 `json:"loop_up_radius"`
	WebPort        string  `json:"web_port"`
	DBHost         string  `json:"db_host"`
	DBPort         string  `json:"db_port"`
	DBUser         string  `json:"db_user"`
	DBPassword     string  `json:"db_password"`
	DBName         string  `json:"db_name"`
}

func LoadSettings() *Settings {
	if err := loadEnv(".env"); err != nil {
		slog.Warn("could not load .env file, using environment variables only", "error", err)
	}

	s := Settings{
		APIKey:         getString("API_KEY", "", true),
		LookUpInterval: getUInt("LOOK_UP_INTERVAL", 60, false),
		LookUpLat:      getFloat64("LOOK_UP_LAT", 0, true),
		LookUpLng:      getFloat64("LOOK_UP_LNG", 0, true),
		LoopUpRadius:   getFloat64("LOOK_UP_RADIUS", 10, false),
		WebPort:        getString("WEB_PORT", "8080", false),
		DBHost:         getString("DB_HOST", "", true),
		DBPort:         getString("DB_PORT", "", true),
		DBUser:         getString("DB_USER", "", true),
		DBPassword:     getString("DB_PASSWORD", "", true),
		DBName:         getString("DB_NAME", "", true),
	}

	slog.Info("settings loaded", "interval_min", s.LookUpInterval, "lat", s.LookUpLat, "lng", s.LookUpLng, "radius", s.LoopUpRadius)
	return &s
}

func getString(key string, def string, required bool) string {
	v, ok := os.LookupEnv(key)

	if !ok || v == "" {
		if required {
			slog.Error("missing required env value", "key", key)
			return ""
		}
		slog.Info("missing env value, taking default instead", "key", key, "default", def)
		return def
	}

	return v
}

func getFloat64(key string, def float64, required bool) float64 {
	v, ok := os.LookupEnv(key)

	if !ok || v == "" {
		if required {
			slog.Error("missing required env value", "key", key)
			return 0
		}
		slog.Info("missing env value, taking default instead", "key", key, "default", def)
		return def
	}

	vf, err := strconv.ParseFloat(v, 64)
	if err != nil {
		if required {
			slog.Error("cannot convert required env value", "key", key, "value", v, "error", err)
			return 0
		}
		slog.Warn("cannot convert env value, taking default instead", "key", key, "value", v, "error", err)
		return def
	}

	return vf
}

func getUInt(key string, def uint, required bool) uint {
	v, ok := os.LookupEnv(key)

	if !ok || v == "" {
		if required {
			slog.Error("missing required env value", "key", key)
			return 0
		}
		slog.Info("missing env value, taking default instead", "key", key, "default", def)
		return def
	}

	vi, err := strconv.Atoi(v)
	if err != nil {
		if required {
			slog.Error("cannot convert required env value", "key", key, "value", v, "error", err)
			return 0
		}
		slog.Warn("cannot convert env value, taking default instead", "key", key, "value", v, "error", err)
		return def
	}

	if vi < 0 {
		if required {
			slog.Error("required env value must not be negative", "key", key, "value", vi)
			return 0
		}
		slog.Warn("env value is negative, taking default instead", "key", key, "value", vi, "default", def)
		return def
	}

	return uint(vi)
}

func loadEnv(filename string) (retErr error) {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && retErr == nil {
			retErr = cerr
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.Trim(
			strings.TrimSpace(parts[1]), `"'`)

		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return scanner.Err()
}
