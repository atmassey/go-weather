package main

import (
	"fmt"
)

func main() {
	var conf config

	conf.ReadConfig()
	fmt.Println(conf.APIKey)
}
