package libninty

import "strconv"

/*

UnstringifyWiiUTID converts an encoded wiiu titleid into a proper one

*/
func UnstringifyWiiUTID(stringifiedTID string) (string, error) {

	// convert the string to an int
	intTID, err := strconv.ParseUint(stringifiedTID, 10, 64)

	// return an error if there is one
	if err != nil {

		// return the error
		return "", err

	}

	// convert it back to a string
	tid := strconv.FormatUint(intTID, 16)

	// pad it to 16 characters
	if len(tid) != 16 {

		// loop until it is 16 characters
		for x := len(tid); x < 16; x++ {

			// pad with zeroes
			tid = "0" + tid

		}

	}

	// return the zero padded tid
	return tid, nil

}
