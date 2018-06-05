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
		t.FailNow()

	}

}
