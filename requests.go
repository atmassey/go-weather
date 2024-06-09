package main

import (
	"fmt"

	owm "github.com/briandowns/openweathermap"
)

func UpdateCurrentWeather() *Weather {
	w, err := owm.NewCurrent("F", "EN", APIKey)
	if err != nil {
		fmt.Println(err)
	}
	w.CurrentByName("Bowling Green, KY")

	return &Weather{
		Current: *w,
	}
}

func UpdateForecast() *Weather {
	f, err := owm.NewForecast("5", "F", "EN", APIKey)
	if err != nil {
		fmt.Println(err)
	}
	f.DailyByName("Bowling Green, KY", 5)

	return &Weather{
		Forecast: *f,
	}
}
