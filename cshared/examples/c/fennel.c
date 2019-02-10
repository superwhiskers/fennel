// fennel.c - test code for fennel.so
#include <stdio.h>
#include "fennel.h"

int main(int argc, char *argv[]) {

  struct ClientInformation clientInfo;
  clientInfo.ClientID = "ea25c66c26b403376b4c5ed94ab9cdea";
  clientInfo.ClientSecret = "d137be62cb6a2b831cad8c013b92fb55";
  clientInfo.DeviceCert = "";
  clientInfo.Environment = "";
  clientInfo.Country = "US";
  clientInfo.Region = "2";
  clientInfo.SysVersion = "1111";
  clientInfo.Serial = "1";
  clientInfo.DeviceID = "1";
  clientInfo.DeviceType = "";
  clientInfo.PlatformID = "1";

  Client client;
  client = fennel_newClient("https://account.nintendo.net/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", clientInfo);

  if (fennel_doesUserExist(client, argv[1]) == 0) {

    printf("no, the user does not exist\n");

  } else {

    printf("yes, the user does exist\n");

  }
  
}
