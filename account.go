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
	"strings"

	"github.com/superwhiskers/libninty/errors"
	"github.com/superwhiskers/libninty/formats/xmls"
	"github.com/valyala/fasthttp"
)

// DoesUserExist checks if a user with the given nnid exists on nintendo network
func (c *Client) DoesUserExist(nnid string) (bool, xmls.ErrorXML, error) {

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	requestHeader := fasthttp.RequestHeader{}

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	requestHeader.SetMethod("GET")
	request.Header = requestHeader
	request.SetRequestURI(strings.Join([]string{c.AccountServerAPIEndpoint, "/people/", nnid}, ""))

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

		if error.Code == errors.AccountIDExistsError {

			return true, xmls.NilErrorXML, nil

		}

	}

	return false, errorXML, nil

}

// GetEULA retrieves the Nintendo Network EULA for the specified country.
// if version is `@latest`, it returns the latest version. otherwise, it returns the specified version
func (c *Client) GetEULA(countryCode, version string) (xmls.AgreementXML, xmls.ErrorXML, error) {

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	requestHeader := fasthttp.RequestHeader{}

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	requestHeader.SetMethod("GET")
	request.Header = requestHeader
	request.SetRequestURI(strings.Join([]string{c.AccountServerAPIEndpoint, "/content/agreements/Nintendo-Network-EULA/", countryCode, "/", version}, ""))

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
