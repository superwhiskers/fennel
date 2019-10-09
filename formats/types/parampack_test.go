/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

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
