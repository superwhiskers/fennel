/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

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
	Mii     *types.Mii              `xml:"-"`
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
	CachedURL string `xml:"cached_url"`
	ID        int64  `xml:"id"`
	URL       string `xml:"url"`
	Type      string `xml:"type"`
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
