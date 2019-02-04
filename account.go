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
	"strings"

	"github.com/superwhiskers/fennel/errors"
	"github.com/superwhiskers/fennel/formats/xmls"
	"github.com/valyala/fasthttp"
)

// AccountServerClient implements a client for the nintendo network account servers
type AccountServerClient struct {
	APIEndpoint       string
	HTTPClient        *fasthttp.Client
	ClientInformation ClientInformation
}

// NewAccountServerClient is a constructor function for creating a client for the nintendo network account servers
func NewAccountServerClient(accountServer string, certificatePath string, keyPath string, clientInfo ClientInformation) (*AccountServerClient, error) {

	keyPair, err := tls.LoadX509KeyPair(certificatePath, keyPath)
	if err != nil {

		return &AccountServerClient{}, err

	}

	httpClient := &fasthttp.Client{
		TLSConfig: &tls.Config{
			Certificates:       []tls.Certificate{keyPair},
			ClientAuth:         tls.RequireAndVerifyClientCert,
			InsecureSkipVerify: true,
		},
	}

	client := &AccountServerClient{
		APIEndpoint:       accountServer,
		HTTPClient:        httpClient,
		ClientInformation: clientInfo,
	}

	return client, nil

}

// Do makes a request with headers set to make it look like you are a nintendo console
func (c *AccountServerClient) Do(request *fasthttp.Request, response *fasthttp.Response) error {

	request.Header.Set("X-Nintendo-Client-ID", c.ClientInformation.ClientID)
	request.Header.Set("X-Nintendo-FPD-Version", c.ClientInformation.FPDVersion)
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

	return c.HTTPClient.Do(request, response)

}

// DoesUserExist checks if a user with the given nnid exists on nintendo network
func (c *AccountServerClient) DoesUserExist(nnid string) (bool, xmls.ErrorXML, error) {

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	requestHeader := fasthttp.RequestHeader{}

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	requestHeader.SetMethod("GET")
	request.Header = requestHeader
	request.SetRequestURI(strings.Join([]string{c.APIEndpoint, "/people/", nnid}, ""))

	err := c.Do(request, response)
	if err != nil {

		return false, xmls.NilErrorXML, err

	}

	if response.StatusCode() == 200 {

		return false, xmls.NilErrorXML, nil

	}

	errorXML, err := xmls.ParseErrorXML(response.Body())
	if err != nil {

		return false, xmls.NilErrorXML, err

	}

	for _, error := range errorXML.Errors {

		if error.Code == string(errors.AccountIDExistsError) {

			return true, xmls.NilErrorXML, nil

		}

	}

	return false, errorXML, nil

}

// GetEULA retrieves the Nintendo Network EULA for the specified country.
// if version is `@latest`, it returns the latest version. otherwise, it returns the specified version
func (c *AccountServerClient) GetEULA(countryCode, version string) (xmls.AgreementXML, xmls.ErrorXML, error) {

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	requestHeader := fasthttp.RequestHeader{}

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	requestHeader.SetMethod("GET")
	request.Header = requestHeader
	request.SetRequestURI(strings.Join([]string{c.APIEndpoint, "/content/agreements/Nintendo-Network-EULA/", countryCode, "/", version}, ""))

	err := c.Do(request, response)
	if err != nil {

		return xmls.NilAgreementXML, xmls.NilErrorXML, err

	}

	if response.StatusCode() == 200 {

		axml, err := xmls.ParseAgreementXML(response.Body())
		if err != nil {

			return xmls.NilAgreementXML, xmls.NilErrorXML, err

		}

		return axml, xmls.NilErrorXML, nil

	}

	errorXML, err := xmls.ParseErrorXML(response.Body())
	if err != nil {

		return xmls.NilAgreementXML, xmls.NilErrorXML, err

	}

	return xmls.NilAgreementXML, errorXML, nil

}
