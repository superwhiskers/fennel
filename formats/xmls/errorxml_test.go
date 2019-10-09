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

func TestErrorXML(t *testing.T) {

	var (
		expectedErrorXML = ErrorXML{
			Errors: []ErrorXMLError{
				{
					Cause:   "client_id",
					Code:    "0004",
					Message: "API application invalid or incorrect application credentials",
				},
			},
		}
		expectedByteErrorXML = []byte("<errors><error><cause>client_id</cause><code>0004</code><message>API application invalid or incorrect application credentials</message></error></errors>")
	)

	errorxml, err := ParseErrorXML(expectedByteErrorXML)
	if err != nil {

		t.Errorf("couldn't parse the errorxml to an ErrorXML. error: %v\n", err)

	}

	for i, nnError := range errorxml.Errors {

		if nnError != expectedErrorXML.Errors[i] {

			t.Errorf("the ErrorXML doesn't match the expected one")

		}

	}

	t.Logf("got ErrorXML: %#v\n", errorxml)

	errorXMLByte, err := expectedErrorXML.FormatXML()
	if err != nil {

		t.Errorf("couldn't format the ErrorXML as an errorxml. error: %v\n", err)

	}

	if !bytes.Equal(errorXMLByte, expectedByteErrorXML) {

		t.Errorf("the byte-formatted errorxml doesn't match the expected one")

	}

	t.Logf("got byte-formatted errorxml: %s\n", string(errorXMLByte))

}
