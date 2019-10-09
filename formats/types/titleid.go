/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

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

	copy(tid[:], sliceTID)

	return tid, nil

}
