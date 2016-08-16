package server

import (
  "fmt"
  "log"
  "strings"
  "net/http"
  "encoding/json"
  "encoding/xml"
  "mymtgapi/datamodel"
  "github.com/gorilla/mux"
)

type Root struct {
  EditionsArray []datamodel.Edition `xml:"edition"`
}

var EditionsMap map[string]datamodel.Edition
var RootElement Root
//var EditionsArray []datamodel.Edition

func InitServer() {
  EditionsMap, RootElement.EditionsArray = datamodel.LoadMap();
  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/listeditions", listEditions )
  myRouter.HandleFunc("/listeditions/{format}", listEditions )
  myRouter.HandleFunc("/getedition/{key}", getEdition )
  myRouter.HandleFunc("/getedition/{key}/{format}", getEdition )
  log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func getEdition(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  displayFormat := vars["format"]
  editionCode := vars["key"]
  w.Header().Set("charset", "utf-8")
  if (displayFormat != "") {
    if (strings.Compare(displayFormat, "xml") == 0){
        w.Header().Set("Content-Type", "text/xml")
      xml.NewEncoder(w).Encode(EditionsMap[editionCode])
    }
  } else {
    w.Header().Set("Content-Type", "text/json")
    json.NewEncoder(w).Encode(EditionsMap[editionCode])
  }
  fmt.Println("GetEditions endpoint ", displayFormat)
}

func listEditions(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  displayFormat := vars["format"]
  fmt.Println("ListEditions endpoint", displayFormat)
  w.Header().Set("charset", "utf8")
  if (displayFormat != "" ) {
    if (strings.Compare(displayFormat, "xml") == 0){
      w.Header().Set("Content-Type", "text/xml")
      err := xml.NewEncoder(w).Encode(RootElement)
      if(err != nil){
        fmt.Println("Error occurred", err)
      }
    } else {
      w.Header().Set("Content-Type", "text/json")
      json.NewEncoder(w).Encode(EditionsMap)
    }
  } else {
    w.Header().Set("Content-Type", "text/json")
    json.NewEncoder(w).Encode(EditionsMap)
  }
}
