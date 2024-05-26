package main

import (
	"fmt"
)

func main() {
	var c config

	c.getAPIKey()
	fmt.Println(c.APIKey)
}
