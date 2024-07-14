package main

import owm "github.com/briandowns/openweathermap"

type Weather struct {
	Current     owm.CurrentWeatherData
	Forecast    owm.ForecastWeatherData
	Coordinates owm.Coordinates
}

type WeatherDisplay struct {
	City        string
	Temperature float64
	Condition   string
	Humidity    int
	WindSpeed   float64
	Icon        string
}

type ForecastWeather struct {
	City     string
	Forecast []ForecastWeatherData
}

type ForecastWeatherData struct {
	Description string
	TempMax     float64
	TempMin     float64
	Icon        string
	Condition   string
	Date        string
}
