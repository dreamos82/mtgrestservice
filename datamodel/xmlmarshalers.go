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
		start.Attr = []xml.Attr{xml.Attr{Name: xml.Name{Local: "decks"}, Value: strconv.Itoa(preconstructed.Size)}}
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
