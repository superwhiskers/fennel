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

import (
	"encoding/base64"
	"encoding/xml"

	"github.com/superwhiskers/fennel/formats/types"
)

// AccountMiiXML represents a nintendo network miis xml from the account server
type AccountMiiXML struct {
	XMLName xml.Name           `xml:"miis"`
	Miis    []AccountMiiXMLMii `xml:"mii,omitempty"`
}

// NilAccountMiiXML represents an AccountMiiXML with no value
var NilAccountMiiXML = AccountMiiXML{}

// FormatXML formats the AccountMiiXML as a byte array
func (a AccountMiiXML) FormatXML() ([]byte, error) {

	return xml.Marshal(a)

}

// AccountMiiXMLMii represents a mii in a nintendo network miis xml
type AccountMiiXMLMii struct {
	Mii     types.Mii               `xml:"-"`
	Data    string                  `xml:"data"`
	ID      int64                   `xml:"id"`
	Name    string                  `xml:"name"`
	PID     int64                   `xml:"pid"`
	Primary string                  `xml:"primary"`
	UserID  string                  `xml:"user_id"`
	Images  []AccountMiiXMLMiiImage `xml:"images>image"`
}

// AccountMiiXMLMiiImage represents an image of a nintendo network miis xml
type AccountMiiXMLMiiImage struct {
	CachedURL string   `xml:"cached_url"`
	ID        int64    `xml:"id"`
	URL       string   `xml:"url"`
	Type      string   `xml:"type"`
}

// ParseAccountMiiXML parses a nintendo network account mii xml to an AccountMiiXML
func ParseAccountMiiXML(accountMiiXML []byte) (accountMiiXMLParsed AccountMiiXML, err error) {
	err = xml.Unmarshal(accountMiiXML, &accountMiiXMLParsed)
	if err != nil {

		return NilAccountMiiXML, err

	}

	var bytemii []byte
	for i, b64mii := range accountMiiXMLParsed.Miis {

		bytemii, err = base64.StdEncoding.DecodeString(b64mii.Data)
		if err != nil {

			return NilAccountMiiXML, err

		}

		accountMiiXMLParsed.Miis[i].Mii = types.ParseMii(bytemii)

	}

	return accountMiiXMLParsed, nil

}
