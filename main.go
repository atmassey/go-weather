package main

import (
	"log"
	"net/http"
	"os"
)

var APIKey string

func main() {
	var err error
	APIKey = os.Getenv("API_KEY")
	if APIKey == "" {
		log.Panic("API_KEY environment variable not set")
	}

	err = CheckAPIKey(APIKey)
	if err != nil {
		log.Fatalf("Invalid API Key. Error: %v", err)
	}

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
