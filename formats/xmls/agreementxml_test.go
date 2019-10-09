/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

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
