/*

fennel - nintendo network utility library for golang
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

package xmls

import (
	"encoding/xml"
	"fmt"
)

// NilErrorXML represents an ErrorXML with no value
var NilErrorXML = ErrorXML{
	Errors: []ErrorXMLError{},
}

// ErrorXML represents a nintendo network error xml sheet
type ErrorXML struct {
	XMLName xml.Name        `xml:"errors"`
	Errors  []ErrorXMLError `xml:"error"`
}

// FormatXML formats an ErrorXML struct as a byte array
func (e ErrorXML) FormatXML() ([]byte, error) {

	errorxml, err := xml.Marshal(e)
	return errorxml, err

}

// ErrorXMLError represents an error of an ErrorXML
type ErrorXMLError struct {
	Cause   string `xml:"cause"`
	Code    string `xml:"code"`
	Message string `xml:"message"`
}

// Error returns a console-friendly version of the error contained in the struct
func (e ErrorXMLError) Error() string {

	return fmt.Sprintf("code %s, caused by: %s: %s", e.Code, e.Cause, e.Message)

}

// ParseErrorXML parses a nintendo network error xml sheet to an ErrorXML struct
func ParseErrorXML(errorXML []byte) (errorXMLParsed ErrorXML, err error) {

	err = xml.Unmarshal(errorXML, &errorXMLParsed)
	if err != nil {

		return NilErrorXML, err

	}

	return errorXMLParsed, nil

}
