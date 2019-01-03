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
)

//export libninty_newClient
func libninty_newClient(accountServer, certificatePath, keyPath string, clientInfo C.struct_ClientInformation) C.Client {

	Client, err := libninty.NewClient(accountServer, certificatePath, keyPath, libninty.ClientInformation{
		ClientID: C.GoString(clientInfo.ClientID),
		ClientSecret: C.GoString(clientInfo.ClientSecret),
		DeviceCert: C.GoString(clientInfo.DeviceCert),
		Environment: C.GoString(clientInfo.Environment),
		Country: C.GoString(clientInfo.Country),
		Region: C.GoString(clientInfo.Region),
		SysVersion: C.GoString(clientInfo.SysVersion),
		Serial: C.GoString(clientInfo.Serial),
		DeviceID: C.GoString(clientInfo.DeviceID),
		DeviceType: C.GoString(clientInfo.DeviceType),
		PlatformID: C.GoString(clientInfo.PlatformID),
	})
	if err != nil {

		panic(err)

	}

	return C.Client(unsafe.Pointer(Client))

}

//export libninty_doesUserExist
func libninty_doesUserExist(clientPtr C.Client, nnid string) int {

	client := convertPointerToClient(clientPtr)

	exists, xml, err := client.DoesUserExist(nnid)
	if err != nil {

		panic(err)

	} else if len(xml.Errors) != 0 {

		panic(xml.Errors[0])

	}

	if exists {

		return 1

	} else {

		return 0	

	}

}

func convertPointerToClient(ptr C.Client) *libninty.Client {

	return (*libninty.Client)(unsafe.Pointer(ptr))

}

func main() {}
