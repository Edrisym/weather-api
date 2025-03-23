package main

import (
	"log"
	"net/http"

	"weather-api/handlers"
)

func main() {
	// if os.Getenv("OPENWEATHER_API_KEY") == "" {
	// 	log.Fatal("OPENWEATHER_API_KEY environment variable is not set")
	// }

	http.HandleFunc("/weather", handlers.WeatherHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
