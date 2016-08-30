package main

import (
	"fmt"
	"mtgrestservice/server"
)

func main() {
	/*xml_file, err := os.Open("assets/magicsymbols.xml")
	  if( err != nil ){
	    fmt.Println("Error opening file: ", err)
	    return
	  }

	  defer xml_file.Close()

	  b, _ := ioutil.ReadAll(xml_file)

	  var q datamodel.Query
		xml.Unmarshal(b, &q)

	  editionsMap = make(map[string]datamodel.Edition)

		for _, edition := range q.Editions {
	    		//fmt.Printf("\t%s - %s - %s\n", edition.Code, edition.Names[0].Name, edition.Names[0].Lang)
	        EditionsMap[edition.Code] = edition
	        result, err := json.Marshal(edition)
	        if ( err != nil) {
	          fmt.Println("Error: ", err)
	          return
	        }
	        fmt.Printf("%s\n", result)
		}
	  fmt.Println(EditionsMap["EMN"])*/
	/*EditionsMap = datamodel.LoadMap()
	  fmt.Println(EditionsMap["EMN"])*/
	fmt.Println("Launching server")
	server.InitServer()
}
