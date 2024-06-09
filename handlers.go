package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	owm "github.com/briandowns/openweathermap"
)

func CheckAPIKey(apiKey string) error {

	err := owm.ValidAPIKey(apiKey)
	if err != nil {
		log.Fatalf("Invalid API Key. Error: %v", err)
		return err
	}
	return nil
}

func CurrentWeatherHandler(w http.ResponseWriter, r *http.Request) {

	weather_request := UpdateCurrentWeather()
	// Check if weather data is found
	if len(weather_request.Current.Weather) == 0 {
		http.Error(w, "Error: No weather data found", http.StatusInternalServerError)
		return
	}

	weatherData := WeatherDisplay{
		City:        weather_request.Current.Name,
		Temperature: weather_request.Current.Main.Temp,
		Condition:   weather_request.Current.Weather[0].Description,
		Humidity:    weather_request.Current.Main.Humidity,
		WindSpeed:   weather_request.Current.Wind.Speed,
	}
	log.Printf("Current Weather: %v", weatherData)

	tmpl, err := template.ParseFS(os.DirFS("./web"), "current_template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, weatherData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
