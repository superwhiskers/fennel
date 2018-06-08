package libninty

import "testing"

func TestDecodeServiceToken(t *testing.T) {

	// define a test servicetoken
	// (obatained from ariankordi.net, it's a banned servicetoken)
	serviceToken := "3YysZBU8Xm+Uqb60uTU69cksB4SrTGcVJtU9JlzJhMYcbnpBq25KgWlrPt18zOfa+JHJzF6Ha36NuGjQ+BCrRsfRmnGSMz5muA7GLV195inaPza2AQLoEdRp4qjsklle1oqswldQS8m3pAHpZAHYYOORoJe3KLF6uEvleJRiCfk="

	// define the expected output
	expectedOutput := "dd8cac64153c5e6f94a9beb4b9353af5c92c0784ab4c671526d53d265cc984c61c6e7a41ab6e4a81696b3edd7ccce7daf891c9cc5e876b7e8db868d0f810ab46c7d19a7192333e66b80ec62d5d7de629da3f36b60102e811d469e2a8ec92595ed68aacc257504bc9b7a401e96401d860e391a097b728b17ab84be578946209f9"

	// see if it decodes
	output, err := DecodeServiceToken(serviceToken)

	// test failed if an error occured
	if err != nil {

		// fail the test
		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	// print what we got vs what we expected
	t.Logf("expected: %s", expectedOutput)
	t.Logf("got: %s", output)

	// compare the expected output with the output
	if output != expectedOutput {

		// if they do not match, fail
		t.Errorf("output mismatch...")

	}

}
