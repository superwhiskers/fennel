/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

package types

import (
	"encoding/base64"
	"encoding/hex"
)

// ParseServicetoken is a function that takes a base64ed servicetoken and converts it to hexadecimal
func ParseServicetoken(servicetoken string) (string, error) {

	decodedServicetoken, err := base64.StdEncoding.DecodeString(servicetoken)
	if err != nil {

		return "", err

	}

	return hex.EncodeToString(decodedServicetoken), nil

}
