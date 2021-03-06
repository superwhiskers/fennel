/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

package fennel

import (
	"crypto/tls"
	"strconv"
	"strings"

	"github.com/superwhiskers/fennel/errors"
	"github.com/superwhiskers/fennel/formats/xmls"
	"github.com/superwhiskers/fennel/utils"
	"github.com/valyala/fasthttp"
)

// AccountServerClient implements a client for the nintendo network account servers
type AccountServerClient struct {
	APIEndpoint       string
	HTTPClient        *fasthttp.Client
	ClientInformation ClientInformation
	ReqHeader         fasthttp.RequestHeader
}

// NewAccountServerClient is a constructor function for creating a client for the nintendo network account servers
func NewAccountServerClient(accountServer string, certificate []byte, key []byte, clientInfo ClientInformation) (*AccountServerClient, error) {

	keyPair, err := tls.X509KeyPair(certificate, key)
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

// GetMiis retrieves the miis for the provided pids
func (c *AccountServerClient) GetMiis(pids []int64) (xmls.AccountMiiXML, xmls.ErrorXML, error) {

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	requestHeader := fasthttp.RequestHeader{}

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	requestHeader.SetMethod("GET")
	request.Header = requestHeader

	spids := utils.ConvertInt64SliceToStringSlice(pids)
	request.SetRequestURI(strings.Join([]string{c.APIEndpoint, "/miis?pids=", strings.Join(spids, ",")}, ""))

	err := c.Do(request, response)
	if err != nil {

		return xmls.AccountMiiXML{}, xmls.NilErrorXML, err

	}

	if response.StatusCode() == 200 {

		amxml, err := xmls.ParseAccountMiiXML(response.Body())
		if err != nil {

			return xmls.AccountMiiXML{}, xmls.NilErrorXML, err

		}

		return amxml, xmls.NilErrorXML, nil

	}

	errorXML, err := xmls.ParseErrorXML(response.Body())
	if err != nil {

		return xmls.AccountMiiXML{}, xmls.NilErrorXML, err

	}

	return xmls.AccountMiiXML{}, errorXML, nil

}

// GetPIDs retrieves the pids for the provided nnids
func (c *AccountServerClient) GetPIDs(nnids []string) ([]int64, xmls.ErrorXML, error) {

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	requestHeader := fasthttp.RequestHeader{}

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	requestHeader.SetMethod("GET")
	request.Header = requestHeader
	request.SetRequestURI(strings.Join([]string{c.APIEndpoint, "/admin/mapped_ids?input_type=user_id&output_type=pid&input=", strings.Join(nnids, ",")}, ""))

	err := c.Do(request, response)
	if err != nil {

		return []int64{}, xmls.NilErrorXML, err

	}

	if response.StatusCode() == 200 {

		mixml, err := xmls.ParseMappedIDsXML(response.Body())
		if err != nil {

			return []int64{}, xmls.NilErrorXML, err

		}

		mis := []int64{}
		var id int64
		for _, mi := range mixml.MappedIDs {

			id, err = strconv.ParseInt(mi.OutID, 10, 64)
			if err != nil {

				return []int64{}, xmls.NilErrorXML, err

			}

			mis = append(mis, id)

		}

		return mis, xmls.NilErrorXML, nil

	}

	errorXML, err := xmls.ParseErrorXML(response.Body())
	if err != nil {

		return []int64{}, xmls.NilErrorXML, err

	}

	return []int64{}, errorXML, nil

}

// GetNNIDs retrieves the nnids for the provided pids
func (c *AccountServerClient) GetNNIDs(pids []int64) ([]string, xmls.ErrorXML, error) {

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	requestHeader := fasthttp.RequestHeader{}

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	requestHeader.SetMethod("GET")
	request.Header = requestHeader

	spids := utils.ConvertInt64SliceToStringSlice(pids)
	request.SetRequestURI(strings.Join([]string{c.APIEndpoint, "/admin/mapped_ids?input_type=pid&output_type=user_id&input=", strings.Join(spids, ",")}, ""))

	err := c.Do(request, response)
	if err != nil {

		return []string{}, xmls.NilErrorXML, err

	}

	if response.StatusCode() == 200 {

		mixml, err := xmls.ParseMappedIDsXML(response.Body())
		if err != nil {

			return []string{}, xmls.NilErrorXML, err

		}

		mis := []string{}
		for _, mi := range mixml.MappedIDs {

			mis = append(mis, mi.OutID)

		}

		return mis, xmls.NilErrorXML, nil

	}

	errorXML, err := xmls.ParseErrorXML(response.Body())
	if err != nil {

		return []string{}, xmls.NilErrorXML, err

	}

	return []string{}, errorXML, nil

}
