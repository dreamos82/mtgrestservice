package server

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"mtgrestservice/datamodel"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Root struct {
	EditionsArray []datamodel.Edition `xml:"edition"`
}

var EditionsMap map[string]datamodel.Edition
var RootElement Root

//var EditionsArray []datamodel.Edition

//InitServer will initialize the server
func InitServer() {
	EditionsMap, RootElement.EditionsArray = datamodel.LoadMap()
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/listeditions", listEditions)
	myRouter.HandleFunc("/listeditions/{format}", listEditions)
	myRouter.HandleFunc("/getedition/{key}", getEdition)
	myRouter.HandleFunc("/getedition/{key}/{format}", getEdition)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func getEdition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	displayFormat := vars["format"]
	editionCode := vars["key"]
	w.Header().Set("charset", "utf-8")
	if displayFormat != "" && strings.Compare(displayFormat, "xml") == 0 {
		w.Header().Set("Content-Type", "text/xml")
		xml.NewEncoder(w).Encode(EditionsMap[editionCode])
	} else if displayFormat == "" || strings.Compare(displayFormat, "json") == 0 {
		fmt.Println("else JSON")
		w.Header().Set("Content-Type", "text/json")
		if err := json.NewEncoder(w).Encode(EditionsMap[editionCode]); err != nil {
			panic(err)
		}
	}
	fmt.Println("GetEditions endpoint ", displayFormat)
}

func listEditions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	displayFormat := vars["format"]
	fmt.Println("ListEditions endpoint hit", displayFormat)
	w.Header().Set("charset", "utf8")
	if displayFormat != "" && strings.Compare(displayFormat, "xml") == 0 {
		w.Header().Set("Content-Type", "text/xml")
		if err := xml.NewEncoder(w).Encode(RootElement); err != nil {
			panic(err)
		}
	} else if displayFormat == "" || strings.Compare(displayFormat, "json") {
		fmt.Println("JSON list")
		w.Header().Set("Content-Type", "text/json")
		if err := json.NewEncoder(w).Encode(EditionsMap); err != nil {
			panic(err)
		}
	}
}
