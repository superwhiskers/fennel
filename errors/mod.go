/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

package errors

import "fmt"

// FennelError implements a custom error type used in fennel
type FennelError struct {
	scope string
	error string
}

// Error formats the error held in a FennelError as a string
func (e FennelError) Error() string {

	return fmt.Sprintf("fennel: %s: %s", e.scope, e.error)

}

var (
	ByteBufferOverreadError = FennelError{
		scope: "bytebuffer",
		error: "read exceeds buffer capacity",
	}
	ByteBufferOverwriteError = FennelError{
		scope: "bytebuffer",
		error: "write exceeds buffer capacity",
	}
	ByteBufferInvalidIntegerSize = FennelError{
		scope: "bytebuffer",
		error: "invalid integer size specified",
	}
	ByteBufferInvalidEndianness = FennelError{
		scope: "bytebuffer",
		error: "invalid endianness specified",
	}
	ByteBufferInvalidByteCount = FennelError{
		scope: "bytebuffer",
		error: "invalid byte count requested",
	}
	BitfieldInvalidBit = FennelError{
		scope: "bitfield",
		error: "invalid bit value specified",
	}
	BitfieldOverreadError = FennelError{
		scope: "bitfield",
		error: "read exceeds bitfield capacity",
	}
	BitfieldOverwriteError = FennelError{
		scope: "bitfield",
		error: "write exceeds bitfield capacity",
	}
)
