package datamodel

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// MarshalXML generate XML output for Edition
func (preconstructed PreconstructedInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if (PreconstructedInfo{} == preconstructed) {
		fmt.Print("called")
		return nil
	}
	if preconstructed.Decks > 0 {
		start.Attr = []xml.Attr{xml.Attr{Name: xml.Name{Local: "decks"}, Value: strconv.Itoa(preconstructed.Size)}}
		fmt.Println(len(start.Attr))
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"}, Value: preconstructed.Type})
		fmt.Println(len(start.Attr))
	}
	err = e.EncodeToken(start)
	return
	/*if err != nil {
		fmt.Print("error: ", err)
		return
	}
	fmt.Print(start)

	return e.EncodeToken(preconstructed)
	return nil*/
}
