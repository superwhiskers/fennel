/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

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

import "strconv"

// TODO: add a 3ds tid unstringify

// TitleID reperesents a parsed titleid from any source format
type TitleID [16]byte

// FormatWiiU converts a TitleID to a wiiu titleid
func (tid TitleID) FormatWiiU() (string, error) {

	intTID, err := strconv.ParseUint(string(tid[:]), 16, 64)
	if err != nil {

		return "", err

	}

	wiiuTID := strconv.FormatUint(intTID, 10)

	for wiiuTID[0] == '0' {

		wiiuTID = wiiuTID[1:]

	}

	return wiiuTID, nil
}

// ParseWiiUTID parses a wiiu titleid to a TitleID
func ParseWiiUTID(wiiuTID string) (TitleID, error) {

	var tid TitleID

	intTID, err := strconv.ParseUint(wiiuTID, 10, 64)
	if err != nil {

		return TitleID{}, err

	}

	sliceTID := []byte(strconv.FormatUint(intTID, 16))

	for x := len(sliceTID); x < 16; x++ {

		sliceTID = append([]byte{'0'}, sliceTID...)

	}

	copy(tid[:], sliceTID[:])

	return tid, nil

}
