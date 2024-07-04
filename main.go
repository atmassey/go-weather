package main

import (
	"log"
	"net/http"
)

var APIKey string
var APIKeyValid bool

func main() {
	var err error

	s := &http.Server{
		Addr:    ":8081",
		Handler: nil,
	}

	http.HandleFunc("/", CurrentWeatherHandler)
	log.Println("Server starting on port 8081...")
	err = s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start. Error: %v", err)
	}
}
