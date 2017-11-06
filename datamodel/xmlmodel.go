package datamodel

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//Query Struct used to perform Queries on Editions Map
type Query struct {
	Name     string    `xml:"root" json:"root"`
	Editions []Edition `xml:"edition" json:"edition"`
}

//Edition basic magic edition object.
type Edition struct {
	Launch         string             `xml:"launch" json:"launch"`
	Code           string             `xml:"code" json:"code"`
	Names          []NameNode         `xml:"names>name"`
	Cards          CardsComposition   `xml:"cards" json:"cards,omitempty"`
	Preconstructed PreconstructedInfo `xml:"preconstructed,omitempty" json:"preconstructed,omitempty"`
	Vault          *struct{}          `xml:"vault" json:"vault"`
	Online         *struct{}          `xml:"online" json:"online"`
}

//NameNode it contains the edition name Name in language Lang
type NameNode struct {
	Name string `xml:",chardata" json:",chardata"`
	Lang string `xml:"lang,attr" json:"lang"`
}

//CardsComposition it contain the cards composition for an edition (number of common, uncommon, etc..)
type CardsComposition struct {
	Lands    int `xml:"lands,omitempty" json:"lands,omitempty"`
	Common   int `xml:"common,omitempty" json:"common,omitempty"`
	Uncommon int `xml:"uncommon,omitempty" json:"uncommon,omitempty"`
	Rare     int `xml:"rare,omitempty" json:"rare,omitempty"`
	Mythic   int `xml:"mythic,omitempty" json:"mythic,omitempty"`
	Oversize int `xml:"oversize,omitempty" json:"oversize,omitempty"`
	Other    int `xml:"other,omitempty" json:"other,omitempty"`
}

//PreconstructedInfo it contain information about Preconstructed decks (number of decks, type, etc.)
type PreconstructedInfo struct {
	Type  string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Decks int    `xml:"decks,attr,omitempty" json:"decks,omitempty"`
	Size  int    `xml:"size,omitempty" json:"size,omitempty"`
}

//LoadMap it loads the edition map into memory
func LoadMap(fullpath string) (map[string]Edition, []Edition) {
	xmlFile, err := os.Open(fullpath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, nil
	}

	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

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
