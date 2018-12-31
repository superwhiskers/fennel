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
