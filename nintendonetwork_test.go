/*

libninty - nintendo network utility library for golang
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

package libninty

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
)

func TestNewNintendoNetworkClient(t *testing.T) {

	nnClientInfo := NintendoNetworkClientInformation{
		ClientID:     "ea25c66c26b403376b4c5ed94ab9cdea",
		ClientSecret: "d137be62cb6a2b831cad8c013b92fb55",
		DeviceCert:   "",
		Environment:  "",
		Country:      "",
		Region:       "",
		SysVersion:   "",
		Serial:       "",
		DeviceID:     "",
		DeviceType:   "",
		PlatformID:   "",
	}

	keyPair, err := tls.LoadX509KeyPair("keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem")
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	expectedOutput := NintendoNetworkClient{
		AccountServerAPIEndpoint: "https://account.pretendo.cc/v1/api",
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					Certificates:       []tls.Certificate{keyPair},
					ClientAuth:         tls.RequireAndVerifyClientCert,
					InsecureSkipVerify: true,
				},
			},
		},
		ClientInformation: NintendoNetworkClientInformation{
			ClientID:     "ea25c66c26b403376b4c5ed94ab9cdea",
			ClientSecret: "d137be62cb6a2b831cad8c013b92fb55",
			DeviceCert:   "",
			Environment:  "",
			Country:      "",
			Region:       "",
			SysVersion:   "",
			Serial:       "",
			DeviceID:     "",
			DeviceType:   "",
			PlatformID:   "",
		},
	}

	output, err := NewNintendoNetworkClient("https://account.pretendo.cc/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", nnClientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	t.Logf("expected: %+v", expectedOutput)
	t.Logf("got: %+v", output)

	if output == expectedOutput {

		t.Errorf("output mismatch...")

	}

}

func TestDoesUserExist(t *testing.T) {

	nnClientInfo := NintendoNetworkClientInformation{
		ClientID:     "ea25c66c26b403376b4c5ed94ab9cdea",
		ClientSecret: "d137be62cb6a2b831cad8c013b92fb55",
		DeviceCert:   "",
		Environment:  "",
		Country:      "",
		Region:       "",
		SysVersion:   "",
		Serial:       "",
		DeviceID:     "",
		DeviceType:   "",
		PlatformID:   "",
	}

	client, err := NewNintendoNetworkClient("https://account.nintendo.net/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", nnClientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	output, _, err := client.DoesUserExist("whiskers")
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	t.Logf("expected: true")
	t.Logf("got: %+v", output)

	if output != true {

		t.Errorf("output mismatch...")

	}

}
