package main

import (
	"fmt"

	owm "github.com/briandowns/openweathermap"
)

func UpdateWeather() *Weather {
	w, err := owm.NewCurrent("F", "EN", conf.APIKey)
	if err != nil {
		fmt.Println(err)
	}
	w.CurrentByName("Bowling Green, US")

	return &Weather{
		Current: *w,
	}

}
