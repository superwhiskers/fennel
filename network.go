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

	"github.com/valyala/fasthttp"
)

// ClientInformation holds data for headers sent in requests to nintendo network
type ClientInformation struct {
	/* (mostly) constant information */
	PlatformID   string // this is pretty much always "1"
	DeviceType   string // this is either 1 or 2. 1 being debug, 2 being retail
	ClientID     string // this is pretty much always "a2efa818a34fa16b8afbc8a74eba3eda"
	ClientSecret string // this is pretty much always "c91cdb5658bd4954ade78533a339cf9a"
	FPDVersion   string // this is pretty much always "0000"
	Environment  string // this is pretty much always "L1"

	/* fluctuating information */
	DeviceID   string // device id obtained from your console
	Serial     string // string representing the serial number of your console
	SysVersion string // unsigned integer representing your system version
	Region     string // 1 = JPN, 2 = USA, 4 = EUR, 8 = AUS, 16 = CHN, 32 = KOR, 64 = TWN
	Country    string // two-letter country code
	DeviceCert string // device certificate obtained from your console
}

// ApplicationInformation holds data about the application accessing the api
type ApplicationInformation struct {
	TitleID            uint64
	ApplicationVersion uint64
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
