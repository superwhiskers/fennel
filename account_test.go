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

package fennel

import (
	"crypto/tls"
	"testing"

	"github.com/valyala/fasthttp"
)

var clientInfo = ClientInformation{
	ClientID:     "ea25c66c26b403376b4c5ed94ab9cdea",
	ClientSecret: "d137be62cb6a2b831cad8c013b92fb55",
	DeviceCert:   "",
	Environment:  "L1",
	Country:      "US",
	Region:       "2",
	SysVersion:   "1111",
	Serial:       "1",
	DeviceID:     "1",
	DeviceType:   "",
	PlatformID:   "1",
	FPDVersion:   "0000",
}

func TestNewAccountServerClient(t *testing.T) {

	keyPair, err := tls.LoadX509KeyPair("keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem")
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	expectedOutput := &AccountServerClient{
		APIEndpoint: "https://account.nintendo.net/v1/api",
		HTTPClient: &fasthttp.Client{
			TLSConfig: &tls.Config{
				Certificates:       []tls.Certificate{keyPair},
				ClientAuth:         tls.RequireAndVerifyClientCert,
				InsecureSkipVerify: true,
			},
		},
		ClientInformation: clientInfo,
	}

	output, err := NewAccountServerClient("https://account.nintendo.net/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", clientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	if output.ClientInformation != expectedOutput.ClientInformation {

		t.Errorf("invalid output")

	}

}

func TestDoesUserExist(t *testing.T) {

	client, err := NewAccountServerClient("https://account.nintendo.net/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", clientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	output, exml, err := client.DoesUserExist("abcdefg")
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	if len(exml.Errors) != 0 {

		t.Errorf("expected no error to occur, instead got %v\n", exml.Errors[0])

	}

	if output != true {

		t.Errorf("invalid output")

	}

}

func TestGetEULA(t *testing.T) {

	client, err := NewAccountServerClient("https://account.nintendo.net/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", clientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	output, exml, err := client.GetEULA("US", "@latest")
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	if len(exml.Errors) != 0 {

		t.Errorf("expected no error to occur, instead got %v\n", exml.Errors[0])

	}

	for _, agreement := range output.Agreements {

		if agreement.Language == "en" && agreement.Title.Data == "Nintendo Network Services Agreement" {

			return

		}

	}

	t.Errorf("invalid output")

}

func TestGetPIDs(t *testing.T) {

	client, err := NewAccountServerClient("https://account.nintendo.net/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", clientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	output, exml, err := client.GetPIDs([]string{"abcdefg"})
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	if len(exml.Errors) != 0 {

		t.Errorf("expected no error to occur, instead got %v\n", exml.Errors[0])

	}

	if output[0] != 1799704789 {

		t.Errorf("invalid output")

	}
	
}

func TestGetNNIDs(t *testing.T) {

	client, err := NewAccountServerClient("https://account.nintendo.net/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", clientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	output, exml, err := client.GetNNIDs([]int64{0x6b4550d5})
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	if len(exml.Errors) != 0 {

		t.Errorf("expected no error to occur, instead got %v\n", exml.Errors[0])

	}

	if output[0] != "ABCDEFG" {

		t.Errorf("invalid output")

	}

}

func TestGetMiis(t *testing.T) {

	client, err := NewAccountServerClient("https://account.nintendo.net/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", clientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	_, exml, err := client.GetMiis([]int64{0x6b4550d5})
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	if len(exml.Errors) != 0 {

		t.Errorf("expected no error to occur, instead got %v\n", exml.Errors[0])

	}

}
