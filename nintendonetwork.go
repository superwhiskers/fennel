/*

nintendonetwork.go -
contains things for interacting with the nintendo network api

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

/*

NintendoNetworkErrorXML is a struct that holds information from a nintendo network error xml

*/
type NintendoNetworkErrorXML struct {
	XMLName xml.Name `xml:"errors"`
	Cause   string   `xml:"error>cause"`
	Code    int      `xml:"error>code"`
	Message string   `xml:"error>message"`
}

/*

NintendoNetworkClientInformation is a struct that holds data that is used to make the servers believe we are an actual 3ds or wiiu

*/
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

/*

NintendoNetworkClient is a struct that holds data used for connecting to nintendo network servers

*/
type NintendoNetworkClient struct {
	AccountServerAPIEndpoint string
	HTTPClient               *http.Client
	ClientInformation        NintendoNetworkClientInformation
}

/*

NewNintendoNetworkClient is a constructor function for creating a client to nintendo network servers

*/
func NewNintendoNetworkClient(accountServer string, certificatePath string, keyPath string, nnClientInfo NintendoNetworkClientInformation) (NintendoNetworkClient, error) {

	// load the certificate and key
	keyPair, err := tls.LoadX509KeyPair(certificatePath, keyPath)

	// handle errors
	if err != nil {

		// if there is one, return it
		return NintendoNetworkClient{}, err

	}

	// then we create a *http.Client with tls
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{keyPair},
				ClientAuth:         tls.RequireAndVerifyClientCert,
				InsecureSkipVerify: true,
			},
		},
	}

	// then create a NintendoNetworkClient
	nnClient := NintendoNetworkClient{
		AccountServerAPIEndpoint: accountServer,
		HTTPClient:               httpClient,
		ClientInformation:        nnClientInfo,
	}

	// and then return it
	return nnClient, nil

}

/*

DoesUserExist is a method on NintendoNetworkClient that requests info about a user from the nintendo network servers

*/
func (c NintendoNetworkClient) DoesUserExist(nnid string) (bool, NintendoNetworkErrorXML, error) {

	// construct the request
	request, err := http.NewRequest("GET", strings.Join([]string{c.AccountServerAPIEndpoint, "/people/", nnid}, ""), nil)

	// check if there was an error
	if err != nil {

		// if there was one, we return it
		return false, NintendoNetworkErrorXML{}, err

	}

	// set the headers
	request.Header.Set("X-Nintendo-Client-ID", c.ClientInformation.ClientID)
	request.Header.Set("X-Nintendo-Client-Secret", c.ClientInformation.ClientSecret)

	// perform the request
	res, err := c.HTTPClient.Do(request)

	// check if there was an error
	if err != nil {

		// if there was one, we return it
		return false, NintendoNetworkErrorXML{}, err

	}

	// close the response body when we are done
	defer res.Body.Close()

	// get the response body
	resData, err := ioutil.ReadAll(res.Body)

	// check if there were errors
	if err != nil {

		// return an error if there was one
		return false, NintendoNetworkErrorXML{}, err

	}

	// the error xml struct
	var errorXML NintendoNetworkErrorXML

	// attempt to parse it as xml
	err = xml.Unmarshal(resData, &errorXML)

	// check an error occured
	if err != nil {

		// if there was one, it means that the nnid does not exist
		return false, NintendoNetworkErrorXML{}, nil

	}

	// we check the error code
	if errorXML.Code == 100 {

		// it exists
		return true, errorXML, nil

	} else if errorXML.Code == 1104 {

		// user id format invalid
		return false, errorXML, errors.New("your user id format is invalid")

	} else if errorXML.Code == 4 {

		// invalid creds
		return false, errorXML, errors.New("there is an error in your credentials")

	}

	// if we get here, there was an unknown error
	return false, errorXML, errors.New("an unknown and unhandlable error occured")

}
