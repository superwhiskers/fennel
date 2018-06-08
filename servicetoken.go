/*

servicetoken.go -
contains things for making encrypted servicetokens easier to deal with

*/

package libninty

import (
	"encoding/base64"
	"encoding/hex"
)

/*

DecodeServiceToken is a function that takes a base64ed servicetoken and converts it to hexadecimal

*/
func DecodeServiceToken(serviceToken string) (string, error) {

	// decode it from base64
	decodedServiceToken, err := base64.StdEncoding.DecodeString(serviceToken)

	// if there is an error
	if err != nil {

		// exit the function and return the error
		return "", err

	}

	// temporary workaround for now
	return hex.EncodeToString(decodedServiceToken), nil

}
