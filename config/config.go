package config

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "errors"
)

func ReadConfigFile(fileName string) (map[string]string, error){
  configFile, err := os.Open(fileName)
  if err != nil {
    fmt.Println(err)
    return nil, err
  } else if configFile != nil {
    scanner := bufio.NewScanner(configFile)
    //configuration := NewConfiguration()
    var configurationmap map[string]string
    configurationmap = make(map[string]string)
    for scanner.Scan() {             // internally, it advances token based on sperator
        property := strings.Split(scanner.Text(), "=")
        SetConfigurationProperty(&configurationmap, property)
    }
    return configurationmap, nil
  } else {
    err = errors.New("emit macho dwarf: elf header corrupted")
    return nil, err
  }
}

func CreateDefaultConfigurationMap() map[string]string {
  var configurationmap map[string]string
  configurationmap = make(map[string]string)
  configurationmap["port"]="8080"
  configurationmap["assetsfolder"] = "./assets"
  configurationmap["hostname"] = "localhost"
  configurationmap["maxconnections"] = "0"
  return configurationmap
}

func SetConfigurationProperty(configuration *map[string]string, property []string){
  if(len(property) == 2){
    (*configuration)[property[0]] = property[1]
  }
}
