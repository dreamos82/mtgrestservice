package main

import (
	"fmt"
	"mtgrestservice/server"
	"mtgrestservice/config"
)

func main() {
	fmt.Println("Reading configuration")
	configuration, err := config.ReadConfigFile("config.properties")
	fmt.Println(configuration)
	fmt.Println(err)
	if err == nil && configuration != nil {
		fmt.Println("Launching server")
		server.InitServer()
	}
}
