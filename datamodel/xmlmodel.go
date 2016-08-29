package datamodel

import (
  "fmt"
  "strings"
  "encoding/xml"
  "encoding/json"
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
  Cards CardsComposition `xml:"cards" json:"cards,omitempty"`
  Preconstructed PreconstructedInfo `xml:"preconstructed" json:"preconstructed,omitempty"`
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
  Type string `xml:"type,attr" json:"type,omitempty"`
  Decks string `xml:"decks,attr" json:"decks,omitempty"`
  Size int `xml:"size" json:"size,omitempty"`
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
  var vaultValue bool
  vaultValue = true
  var onlineValue bool
  vaultValue = false
  var preconstructedItem *PreconstructedInfo
  preconstructedItem = &u.Preconstructed
  var cardsItem *CardsComposition
  cardsItem = &u.Cards
  if(u.Vault == nil) {
    vaultValue = false
  }

  if(u.Online == nil) {
    onlineValue = false
  }

  if( PreconstructedInfo{} == u.Preconstructed ) {
    fmt.Println("PreconstructedEmpty")
    preconstructedItem = nil
  }

  if( CardsComposition{} == u.Cards ) {
    cardsItem = nil
  }
  fmt.Println(PreconstructedInfo{} == u.Preconstructed)
  fmt.Println("Cards ", CardsComposition{} == u.Cards)
  type AliasEdition Edition
	return json.Marshal(struct {
    AliasEdition
    Cards *CardsComposition `json:"cards,omitempty"`
    Preconstructed *PreconstructedInfo `json:"preconstructed,omitempty"`
    Vault bool `json:"vault"`
    Online bool `json:"online"`
	}{
    AliasEdition:   AliasEdition(u),
    Cards:  cardsItem,
    Preconstructed: preconstructedItem,
		Vault:       vaultValue,
		Online:      onlineValue,
	})
}

func (nameNode NameNode) MarshalJSON() ([]byte, error) {
  language := nameNode.Lang
  if( strings.Compare( nameNode.Lang, "") == 0){
    language = "en"
  }
  return json.Marshal(struct{
    Name string `json:"name"`
    Lang string `json:"lang"`
  }{
    Name: nameNode.Name,
    Lang: language,
  })
}
