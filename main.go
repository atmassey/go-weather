package main

import (
	"log"
	"net/http"
	"os"
)

var APIKey string

func main() {

	APIKey = os.Getenv("API_KEY")
	if APIKey == "" {
		log.Panic("API_KEY environment variable not set")
	}

	s := &http.Server{
		Addr:    ":8081",
		Handler: nil,
	}

	http.HandleFunc("/", WeatherHandler)
	log.Println("Server starting on port 8081...")
	s.ListenAndServe()
}
