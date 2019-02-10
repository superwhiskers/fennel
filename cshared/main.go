package main

/*
struct ClientInformation {
	char *ClientID;
	char *ClientSecret;
	char *DeviceCert;
	char *Environment;
	char *Country;
	char *Region;
	char *SysVersion;
	char *Serial;
	char *DeviceID;
	char *DeviceType;
	char *PlatformID;
};

typedef void* Client;
*/
import "C"

import (
	"unsafe"
	"strings"

	"github.com/superwhiskers/fennel"
)

//export fennel_newClient
func fennel_newClient(accountServer, certificatePath, keyPath *C.char, clientInfo C.struct_ClientInformation) C.Client {

	client, err := fennel.NewAccountServerClient(gostring(accountServer), gostring(certificatePath), gostring(keyPath), fennel.ClientInformation{
		ClientID: gostring(clientInfo.ClientID),
		ClientSecret: gostring(clientInfo.ClientSecret),
		DeviceCert: gostring(clientInfo.DeviceCert),
		Environment: gostring(clientInfo.Environment),
		Country: gostring(clientInfo.Country),
		Region: gostring(clientInfo.Region),
		SysVersion: gostring(clientInfo.SysVersion),
		Serial: gostring(clientInfo.Serial),
		DeviceID: gostring(clientInfo.DeviceID),
		DeviceType: gostring(clientInfo.DeviceType),
		PlatformID: gostring(clientInfo.PlatformID),
	})
	if err != nil {

		panic(err)

	}

	return C.Client(unsafe.Pointer(client))

}

//export fennel_doesUserExist
func fennel_doesUserExist(clientPtr C.Client, nnid *C.char) C.int {

	client := convertPointerToClient(clientPtr)

	exists, xml, err := client.DoesUserExist(gostring(nnid))
	if err != nil {

		panic(err)

	} else if len(xml.Errors) != 0 {

		panic(xml.Errors[0])

	}

	if exists {

		return C.int(1)

	} else {

		return C.int(0)	

	}

}

func convertPointerToClient(ptr C.Client) *fennel.AccountServerClient {

	return (*fennel.AccountServerClient)(unsafe.Pointer(ptr))

}

func gostring(str *C.char) string {

	return strings.TrimSuffix(C.GoString(str), "\x00")

}

func main() {}
