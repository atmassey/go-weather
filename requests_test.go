package main

import (
	"fmt"
	"os"
	"testing"

	owm "github.com/briandowns/openweathermap"
	"github.com/joho/godotenv"
)

func TestValidAPIKey(t *testing.T) {

	godotenv.Load()
	api_key := os.Getenv("API_KEY")

	err := owm.ValidAPIKey(api_key)
	if err != nil {
		t.Error("Error: Invalid Key")
	}
	t.Log("API key is valid")
}

func TestWeatherAPI(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		t.Error("Error loading .env file")
	}
	api_key := os.Getenv("API_KEY")
	w, err := owm.NewCurrent("F", "en", api_key) // fahrenheit (imperial) with Russian output
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	w.CurrentByName("Bowling Green")
	fmt.Println(w.Main.Temp)
	fmt.Println(w.Main.Humidity)

	if w.Main.Temp == 0 {
		t.Error("Temperature is 0")
	}

}

func TestForecastAPI(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		t.Error("Error loading .env file")
	}
	api_key := os.Getenv("API_KEY")
	forecast, err := GetForecast5("Bowling Green", "F", "en", api_key)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	fmt.Printf("Forecast: %v\n", forecast)

}

func TestUpdateCurrentWeather(t *testing.T) {
	weather_request := UpdateCurrentWeather()

	weatherData := WeatherDisplay{
		City:        weather_request.Current.Name,
		Temperature: weather_request.Current.Main.Temp,
		Humidity:    weather_request.Current.Main.Humidity,
		WindSpeed:   weather_request.Current.Wind.Speed,
	}
	t.Logf("Current City: %v", weatherData.City)
	t.Logf("Current Temperature: %v", weatherData.Temperature)
	t.Logf("Current Humidity: %v", weatherData.Humidity)
	t.Logf("Current Wind Speed: %v", weatherData.WindSpeed)

}
