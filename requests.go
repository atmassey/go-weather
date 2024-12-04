package main

import (
	"fmt"

	owm "github.com/briandowns/openweathermap"
)

func UpdateCurrentWeather() *Weather {
	w, err := owm.NewCurrent(Units, "EN", APIKey)
	if err != nil {
		fmt.Println(err)
	}
	err = w.CurrentByName(City)
	if err != nil {
		fmt.Println(err)
	}

	return &Weather{
		Current: *w,
	}
}

func GetForecast5(location, units, lang string, api_key string) (*owm.Forecast5WeatherData, error) {
	w, err := owm.NewForecast("5", units, lang, api_key)
	if err != nil {
		return nil, err
	}
	err = w.DailyByName(location, 5)
	if err != nil {
		return nil, err
	}
	forecast := w.ForecastWeatherJson.(*owm.Forecast5WeatherData)
	return forecast, err
}

func GetForecast5New(location, units, lang string, api_key string) (*ForecastWeather, error) {
	w, err := owm.NewForecast("5", units, lang, api_key)
	if err != nil {
		return nil, err
	}
	err = w.DailyByName(location, 5)
	if err != nil {
		return nil, err
	}
	forecast := ForecastWeather{}
	for _, data := range w.ForecastWeatherJson.(*owm.Forecast5WeatherData).List {
		forecast.Forecast = append(forecast.Forecast, ForecastWeatherData{
			Description: data.Weather[0].Description,
			TempMax:     data.Main.TempMax,
			TempMin:     data.Main.TempMin,
			Icon:        data.Weather[0].Icon,
			Condition:   data.Weather[0].Main,
			Date:        data.DtTxt.Time.Format("2006-01-02 15:04:05"),
		})
	}
	forecast.City = w.ForecastWeatherJson.(*owm.Forecast5WeatherData).City.Name
	fmt.Printf("Forecast: %v\n", forecast)
	return &forecast, err
}
