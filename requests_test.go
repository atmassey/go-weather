package main

import (
	"fmt"
	"os"
	"testing"

	owm "github.com/briandowns/openweathermap"
	"github.com/joho/godotenv"
)

func TestValidAPIKey(t *testing.T) {
	var api_key string
	godotenv.Load()
	api_key = os.Getenv("API_KEY")

	err := owm.ValidAPIKey(api_key)
	if err != nil {
		t.Error("Error: Invalid Key")
	}
	t.Log("API key is valid")
}

func TestWeatherAPI(t *testing.T) {
	var api_key string
	err := godotenv.Load()
	if err != nil {
		t.Error("Error loading .env file")
	}
	api_key = os.Getenv("API_KEY")
	w, err := owm.NewCurrent("F", "en", api_key) // fahrenheit (imperial) with Russian output
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	w.CurrentByName("Bowling Green")
	fmt.Println(w.Main.Temp)
	fmt.Println(w.Main.Humidity)
}
