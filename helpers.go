package main

import (
	"fmt"
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
	"gopkg.in/yaml.v2"
)

func (c *config) ReadConfig() *config {
	yamlFile, err := os.ReadFile("prod/config.yml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err)
	}

	return c
}

func (conf *config) CheckAPIKey() error {

	err := owm.ValidAPIKey(conf.APIKey)
	if err != nil {
		log.Fatalf("Invalid API Key. Error: %v", err)
		return err
	}
	return nil
}
