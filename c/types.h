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

// account server client type used as an argument that points to the account server client
typedef void *fennel_AccountServerClient;

// nintendo network error xml structure
struct fennel_ErrorXML {
  char *Cause;
  char *Code;
  char *Message;
};

// enumeration for the type of an error
enum fennel_ErrorType { fennel_ErrorTypeErrorXML, fennel_ErrorTypeError, fennel_ErrorTypeNone };

// error structure
struct fennel_Error {
  enum fennel_ErrorType Type;
  char *Error;
  struct fennel_ErrorXML ErrorXML;
};

// client information structure
struct fennel_ClientInformation {
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

