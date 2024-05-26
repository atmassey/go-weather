package main

import (
	"fmt"
)

func main() {
	var conf config

	conf.getAPIKey()
	fmt.Println(conf.APIKey)
}
