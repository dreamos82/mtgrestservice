package datamodel

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// MarshalXML generate XML output for Edition
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
	/*if err != nil {
		fmt.Print("error: ", err)
		return
	}
	fmt.Print(start)

	return e.EncodeToken(preconstructed)
	return nil*/
}
