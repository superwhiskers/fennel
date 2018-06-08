package libninty

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
)

func TestNewNintendoNetworkClient(t *testing.T) {

	// nintendo network client information
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

	// load the certificate and key
	keyPair, err := tls.LoadX509KeyPair("keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem")

	// test failed if an error occured
	if err != nil {

		// fail the test
		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	// expected output generated from the input
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

	// generate a new client
	output, err := NewNintendoNetworkClient("https://account.pretendo.cc/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", nnClientInfo)

	// test failed if an error occured
	if err != nil {

		// fail the test
		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	// print what we got vs what we expected
	t.Logf("expected: %+v", expectedOutput)
	t.Logf("got: %+v", output)

	// compare the expected output with the output
	if output == expectedOutput {

		// if they do not match, fail
		t.Errorf("output mismatch...")

	}

}

func TestDoesUserExist(t *testing.T) {

	// nintendo network client information
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

	// generate a new client
	client, err := NewNintendoNetworkClient("https://account.nintendo.net/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", nnClientInfo)

	// test failed if an error occured
	if err != nil {

		// fail the test
		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	// now check if a user exists
	output, _, err := client.DoesUserExist("whiskers")

	// fail if an error occured
	if err != nil {

		// fail the test
		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	// print what we got vs what we expected
	t.Logf("expected: true")
	t.Logf("got: %+v", output)

	// compare the expected output with the output
	if output != true {

		fmt.Printf("failed")

		// if they do not match, fail
		t.Errorf("output mismatch...")

	}

}
