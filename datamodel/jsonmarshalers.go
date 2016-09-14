package datamodel

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MarshalJSON generate JSON output for Edition
func (edition Edition) MarshalJSON() ([]byte, error) {
	var vaultValue bool
	var onlineValue bool
	vaultValue = false
	onlineValue = false
	var preconstructedItem *PreconstructedInfo
	preconstructedItem = &edition.Preconstructed
	var cardsItem *CardsComposition
	cardsItem = &edition.Cards
	if edition.Vault == nil {
		vaultValue = false
	}

	if edition.Online == nil {
		onlineValue = false
	}

	if (PreconstructedInfo{} == edition.Preconstructed) {
		fmt.Println("PreconstructedEmpty")
		preconstructedItem = nil
	}

	if (CardsComposition{} == edition.Cards) {
		cardsItem = nil
	}
	fmt.Println(PreconstructedInfo{} == edition.Preconstructed)
	fmt.Println("Cards ", CardsComposition{} == edition.Cards)
	type AliasEdition Edition
	return json.Marshal(struct {
		AliasEdition
		Cards          *CardsComposition   `json:"cards,omitempty"`
		Preconstructed *PreconstructedInfo `json:"preconstructed,omitempty"`
		Vault          bool                `json:"vault"`
		Online         bool                `json:"online"`
	}{
		AliasEdition:   AliasEdition(edition),
		Cards:          cardsItem,
		Preconstructed: preconstructedItem,
		Vault:          vaultValue,
		Online:         onlineValue,
	})
}

//MarshalJSON return json for nameNode object
func (nameNode NameNode) MarshalJSON() ([]byte, error) {
	language := nameNode.Lang
	if strings.Compare(nameNode.Lang, "") == 0 {
		language = "en"
	}
	return json.Marshal(struct {
		Name string `json:"name"`
		Lang string `json:"lang"`
	}{
		Name: nameNode.Name,
		Lang: language,
	})
}
