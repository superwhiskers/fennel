package libninty

import "testing"

func TestUnstringifyWiiUTID(t *testing.T) {

	// define a test titleid
	stringifiedTitleID := "1407443871727872"

	// define the expected output
	expectedOutput := "0005001010040100"

	// see if it decodes
	output, err := UnstringifyWiiUTID(stringifiedTitleID)

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
		t.FailNow()

	}

}
