package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func (c *config) getAPIKey() *config {
	yamlFile, err := ioutil.ReadFile("prod/config.yml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err)
	}

	return c
}
