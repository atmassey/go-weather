package main

type config struct {
	APIKey string `yaml:"api_key"`
}

type weather struct {
	temp       float64 `json:"temp"`
	feels_like float64 `json:"feels_like"`
	humidity   int     `json:"humidity"`
}
