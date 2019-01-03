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

	"github.com/superwhiskers/libninty"
	"github.com/superwhiskers/libninty/formats/xmls"
)

//export libninty_newClient
func libninty_newClient(accountServer, certificatePath, keyPath string, clientInfo C.struct_ClientInformation) C.Client {

	Client := libninty.NewClient(accountServer, certificatePath, keyPath, libninty.ClientInformation{
		ClientID: C.GoString(clientInfo._ClientID),
		ClientSecret: C.GoString(clientInfo._ClientSecret),
		DeviceCert: C.GoString(clientInfo._DeviceCert),
		Environment: C.GoString(clientInfo._Environment),
		Country: C.GoString(clientInfo._Country),
		Region: C.GoString(clientInfo._Region),
		SysVersion: C.GoString(clientInfo._SysVersion),
		Serial: C.GoString(clientInfo._Serial),
		DeviceID: C.GoString(clientInfo._DeviceID),
		DeviceType: C.GoString(clientInfo._DeviceType),
		PlatformID: C.GoString(clientInfo._PlatformID),
	})

	return C.Client(unsafe.Pointer(Client))

}

//export libninty_doesUserExist
func libninty_doesUserExist(clientPtr C.Client, nnid string) int {

	client := convertPointerToClient(clientPtr)

	exists, xml, err := client.DoesUserExist(nnid)
	if err != nil {

		panic(err)

	} else if xml != xmls.NilErrorXML {

		panic(xml)

	}

	return int(exists)

}

func convertPointerToClient(ptr C.Client) *libninty.Client {

	return *libninty.Client(unsafe.Pointer(ptr))

}

func main() {}
