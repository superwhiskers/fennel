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
	"time"
)

// NilAgreementXML represents an AgreementXML with no value
var NilAgreementXML = AgreementXML{
	Agreements: []AgreementXMLAgreement{},
}

// AgreementXML represents a nintendo network agreement xml sheet
type AgreementXML struct {
	XMLName    xml.Name                `xml:"agreements"`
	Agreements []AgreementXMLAgreement `xml:"agreement"`
}

// FormatXML formats the AgreementXML as a byte array
func (a AgreementXML) FormatXML() ([]byte, error) {

	agreementxml, err := xml.Marshal(a)
	return agreementxml, err

}

// AgreementXMLAgreement represents an agreement of an AgreementXML
type AgreementXMLAgreement struct {
	Country        string `xml:"country"`
	Language       string `xml:"language"`
	LanguageName   string `xml:"language_name"`
	PublishDateRaw string `xml:"publish_date"`
	AcceptText     CData  `xml:"texts>agree_text"`
	CancelText     CData  `xml:"texts>non_agree_text"`
	Title          CData  `xml:"texts>main_title"`
	Body           CData  `xml:"texts>main_text"`
	Type           string `xml:"type"`
	Version        string `xml:"version"`
}

// PublishDate returns the publish date as time.Time
func (a AgreementXMLAgreement) PublishDate() (time.Time, error) {

	time, err := time.Parse(time.RFC3339, a.PublishDateRaw)
	return time, err

}

// CData represents cdata in xml
type CData struct {
	Data string `xml:",cdata"`
}

// ParseAgreementXML parses a nintendo network agreement xml to an AgreementXML
func ParseAgreementXML(agreementXML []byte) (AgreementXML, error) {

	var agreementXMLParsed AgreementXML

	err := xml.Unmarshal(agreementXML, &agreementXMLParsed)
	if err != nil {

		return AgreementXML{}, err

	}

	return agreementXMLParsed, nil

}
