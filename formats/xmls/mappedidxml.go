/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

package xmls

import "encoding/xml"

// NilMappedIDsXML represents a MappedIDsXML with no value
var NilMappedIDsXML = MappedIDsXML{}

// MappedIDsXML represents a nintendo network mappedids xml sheet
type MappedIDsXML struct {
	XMLName   xml.Name               `xml:"mapped_ids"`
	MappedIDs []MappedIDsXMLMappedID `xml:"mapped_id"`
}

// FormatXML formats the MappedIDsXML as a byte array
func (m MappedIDsXML) FormatXML() ([]byte, error) {

	return xml.Marshal(m)

}

// MappedIDsXMLMappedID represents a mapped id of a MappedIDsXML
type MappedIDsXMLMappedID struct {
	InID  string `xml:"in_id"`
	OutID string `xml:"out_id"`
}

// ParseMappedIDsXML parses a nintendo network mapped ids xml to a MappedIDsXML
func ParseMappedIDsXML(mappedidsXML []byte) (mappedidsXMLParsed MappedIDsXML, err error) {

	err = xml.Unmarshal(mappedidsXML, &mappedidsXMLParsed)
	if err != nil {

		return NilMappedIDsXML, err

	}

	return mappedidsXMLParsed, nil

}
