package main

import (
	"fmt"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html><h1>The temperature is %v degrees F</h1></html>", UpdateWeather().Current.Main.Temp)
	})

	s.ListenAndServe()
}
