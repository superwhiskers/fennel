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

import "strconv"

// UnstringifyWiiUTID converts an encoded wiiu titleid into a proper one
func UnstringifyWiiUTID(stringifiedTID string) (string, error) {

	intTID, err := strconv.ParseUint(stringifiedTID, 10, 64)
	if err != nil {

		return "", err

	}

	tid := strconv.FormatUint(intTID, 16)

	if len(tid) != 16 {

		for x := len(tid); x < 16; x++ {

			tid = "0" + tid

		}

	}

	return tid, nil

}
