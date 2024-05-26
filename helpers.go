package main

import (
	"fmt"
	"os"

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
