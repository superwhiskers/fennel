// libninty.js - test code for libninty.so
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

let libninty = ffi.Library('./libninty.so', {
    'libninty_newClient': [ clientType, [ 'string', 'string', 'string', clientInfoStruct ]],
    'libninty_doesUserExist': [ 'int', [ clientType, 'string' ]]
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

let client = libninty.libninty_newClient("https://account.nintendo.net/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", clientInfo)

if (libninty.libninty_doesUserExist(client, process.argv[2]) === 0) {

    console.log("no, that user does not exist")

} else {

    console.log("yes, that user does exist")

}
