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
	"bytes"
	"testing"
)

func TestAgreementXML(t *testing.T) {

	var (
		expectedAgreementXML = AgreementXML{
			Agreements: []AgreementXMLAgreement{
				{
					Country:        "US",
					Language:       "en",
					LanguageName:   "English",
					PublishDateRaw: "2014-09-29T20:07:35",
					AcceptText: CData{
						Data: "i accept",
					},
					CancelText: CData{
						Data: "i don't accept",
					},
					Title: CData{
						Data: "heck",
					},
					Body: CData{
						Data: "heck",
					},
					Type:    "NINTENDO-NETWORK-EULA",
					Version: "0300",
				},
			},
		}
		expectedByteAgreementXML = []byte("<agreements><agreement><country>US</country><language>en</language><language_name>English</language_name><publish_date>2014-09-29T20:07:35</publish_date><texts><agree_text><![CDATA[i accept]]></agree_text><non_agree_text><![CDATA[i don't accept]]></non_agree_text><main_title><![CDATA[heck]]></main_title><main_text><![CDATA[heck]]></main_text></texts><type>NINTENDO-NETWORK-EULA</type><version>0300</version></agreement></agreements>")
	)

	agreementxml, err := ParseAgreementXML(expectedByteAgreementXML)
	if err != nil {

		t.Errorf("couldn't parse the agreementxml to an AgreementXML. error: %v\n", err)

	}

	for i, agreement := range agreementxml.Agreements {

		if agreement != expectedAgreementXML.Agreements[i] {

			t.Errorf("the AgreementXML doesn't match the expected one")

		}

	}

	t.Logf("got AgreementXML: %#v\n", agreementxml)

	agreementXMLByte, err := expectedAgreementXML.FormatXML()
	if err != nil {

		t.Errorf("couldn't format the AgreementXML as an agreementxml. error: %v\n", err)

	}

	if !bytes.Equal(agreementXMLByte, expectedByteAgreementXML) {

		t.Errorf("the byte-formatted agreementxml doesn't match the expected one")

	}

	t.Logf("got byte-formatted agreementxml: %s\n", string(agreementXMLByte))

}
