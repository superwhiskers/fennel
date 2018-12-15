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

	"github.com/valyala/fasthttp"
)

// ClientInformation holds data for headers sent in requests to nintendo network
type ClientInformation struct {
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

// Client implements a client for nintendo network
type Client struct {
	AccountServerAPIEndpoint string
	HTTPClient               *fasthttp.Client
	ClientInformation        ClientInformation
}

// NewClient is a constructor function for creating a client to nintendo network servers
func NewClient(accountServer string, certificatePath string, keyPath string, clientInfo ClientInformation) (*Client, error) {

	keyPair, err := tls.LoadX509KeyPair(certificatePath, keyPath)
	if err != nil {

		return &Client{}, err

	}

	httpClient := &fasthttp.Client{
		TLSConfig: &tls.Config{
			Certificates:       []tls.Certificate{keyPair},
			ClientAuth:         tls.RequireAndVerifyClientCert,
			InsecureSkipVerify: true,
		},
	}

	client := &Client{
		AccountServerAPIEndpoint: accountServer,
		HTTPClient:               httpClient,
		ClientInformation:        clientInfo,
	}

	return client, nil

}

// Do makes a request with headers set to make it look like you are a nintendo console
func (c *Client) Do(request *fasthttp.Request, response *fasthttp.Response) error {

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

	request.Header.Set("X-Nintendo-FPD-Version", "0000")

	return c.HTTPClient.Do(request, response)

}
