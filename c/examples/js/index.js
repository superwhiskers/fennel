/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

let ffi = require('ffi')
let ref = require('ref')
let Struct = require('ref-struct')

let clientType = ref.refType(ref.types.void)
let clientInfoStruct = Struct({
    'ClientID': 'string',
    'ClientSecret': 'string',
    'DeviceCert': 'string',
    'Environment': 'string',
    'Country': 'string',
    'Region': 'string',
    'SysVersion': 'string',
    'Serial': 'string',
    'DeviceID': 'string',
    'DeviceType': 'string',
    'PlatformID': 'string'
})

let fennel = ffi.Library('../fennel.so', {
    'fennel_newClient': [ clientType, [ 'string', 'string', 'string', clientInfoStruct ]],
    'fennel_doesUserExist': [ 'int', [ clientType, 'string' ]]
})

let clientInfo = new clientInfoStruct()
clientInfo.ClientID = "ea25c66c26b403376b4c5ed94ab9cdea"
clientInfo.ClientSecret = "d137be62cb6a2b831cad8c013b92fb55"
clientInfo.DeviceCert = ""
clientInfo.Environment = ""
clientInfo.Country = "US"
clientInfo.Region = "2"
clientInfo.SysVersion = "1111"
clientInfo.Serial = "1"
clientInfo.DeviceID = "1"
clientInfo.DeviceType = ""
clientInfo.PlatformID = "1"

let client = fennel.fennel_newClient("https://account.nintendo.net/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", clientInfo)

if (fennel.fennel_doesUserExist(client, process.argv[2]) === 0) {

    console.log("no, that user does not exist")

} else {

    console.log("yes, that user does exist")

}
