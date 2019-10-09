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
	"time"
)

// NilAgreementXML represents an AgreementXML with no value
var NilAgreementXML = AgreementXML{}

// AgreementXML represents a nintendo network agreement xml sheet
type AgreementXML struct {
	XMLName    xml.Name                `xml:"agreements"`
	Agreements []AgreementXMLAgreement `xml:"agreement"`
}

// FormatXML formats the AgreementXML as a byte array
func (a AgreementXML) FormatXML() ([]byte, error) {

	return xml.Marshal(a)

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

// ParseAgreementXML parses a nintendo network agreement xml to an AgreementXML
func ParseAgreementXML(agreementXML []byte) (agreementXMLParsed AgreementXML, err error) {

	err = xml.Unmarshal(agreementXML, &agreementXMLParsed)
	if err != nil {

		return NilAgreementXML, err

	}

	return agreementXMLParsed, nil

}
