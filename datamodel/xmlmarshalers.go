package datamodel

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// MarshalXML generate XML output for PrecsontructedInfo
func (preconstructed PreconstructedInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if (PreconstructedInfo{} == preconstructed) {
		return nil
	}
	if preconstructed.Decks > 0 {
		start.Attr = []xml.Attr{xml.Attr{Name: xml.Name{Local: "decks"}, Value: strconv.Itoa(preconstructed.Decks)}}
	}
	if strings.Compare(preconstructed.Type, "") != 0 {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"}, Value: preconstructed.Type})
	}

	err = e.EncodeToken(start)
	e.EncodeElement(preconstructed.Size, xml.StartElement{Name: xml.Name{Local: "size"}})
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// MarshalXML generate XML output for Names
func (name NameNode) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if (NameNode{} == name) {
		return nil
	}
	if strings.Compare(name.Lang, "") != 0 {
		start.Attr = []xml.Attr{xml.Attr{Name: xml.Name{Local: "lang"}, Value: name.Lang}}
	}
	err = e.EncodeElement(name.Name, start)
	return err
}

// MarshalXML generate XML output for CardsComposition
func (cards CardsComposition) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if (CardsComposition{} == cards) {
		return nil
	}
	err = e.EncodeToken(start)
	if err != nil {
		return err
	}
	if cards.Common > 0 {
		err = e.EncodeElement(cards.Common, xml.StartElement{Name: xml.Name{Local: "common"}})
	}
	if cards.Lands > 0 {
		err = e.EncodeElement(cards.Lands, xml.StartElement{Name: xml.Name{Local: "lands"}})
	}
	if cards.Mythic > 0 {
		err = e.EncodeElement(cards.Mythic, xml.StartElement{Name: xml.Name{Local: "mythic"}})
	}
	if cards.Other > 0 {
		err = e.EncodeElement(cards.Other, xml.StartElement{Name: xml.Name{Local: "other"}})
	}
	if cards.Oversize > 0 {
		err = e.EncodeElement(cards.Oversize, xml.StartElement{Name: xml.Name{Local: "oversize"}})
	}
	if cards.Rare > 0 {
		err = e.EncodeElement(cards.Rare, xml.StartElement{Name: xml.Name{Local: "rare"}})
	}
	if cards.Uncommon > 0 {
		err = e.EncodeElement(cards.Uncommon, xml.StartElement{Name: xml.Name{Local: "uncommon"}})
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
