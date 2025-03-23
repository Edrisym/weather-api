package utils

import (
	"os"
)

type Config struct {
	OpenWeatherAPIKey string
}

func NewConfig() *Config {
	return &Config{
		OpenWeatherAPIKey: os.Getenv("OPENWEATHER_API_KEY"),
	}
}
