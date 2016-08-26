package datamodel

import (
  "fmt"
  "encoding/xml"
  "os"
  "io/ioutil"
)

type Query struct {
  Name string `xml:"root" json:"root"`
  Editions []Edition `xml:"edition" json:"edition"`
}

type Edition struct {
  Launch  string  `xml:"launch" json:"launch"`
  Code  string  `xml:"code" json:"code"`
  Names []NameNode `xml:"names>name"`
  Cards CardsComposition `xml:"cards" json:"cards"`
  Preconstructed PreconstructedInfo `xml:"preconstructed" json:"preconstructed"`
  Vault *struct{} `xml:"vault" json:"vault"`
  Online *struct{} `xml:"online" json:"online"`
}

type NameNode struct {
  Name string `xml:",chardata" json:",chardata"`
  Lang string `xml:"lang,attr" json:"lang"`
}

type CardsComposition struct {
  Lands int `xml:"lands" json:"lands"`
  Common int `xml:"common" json:"common"`
  Uncommon int `xml:"uncommon" json:"uncommon"`
  Rare int `xml:"rare" json:"rare"`
  Mythic int `xml:"mythic" json:"mythic"`
  Oversize int `xml:"oversize" json:"oversize"`
  Other int `xml:"other" json:"other"`
}

type PreconstructedInfo struct {
  Type string `xml:"type,attr" json:"type"`
  Decks string `xml:"decks,attr" json:"decks"`
  Size int `xml:"size" json:"size"`
}

func LoadMap() (map[string]Edition, []Edition) {
  xml_file, err := os.Open("assets/magicsymbols.xml")
  if( err != nil ){
    fmt.Println("Error opening file: ", err)
    return nil, nil
  }

  defer xml_file.Close()

  b, _ := ioutil.ReadAll(xml_file)

  var q Query
  xml.Unmarshal(b, &q)

  var editionsMap map[string]Edition

  editionsMap = make(map[string]Edition)

  for _, edition := range q.Editions {
        //fmt.Printf("\t%s - %s - %s\n", edition.Code, edition.Names[0].Name, edition.Names[0].Lang)
        editionsMap[edition.Code] = edition
  }
  return editionsMap, q.Editions
}

func (u Edition) MarshalJSON() ([]byte, error) {
  fmt.Println("CALLED")
  vaultValue bool
  vaultValue = true
  onlineValue bool
  vaultValue = false
  if(u.Vault == nil) {
    vaultValue = false
  }
  if(u.Online == nil) {
    onlineValue = false
  }
  type AliasEdition Edition
	return json.Marshal(struct {
    AliasEdition
    Vault bool `json:"vault"`
    Online bool `json:"online"`
	}{
    AliasEdition:   AliasEdition(u),
		Vault:       vaultValue,
		Online:      onlineValue,
	})
}
