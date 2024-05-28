package main

import (
	"fmt"
	"testing"

	owm "github.com/briandowns/openweathermap"
)

func TestValidAPIKey(t *testing.T) {
	var conf config
	var api_key string
	config := conf.ReadConfig()
	api_key = config.APIKey

	err := owm.ValidAPIKey(api_key)
	if err != nil {
		t.Error("Error: Invalid Key")
	}
	t.Log("API key is valid")
}

func TestWeatherAPI(t *testing.T) {
	var conf config
	config := conf.ReadConfig()
	w, err := owm.NewCurrent("F", "en", config.APIKey) // fahrenheit (imperial) with Russian output
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	w.CurrentByName("Bowling Green")
	fmt.Println(w.Main.Temp)
	fmt.Println(w.Main.Humidity)
}
