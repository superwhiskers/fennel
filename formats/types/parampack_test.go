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

package types

import "testing"

func TestParampack(t *testing.T) {

	var (
		expectedB64Parampack    = "XHRpdGxlX2lkXFxhY2Nlc3Nfa2V5XDBccGxhdGZvcm1faWRcMVxyZWdpb25faWRcMlxsYW5ndWFnZV9pZFwxXGNvdW50cnlfaWRcNDlcYXJlYV9pZFwwXG5ldHdvcmtfcmVzdHJpY3Rpb25cMFxmcmllbmRfcmVzdHJpY3Rpb25cMFxyYXRpbmdfcmVzdHJpY3Rpb25cMjBccmF0aW5nX29yZ2FuaXphdGlvblwwXHRyYW5zZmVyYWJsZV9pZFwxMjc1NjE0NDg4NDQ1Mzg5ODc4Mlx0el9uYW1lXEFtZXJpY2EvTmV3X1lvcmtcdXRjX29mZnNldFwtMTQ0MDBccmVtYXN0ZXJfdmVyc2lvblwwXA=="
		expectedStringParampack = "\\title_id\\\\access_key\\0\\platform_id\\1\\region_id\\2\\language_id\\1\\country_id\\49\\area_id\\0\\network_restriction\\0\\friend_restriction\\0\\rating_restriction\\20\\rating_organization\\0\\transferable_id\\12756144884453898782\\tz_name\\America/New_York\\utc_offset\\-14400\\remaster_version\\0\\"
		expectedParampack       = Parampack{
			TitleID:            "",
			AccessKey:          "0",
			PlatformID:         1,
			RegionID:           2,
			LanguageID:         1,
			CountryID:          49,
			AreaID:             0,
			NetworkRestriction: 0,
			FriendRestriction:  0,
			RatingRestriction:  20,
			RatingOrganization: 0,
			TransferableID:     "12756144884453898782",
			TimezoneName:       "America/New_York",
			UTCOffset:          -14400,
			RemasterVersion:    0,
		}
	)

	parampack, err := ParseParampack(expectedB64Parampack)
	if err != nil {

		t.Errorf("couldn't parse the base64ed parampack to a Parampack. error: %v\n", err)

	}

	if parampack != expectedParampack {

		t.Errorf("the Parampack doesn't match the expected one")

	}

	t.Logf("got Parampack: %#v\n", parampack)

	parampack = ParseStringParampack(expectedStringParampack)

	if parampack != expectedParampack {

		t.Errorf("the Parampack doesn't match the expected one")

	}

	t.Logf("got Parampack: %#v\n", parampack)

	encodedParampack := expectedParampack.FormatSource()

	if encodedParampack != expectedB64Parampack {

		t.Errorf("the base64 encoded parampack doesn't match the expected one")

	}

	t.Logf("got b64 parampack: %s\n", encodedParampack)

	encodedParampack = expectedParampack.FormatString()

	if encodedParampack != expectedStringParampack {

		t.Errorf("the string encoded parampack doesn't match the expected one")

	}

	t.Logf("got string parampack: %s\n", encodedParampack)

}
