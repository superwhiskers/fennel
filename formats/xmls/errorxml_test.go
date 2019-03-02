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
