package datamodel

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MarshalJSON generate JSON output for Edition
func (u Edition) MarshalJSON() ([]byte, error) {
	var vaultValue bool
	vaultValue = true
	var onlineValue bool
	vaultValue = false
	var preconstructedItem *PreconstructedInfo
	preconstructedItem = &u.Preconstructed
	var cardsItem *CardsComposition
	cardsItem = &u.Cards
	if u.Vault == nil {
		vaultValue = false
	}

	if u.Online == nil {
		onlineValue = false
	}

	if (PreconstructedInfo{} == u.Preconstructed) {
		fmt.Println("PreconstructedEmpty")
		preconstructedItem = nil
	}

	if (CardsComposition{} == u.Cards) {
		cardsItem = nil
	}
	fmt.Println(PreconstructedInfo{} == u.Preconstructed)
	fmt.Println("Cards ", CardsComposition{} == u.Cards)
	type AliasEdition Edition
	return json.Marshal(struct {
		AliasEdition
		Cards          *CardsComposition   `json:"cards,omitempty"`
		Preconstructed *PreconstructedInfo `json:"preconstructed,omitempty"`
		Vault          bool                `json:"vault"`
		Online         bool                `json:"online"`
	}{
		AliasEdition:   AliasEdition(u),
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
