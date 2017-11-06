package config

import (
  "os"
  "fmt"
)

//Handle application configuration
type Configuration struct {
  Port  string
  MaxConnections  string
  AssetsFolder  string
  HostName  string
}

func ReadConfigFile(fileName string){
  configFile, err := os.Open(fileName)
  fmt.Println("Config")
  if err != nil {
    fmt.Println(err)
    return
  }
  if configFile != nil {
    fmt.Println("Open")
  }
}
