package handlers

import (
	"encoding/json"
	"net/http"

	"weather-api/services"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")
	if lat == "" || lon == "" {
		http.Error(w, "lat and lon parameters are required", http.StatusBadRequest)
		return
	}

	weather, err := services.GetWeather(lat, lon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
