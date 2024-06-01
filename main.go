package main

import (
	"fmt"
	"net/http"
)

var conf config

func main() {

	conf.ReadConfig()
	err := conf.CheckAPIKey()
	if err != nil {
		fmt.Errorf("Error: %v", err)
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
