"""

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

"""

import ctypes, sys

# ClientInformation is an implementation of fennel's ClientInformation struct in python
class ClientInformation(ctypes.Structure):
    _fields_ = [
        ("ClientID", ctypes.c_char_p),
        ("ClientSecret", ctypes.c_char_p),
        ("DeviceCert", ctypes.c_char_p),
        ("Environment", ctypes.c_char_p),
        ("Country", ctypes.c_char_p),
        ("Region", ctypes.c_char_p),
        ("SysVersion", ctypes.c_char_p),
        ("Serial", ctypes.c_char_p),
        ("DeviceID", ctypes.c_char_p),
        ("DeviceType", ctypes.c_char_p),
        ("PlatformID", ctypes.c_char_p),
    ]

    def __init__(
        self,
        client_id=bytes("ea25c66c26b403376b4c5ed94ab9cdea", "utf8"),
        client_secret=bytes("d137be62cb6a2b831cad8c013b92fb55", "utf8"),
        device_cert=bytes("", "utf8"),
        environment=bytes("", "utf8"),
        country=bytes("US", "utf8"),
        region=bytes("2", "utf8"),
        system_version=bytes("1111", "utf8"),
        serial=bytes("1", "utf8"),
        device_id=bytes("1", "utf8"),
        device_type=bytes("", "utf8"),
        platform_id=bytes("1", "utf8"),
    ):
        self.ClientID = client_id
        self.ClientSecret = client_secret
        self.DeviceCert = device_cert
        self.Environment = environment
        self.Country = country
        self.Region = region
        self.SysVersion = system_version
        self.Serial = serial
        self.DeviceID = device_id
        self.DeviceType = device_type
        self.PlatformID = platform_id


# load fennel
fennel = ctypes.CDLL("../fennel.so")

# create a clientinformation struct
clientInfo = ClientInformation()

newClient = fennel.fennel_newClient
newClient.argtypes = [
    ctypes.c_char_p,
    ctypes.c_char_p,
    ctypes.c_char_p,
    ClientInformation,
]
newClient.restype = ctypes.c_void_p

doesUserExist = fennel.fennel_doesUserExist
doesUserExist.argtypes = [ctypes.c_void_p, ctypes.c_char_p]
doesUserExist.restype = ctypes.c_int

client = newClient(
    bytes("https://account.nintendo.net/v1/api", "utf8"),
    bytes("keypair/ctr-common-cert.pem", "utf8"),
    bytes("keypair/ctr-common-key.pem", "utf8"),
    clientInfo,
)

if doesUserExist(client, bytes(sys.argv[1], "utf8")) == 1:
    print("yes, the user does exist")
else:
    print("no, the user does not exist")
