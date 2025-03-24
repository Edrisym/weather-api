package services

import (
	"encoding/json"
	"fmt"

	"net/http"

	"io"
	"weather-api/utils"
)

func GetWeather(lat, lon string) (*map[string]interface{}, error) {
	config := utils.NewConfig()
	apiKey := config.OpenWeatherAPIKey
	apiKey = "3178dd4cd259f93d2fe619bc3ef5af9f"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric", lat, lon, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather API request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherResponse map[string]interface{}
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return nil, err
	}

	return &weatherResponse, nil
}
