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
}

func TestForecastAPI(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		t.Error("Error loading .env file")
	}
	api_key := os.Getenv("API_KEY")
	f, err := owm.NewForecast("5", "F", "en", api_key)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	err = f.DailyByName("Bowling Green, KY", 5)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
