package main

import (
	"log"
	"net/http"
)

var APIKey string
var Units string
var City string

func main() {
	var err error

	s := &http.Server{
		Addr:    ":8081",
		Handler: nil,
		ReadHeaderTimeout: time.Second * 10,
	}

	http.HandleFunc("/", DataEntryHandler)
	http.HandleFunc("/current", CurrentWeatherHandler)
	http.HandleFunc("/forecast", WeatherForecastHandler)
	log.Println("Server starting on port 8081...")
	err = s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start. Error: %v", err)
	}
}
