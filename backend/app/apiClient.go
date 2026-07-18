package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func callApi(settings *Settings) (APIResponse, error) {
	request := ListRequest{
		Lat:    settings.LookUpLat,
		Lng:    settings.LookUpLng,
		Rad:    settings.LoopUpRadius,
		Type:   "all",
		ApiKey: settings.APIKey,
	}
	slog.Debug("calling gas station API", "url", request.URI())
	resp, err := http.Get(request.URI())
	if err != nil {
		return APIResponse{}, fmt.Errorf("API request failed: %w", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

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
