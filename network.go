/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

package fennel

// ClientInformation holds data for headers sent in requests to nintendo network
type ClientInformation struct {
	/* (mostly) constant information */
	PlatformID   string // this is pretty much always "1"
	DeviceType   string // this is either 1 or 2. 1 being debug, 2 being retail
	ClientID     string // this is pretty much always "a2efa818a34fa16b8afbc8a74eba3eda"
	ClientSecret string // this is pretty much always "c91cdb5658bd4954ade78533a339cf9a"
	FPDVersion   string // this is pretty much always "0000"
	Environment  string // this is pretty much always "L1"

	/* fluctuating information */
	DeviceID   string // device id obtained from your console
	Serial     string // string representing the serial number of your console
	SysVersion string // unsigned integer representing your system version
	Region     string // 1 = JPN, 2 = USA, 4 = EUR, 8 = AUS, 16 = CHN, 32 = KOR, 64 = TWN
	Country    string // two-letter country code
	DeviceCert string // device certificate obtained from your console
}

// ApplicationInformation holds data about the application accessing the api
type ApplicationInformation struct {
	TitleID            uint64
	ApplicationVersion uint64
}
