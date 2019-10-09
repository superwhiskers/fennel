/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

package types

import "testing"

func TestTitleID(t *testing.T) {

	var (
		expectedWiiUTID = "1407443871727872"
		expectedTID     = TitleID{'0', '0', '0', '5', '0', '0', '1', '0', '1', '0', '0', '4', '0', '1', '0', '0'}
	)

	tid, err := ParseWiiUTID(expectedWiiUTID)
	if err != nil {

		t.Errorf("couldn't parse the wiiu titleid to a TitleID. error: %v\n", err)

	}

	if tid != expectedTID {

		t.Errorf("the TitleID doesn't match the expected one")

	}

	t.Logf("got TitleID: %#v\n", tid)

	wiiuTID, err := expectedTID.FormatWiiU()
	if err != nil {

		t.Errorf("couldn't convert the TitleID back to a wiiu titleid. error: %v\n", err)

	}

	if wiiuTID != expectedWiiUTID {

		t.Errorf("the wiiu titleid doesn't match the expected one")

	}

	t.Logf("got wiiu tid: %#v\n", wiiuTID)

}
