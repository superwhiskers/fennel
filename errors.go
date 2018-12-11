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

package libninty

import "fmt"

// NintendoNetworkError implements a generic error type for all potential errors returned by the library
type NintendoNetworkError struct {
	Code    int
	Message string
}

// Error returns a stringified version of the error from NintendoNetworkError
func (err *NintendoNetworkError) Error() string {

	return fmt.Sprintf("error code %d: %s", err.Code, err.Message)

}

// error declarations. do not edit these
var (
	BadRequestError = &NintendoNetworkError{
		Code:    1600,
		Message: "unable to process request",
	}
	AccountIDExistsError = &NintendoNetworkError{
		Code:    100,
		Message: "account id already exists",
	}
	InvalidApplicationError = &NintendoNetworkError{
		Code:    4,
		Message: "invalid application credentials were provided",
	}
	InvalidAccountIDError = &NintendoNetworkError{
		Code:    1104,
		Message: "an invalid account id was provided",
	}

	UnknownError = &NintendoNetworkError{
		Code:    -1,
		Message: "an unknown error was returned from the server",
	}
)
