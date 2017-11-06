package main

import (
	"fmt"
	"mtgrestservice/server"
	"mtgrestservice/config"
)

func main() {
	fmt.Println("Reading configuration")
	var configuration map[string]string
	configuration, err := config.ReadConfigFile("config.properties")
	if err != nil {
		fmt.Println("Error opening config file: " + err.Error())
		fmt.Println("Use default configuration")
		configuration = config.CreateDefaultConfigurationMap()
	}
	if configuration != nil {
		fmt.Println("Launching server")
		server.InitServer(configuration)
	}
}
