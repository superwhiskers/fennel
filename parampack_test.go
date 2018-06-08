package libninty

import "testing"

func TestDecodeParampack(t *testing.T) {

	// test parampack
	// (obatained from ariankordi.net, it's a valid parampack)
	testParampack := "XHRpdGxlX2lkXFxhY2Nlc3Nfa2V5XDBccGxhdGZvcm1faWRcMVxyZWdpb25faWRcMlxsYW5ndWFnZV9pZFwxXGNvdW50cnlfaWRcNDlcYXJlYV9pZFwwXG5ldHdvcmtfcmVzdHJpY3Rpb25cMFxmcmllbmRfcmVzdHJpY3Rpb25cMFxyYXRpbmdfcmVzdHJpY3Rpb25cMjBccmF0aW5nX29yZ2FuaXphdGlvblwwXHRyYW5zZmVyYWJsZV9pZFwxMjc1NjE0NDg4NDQ1Mzg5ODc4Mlx0el9uYW1lXEFtZXJpY2EvTmV3X1lvcmtcdXRjX29mZnNldFwtMTQ0MDBccmVtYXN0ZXJfdmVyc2lvblwwXA=="

	// expected output generated from the test parampack
	expectedOutput := Parampack{
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

	// see if it decodes
	output, err := DecodeParampack(testParampack)

	// test failed if an error occured
	if err != nil {

		// fail the test
		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	// print what we got vs what we expected
	t.Logf("expected: %+v", expectedOutput)
	t.Logf("got: %+v", output)

	// compare the expected output with the output
	if output != expectedOutput {

		// if they do not match, fail
		t.Errorf("output mismatch...")

	}

}

func TestUnstringifyParampack(t *testing.T) {

	// test parampack
	// (obatained by unbase64ing the parampack from ariankordi.net, it's also a valid parampack)
	testParampack := "\\title_id\\\\access_key\\0\\platform_id\\1\\region_id\\2\\language_id\\1\\country_id\\49\\area_id\\0\\network_restriction\\0\\friend_restriction\\0\\rating_restriction\\20\\rating_organization\\0\\transferable_id\\12756144884453898782\\tz_name\\America/New_York\\utc_offset\\-14400\\remaster_version\\0\\"

	// expected output generated from the test parampack
	expectedOutput := Parampack{
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

	// see if it decodes
	output := UnstringifyParampack(testParampack)

	// print what we got vs what we expected
	t.Logf("expected: %+v", expectedOutput)
	t.Logf("got: %+v", output)

	// compare the expected output with the output
	if output != expectedOutput {

		// if they do not match, fail
		t.Errorf("output mismatch...")

	}

}

func TestEncodeParampack(t *testing.T) {

	// parampack input
	testParampack := Parampack{
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

	// sample parampack output
	expectedOutput := "XHRpdGxlX2lkXFxhY2Nlc3Nfa2V5XDBccGxhdGZvcm1faWRcMVxyZWdpb25faWRcMlxsYW5ndWFnZV9pZFwxXGNvdW50cnlfaWRcNDlcYXJlYV9pZFwwXG5ldHdvcmtfcmVzdHJpY3Rpb25cMFxmcmllbmRfcmVzdHJpY3Rpb25cMFxyYXRpbmdfcmVzdHJpY3Rpb25cMjBccmF0aW5nX29yZ2FuaXphdGlvblwwXHRyYW5zZmVyYWJsZV9pZFwxMjc1NjE0NDg4NDQ1Mzg5ODc4Mlx0el9uYW1lXEFtZXJpY2EvTmV3X1lvcmtcdXRjX29mZnNldFwtMTQ0MDBccmVtYXN0ZXJfdmVyc2lvblwwXA=="

	// see if it encodes
	output := testParampack.EncodeParampack()

	// print what we got vs what we expected
	t.Logf("expected: %+v", expectedOutput)
	t.Logf("got: %+v", output)

	// compare the expected output with the output
	if output != expectedOutput {

		// if they do not match, fail
		t.Errorf("output mismatch...")

	}

}

func TestStringifyParampack(t *testing.T) {

	// parampack input
	testParampack := Parampack{
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

	// sample parampack output
	expectedOutput := "\\title_id\\\\access_key\\0\\platform_id\\1\\region_id\\2\\language_id\\1\\country_id\\49\\area_id\\0\\network_restriction\\0\\friend_restriction\\0\\rating_restriction\\20\\rating_organization\\0\\transferable_id\\12756144884453898782\\tz_name\\America/New_York\\utc_offset\\-14400\\remaster_version\\0\\"

	// see if it decodes
	output := testParampack.StringifyParampack()

	// print what we got vs what we expected
	t.Logf("expected: %+v", expectedOutput)
	t.Logf("got: %+v", output)

	// compare the expected output with the output
	if output != expectedOutput {

		// if they do not match, fail
		t.Errorf("output mismatch...")

	}

}
