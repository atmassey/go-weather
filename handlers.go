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

	weatherData := WeatherDisplay{
		City:        weather_request.Current.Name,
		Temperature: weather_request.Current.Main.Temp,
		Condition:   weather_request.Current.Weather[0].Description,
		Humidity:    weather_request.Current.Main.Humidity,
		WindSpeed:   weather_request.Current.Wind.Speed,
	}

	tmpl, err := template.ParseFS(os.DirFS("./web"), "template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, weatherData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
