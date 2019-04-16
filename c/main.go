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

package main

// #include <stdlib.h>
// #include "types.h"
import "C"

import (
	"fmt"
	"os"
	"strings"
	"unsafe"

	"github.com/superwhiskers/fennel"
	"github.com/superwhiskers/fennel/formats/xmls"
)

var noError = C.struct_fennel_Error{
	Type: C.fennel_ErrorTypeNone,
}

//export fennel_newAccountServerClient
func fennel_newAccountServerClient(accountServer *C.char, certificate *C.char, certificateLength C.int, key *C.char, keyLength C.int, clientInfo C.struct_fennel_ClientInformation, errorPointer **C.struct_fennel_Error) C.fennel_AccountServerClient {

	client, err := fennel.NewAccountServerClient(gostring(accountServer), C.GoBytes(unsafe.Pointer(certificate), certificateLength), C.GoBytes(unsafe.Pointer(key), keyLength), fennel.ClientInformation{
		ClientID:     gostring(clientInfo.ClientID),
		ClientSecret: gostring(clientInfo.ClientSecret),
		DeviceCert:   gostring(clientInfo.DeviceCert),
		Environment:  gostring(clientInfo.Environment),
		Country:      gostring(clientInfo.Country),
		Region:       gostring(clientInfo.Region),
		SysVersion:   gostring(clientInfo.SysVersion),
		Serial:       gostring(clientInfo.Serial),
		DeviceID:     gostring(clientInfo.DeviceID),
		DeviceType:   gostring(clientInfo.DeviceType),
		PlatformID:   gostring(clientInfo.PlatformID),
	})

	if err != nil {

		reassignError(errorPointer, convertGoErrorToError(err))
		return C.fennel_AccountServerClient(uintptr(0x00))

	}

	reassignError(errorPointer, noError)
	return C.fennel_AccountServerClient(unsafe.Pointer(client))

}

//export fennel_accountServerClient_doesUserExist
func fennel_accountServerClient_doesUserExist(clientPtr *C.fennel_AccountServerClient, nnid *C.char, errorPointer **C.struct_fennel_Error) C.int {

	client := convertPointerToAccountServerClient(clientPtr)

	exists, xml, err := client.DoesUserExist(gostring(nnid))
	if err != nil {

		reassignError(errorPointer, convertGoErrorToError(err))
		return C.int(-1)

	} else if len(xml.Errors) != 0 {

		reassignError(errorPointer, convertErrorXMLToError(xml))
		return C.int(-1)

	}

	reassignError(errorPointer, noError)
	if exists {

		return C.int(1)

	}
	return C.int(0)

}

// converts a pointer to an account server client
func convertPointerToAccountServerClient(ptr *C.fennel_AccountServerClient) *fennel.AccountServerClient {

	if uintptr(unsafe.Pointer(ptr)) == 0x00 {

		fmt.Println("fatal: this program has attempted to use an uninitialized AccountServerClient. stopping...")
		os.Exit(1)

	}
	return (*fennel.AccountServerClient)(unsafe.Pointer(ptr))

}

// reassigns an error pointer an error
func reassignError(eptr **C.struct_fennel_Error, e C.struct_fennel_Error) {

	C.free(unsafe.Pointer(*eptr))
	*eptr = (*C.struct_fennel_Error)(C.malloc(C.size_t(unsafe.Sizeof(e))))
	**eptr = e

}

// converts an errorxml to a fennel error type
func convertErrorXMLToError(exml xmls.ErrorXML) (e C.struct_fennel_Error) {

	e = C.struct_fennel_Error{
		Type: C.fennel_ErrorTypeErrorXML,
		ErrorXML: C.struct_fennel_ErrorXML{
			Cause:   cstring(exml.Errors[0].Cause),
			Code:    cstring(exml.Errors[0].Code),
			Message: cstring(exml.Errors[0].Message),
		},
	}

	return

}

// converts a go error type to a fennel error type
func convertGoErrorToError(err error) C.struct_fennel_Error {

	return C.struct_fennel_Error{
		Type:  C.fennel_ErrorTypeError,
		Error: cstring(err.Error()),
	}

}

// converts a c string to a go string type
func gostring(str *C.char) string {

	return strings.TrimSuffix(C.GoString(str), "\x00")

}

// converts a go string type to a c string type
func cstring(str string) *C.char {

	return C.CString(str)

}

func main() {}
