package main

import owm "github.com/briandowns/openweathermap"

type config struct {
	APIKey string `yaml:"api_key"`
}

type Weather struct {
	Current      owm.CurrentWeatherData
	Forecast5Day owm.Forecast5WeatherData
	Forecast     owm.ForecastWeatherData
	Coordinates  owm.Coordinates
}
