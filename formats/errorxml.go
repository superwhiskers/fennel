/*

libninty - nintendo network utility library for golang
Copyright (C) 2018 superwhiskers <whiskerdev@protonmail.com>

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

package formats

import (
	"encoding/xml"
)

// ErrorXML represents a nintendo network error xml sheet
type ErrorXML struct {
	XMLName xml.Name `xml:"errors"`
	Errors []ErrorXMLError `xml:"error"`
}

// ErrorXMLError represents an error of an ErrorXML
type ErrorXMLError struct {
	Cause string `xml:"cause"`
	Code  string `xml:"code"`
	Message string `xml:"message"`
}

// ParseErrorXML parses a nintendo network error xml sheet to an ErrorXML struct
func ParseErrorXML(errorXML []byte) (ErrorXML, error) {

	var errorXMLParsed ErrorXML

	err := xml.Unmarshal(errorXML, &errorXMLParsed)
	if err != nil {

		return ErrorXML{}, err

	}

	return errorXMLParsed, nil

}

// FormatXML formats an ErrorXML struct as a nintendo network error xml sheet
func (eXML ErrorXML) FormatXML() ([]byte, error) {

	errorxml, err := xml.Marshal(eXML)
	return errorxml, err

}
