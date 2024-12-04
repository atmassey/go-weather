# Go-Weather

[![Build](https://github.com/atmassey/go-weather/actions/workflows/build.yml/badge.svg)](https://github.com/atmassey/go-weather/actions/workflows/build.yml)
[![Docker Image CI](https://github.com/atmassey/go-weather/actions/workflows/docker-image.yml/badge.svg)](https://github.com/atmassey/go-weather/actions/workflows/docker-image.yml)
[![Lint](https://github.com/atmassey/go-weather/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/atmassey/go-weather/actions/workflows/golangci-lint.yml)

Go-Weather is a very simple go web app that displays weather metrics.

The data entry page is a form that requires an OpenWeatherMap API key, a City, and Units.
![Data Entry](https://github.com/atmassey/go-weather/blob/main/docs/data_entry.png?raw=true)

The Current Weather page displays the current weather for the city entered on the Data Entry page.
![Current Weather](https://github.com/atmassey/go-weather/blob/main/docs/current_weather.png?raw=true)

The Forecast page displays the 5 day forecast for the city entered on the Data Entry page.
![Forecast](https://github.com/atmassey/go-weather/blob/main/docs/forecast_weather.png?raw=true)
