/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

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

package errors

import "fmt"

// FennelError implements a custom error type used in fennel
type FennelError struct{
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
