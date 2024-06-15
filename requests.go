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
	w.CurrentByName("Bowling Green")

	return &Weather{
		Current: *w,
	}
}

func GetForecast5(location, units, lang string, api_key string) (*owm.Forecast5WeatherData, error) {
	w, err := owm.NewForecast("5", units, lang, api_key)
	if err != nil {
		return nil, err
	}
	w.DailyByName(location, 5)
	forecast := w.ForecastWeatherJson.(*owm.Forecast5WeatherData)
	return forecast, err
}
