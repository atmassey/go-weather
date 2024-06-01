package main

import (
	"log"

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
