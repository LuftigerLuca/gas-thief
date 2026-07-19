package api

import (
	"encoding/json"
	"fmt"
	"gas-thief/app/settings"
	"log/slog"
	"net/http"
	"time"
)

func Call(settings *settings.Settings) (APIResponse, error) {
	request := ListRequest{
		Lat:    settings.LookUpLat,
		Lng:    settings.LookUpLng,
		Rad:    settings.LoopUpRadius,
		Type:   "all",
		ApiKey: settings.APIKey,
	}
	slog.Debug("calling gas station API", "url", request.URI())
	start := time.Now()
	resp, err := http.Get(request.URI())
	duration := time.Since(start)

	if err != nil {
		slog.Warn("API request failed", "duration_ms", duration.Milliseconds(), "error", err)
		return APIResponse{}, fmt.Errorf("API request failed: %w", err)
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	slog.Info("API response received", "status", resp.StatusCode, "duration_ms", duration.Milliseconds())

	if resp.StatusCode != http.StatusOK {
		return APIResponse{}, fmt.Errorf("unexpected HTTP status %s", resp.Status)
	}

	var pResp APIResponse
	err = json.NewDecoder(resp.Body).Decode(&pResp)
	if err != nil {
		return APIResponse{}, fmt.Errorf("failed to decode API response: %w", err)
	}

	return pResp, nil
}
