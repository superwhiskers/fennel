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

