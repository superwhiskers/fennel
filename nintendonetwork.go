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
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// NintendoNetworkErrorXML is a struct that holds information from a nintendo network error xml
type NintendoNetworkErrorXML struct {
	XMLName xml.Name `xml:"errors"`
	Cause   string   `xml:"error>cause"`
	Code    int      `xml:"error>code"`
	Message string   `xml:"error>message"`
}

// NintendoNetworkClientInformation is a struct that holds data that is used to make the servers believe we are an actual 3ds or wiiu
type NintendoNetworkClientInformation struct {
	ClientID     string
	ClientSecret string
	DeviceCert   string
	Environment  string
	Country      string
	Region       string
	SysVersion   string
	Serial       string
	DeviceID     string
	DeviceType   string
	PlatformID   string
}

// NintendoNetworkClient is a struct that holds data used for connecting to nintendo network servers
type NintendoNetworkClient struct {
	AccountServerAPIEndpoint string
	HTTPClient               *http.Client
	ClientInformation        NintendoNetworkClientInformation
}

// ParseErrorXML is a function that parses error xml and returns a NintendoNetworkErrorXML struct
func ParseErrorXML(errorXML []byte) (NintendoNetworkErrorXML, error) {

	var errorXMLParsed NintendoNetworkErrorXML

	err := xml.Unmarshal(errorXML, &errorXMLParsed)
	if err != nil {

		return NintendoNetworkErrorXML{}, err

	}

	return errorXMLParsed, nil

}

// NewNintendoNetworkClient is a constructor function for creating a client to nintendo network servers
func NewNintendoNetworkClient(accountServer string, certificatePath string, keyPath string, nnClientInfo NintendoNetworkClientInformation) (NintendoNetworkClient, error) {

	keyPair, err := tls.LoadX509KeyPair(certificatePath, keyPath)
	if err != nil {

		return NintendoNetworkClient{}, err

	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{keyPair},
				ClientAuth:         tls.RequireAndVerifyClientCert,
				InsecureSkipVerify: true,
			},
		},
	}

	nnClient := NintendoNetworkClient{
		AccountServerAPIEndpoint: accountServer,
		HTTPClient:               httpClient,
		ClientInformation:        nnClientInfo,
	}

	return nnClient, nil

}

// Do is a method on NintendoNetworkClient that makes a request to any url with the nintendo network headers and clientcert
func (c NintendoNetworkClient) Do(request *http.Request) (*http.Response, error) {

	request.Header.Set("X-Nintendo-Client-ID", c.ClientInformation.ClientID)
	request.Header.Set("X-Nintendo-Client-Secret", c.ClientInformation.ClientSecret)
	request.Header.Set("X-Nintendo-Platform-ID", c.ClientInformation.PlatformID)
	request.Header.Set("X-Nintendo-Device-Type", c.ClientInformation.DeviceType)
	request.Header.Set("X-Nintendo-Device-ID", c.ClientInformation.DeviceID)
	request.Header.Set("X-Nintendo-Serial-Number", c.ClientInformation.Serial)
	request.Header.Set("X-Nintendo-System-Version", c.ClientInformation.SysVersion)
	request.Header.Set("X-Nintendo-Region", c.ClientInformation.Region)
	request.Header.Set("X-Nintendo-Country", c.ClientInformation.Country)
	request.Header.Set("X-Nintendo-Environment", c.ClientInformation.Environment)
	request.Header.Set("X-Nintendo-Device-Cert", c.ClientInformation.DeviceCert)
	return c.HTTPClient.Do(request)

}

// DoesUserExist is a method on NintendoNetworkClient that requests info about a user from the nintendo network servers
func (c NintendoNetworkClient) DoesUserExist(nnid string) (bool, NintendoNetworkErrorXML, error) {

	request, err := http.NewRequest("GET", strings.Join([]string{c.AccountServerAPIEndpoint, "/people/", nnid}, ""), nil)
	if err != nil {

		return false, NintendoNetworkErrorXML{}, err

	}

	res, err := c.Do(request)
	if err != nil {

		return false, NintendoNetworkErrorXML{}, err

	}

	defer res.Body.Close()

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {

		return false, NintendoNetworkErrorXML{}, err

	}

	var errorXML NintendoNetworkErrorXML

	err = xml.Unmarshal(resData, &errorXML)
	if err != nil {

		return false, NintendoNetworkErrorXML{}, err

	}

	switch errorXML.Code {

	case 100:
		return true, errorXML, nil

	case 1104:
		return false, errorXML, errors.New("your user id format is invalid")

	case 4:
		return false, errorXML, errors.New("there is an error in your credentials")

	}

	return false, errorXML, errors.New("an unknown and unhandlable error occured")

}
