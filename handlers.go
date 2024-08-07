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
		log.Printf("Invalid API Key. Error: %v", err)
		return err
	}
	return nil
}

func CurrentWeatherHandler(w http.ResponseWriter, r *http.Request) {

	weather_request := UpdateCurrentWeather()
	// Check if weather data is found
	if len(weather_request.Current.Weather) == 0 {
		NoDataHandler(w, r)
		return
	}
	weatherData := WeatherDisplay{
		City:        weather_request.Current.Name,
		Temperature: weather_request.Current.Main.Temp,
		Condition:   weather_request.Current.Weather[0].Description,
		Humidity:    weather_request.Current.Main.Humidity,
		WindSpeed:   weather_request.Current.Wind.Speed,
		Icon:        weather_request.Current.Weather[0].Icon,
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

func WeatherForecastHandler(w http.ResponseWriter, r *http.Request) {
	forecast, err := GetForecast5New(City, Units, "EN", APIKey)
	if err != nil {
		http.Error(w, "Failed to retrieve forecast", http.StatusInternalServerError)
		return
	}

	if forecast == nil || len(forecast.City) == 0 {
		NoDataHandler(w, r)
		return
	}

	tmpl, err := template.ParseFS(os.DirFS("./web"), "forecast_template.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}
	log.Printf("Forecast: %v", forecast)
	if err := tmpl.Execute(w, forecast); err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func DataEntryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		APIKey = r.FormValue("apikey")
		log.Printf("API Key received: %s", APIKey)
		City = r.FormValue("city")
		Units = r.FormValue("units")

		http.Redirect(w, r, "/current", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFS(os.DirFS("./web"), "data_entry.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NoDataHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(os.DirFS("./web"), "no_data.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
