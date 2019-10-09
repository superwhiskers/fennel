/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

package xmls

import (
	"encoding/xml"
	"fmt"
)

// NilErrorXML represents an ErrorXML with no value
var NilErrorXML = ErrorXML{}

// ErrorXML represents a nintendo network error xml sheet
type ErrorXML struct {
	XMLName xml.Name        `xml:"errors"`
	Errors  []ErrorXMLError `xml:"error"`
}

// FormatXML formats an ErrorXML struct as a byte array
func (e ErrorXML) FormatXML() ([]byte, error) {

	return xml.Marshal(e)

}

// ErrorXMLError represents an error of an ErrorXML
type ErrorXMLError struct {
	Cause   string `xml:"cause"`
	Code    string `xml:"code"`
	Message string `xml:"message"`
}

// Error returns a console-friendly version of the error contained in the struct
func (e ErrorXMLError) Error() string {

	return fmt.Sprintf("code %s: %s", e.Code, e.Message)

}

// ParseErrorXML parses a nintendo network error xml sheet to an ErrorXML struct
func ParseErrorXML(errorXML []byte) (errorXMLParsed ErrorXML, err error) {

	err = xml.Unmarshal(errorXML, &errorXMLParsed)
	if err != nil {

		return NilErrorXML, err

	}

	return errorXMLParsed, nil

}
