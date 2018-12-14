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

package formats

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
