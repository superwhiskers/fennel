/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/

package xmls

import "encoding/xml"

// NillMappedIDsXML represents a MappedIDsXML with no value
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
