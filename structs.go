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
}

type ForecastDisplay struct {
	Date        string
	Temperature float64
	Condition   string
	Humidity    int
	WindSpeed   float64
}
