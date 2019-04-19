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

package fennel

import (
	"crypto/tls"
	"encoding/base64"
	"testing"

	"github.com/valyala/fasthttp"
)

var ctrCommonCert = []byte(`
subject=/C=US/ST=Washington/L=Redmond/O=Nintendo of America, Inc./OU=IS/CN=CTR Common Prod 1/emailAddress=ca@noa.nintendo.com
issuer=/C=US/ST=Washington/O=Nintendo of America Inc./OU=IS/CN=Nintendo CA - G3
-----BEGIN CERTIFICATE-----
MIIEwzCCA6ugAwIBAgIBBjANBgkqhkiG9w0BAQsFADBtMQswCQYDVQQGEwJVUzET
MBEGA1UECBMKV2FzaGluZ3RvbjEhMB8GA1UEChMYTmludGVuZG8gb2YgQW1lcmlj
YSBJbmMuMQswCQYDVQQLEwJJUzEZMBcGA1UEAxMQTmludGVuZG8gQ0EgLSBHMzAe
Fw0xMDA1MTMxOTE5NDZaFw0zNzEyMjIxOTE5NDZaMIGlMQswCQYDVQQGEwJVUzET
MBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEiMCAGA1UEChMZ
TmludGVuZG8gb2YgQW1lcmljYSwgSW5jLjELMAkGA1UECxMCSVMxGjAYBgNVBAMT
EUNUUiBDb21tb24gUHJvZCAxMSIwIAYJKoZIhvcNAQkBFhNjYUBub2EubmludGVu
ZG8uY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA81Vzs324jZwc
NpbFESgDNooVTRP1TlxvYwz8bbHnJHhImjEJNO29YSTpjmF7wonczooeKXfE/Ry2
+ey9mk92UhzSnvuSHQ6P2zFBbcPnE8eBi73oDnErgixiWe1TKP1G5LvwOqrEkVmX
LN/qnLrsfFp4QNyFc+PLvJ9IAfRSBwdRJHAiSgE9nB9eI7AGcM6DCw7+p9zEz6rN
RHUVRc5I132wJpQa8aoWaqPW7LE8exEC3VSfDHRVPjZUMRhfoBVSi2NfiA3xYsqk
v+Ct3E+bzW8y1aAQ7wIshQ/RGcLtVZE+tkoAznXewVLdKtcC67Vy4awhJ/BqK1tv
c26qV3zIJwIDAQABo4IBMzCCAS8wCQYDVR0TBAIwADAsBglghkgBhvhCAQ0EHxYd
T3BlblNTTCBHZW5lcmF0ZWQgQ2VydGlmaWNhdGUwHQYDVR0OBBYEFIzG7XO5Ojx2
G45r5dTszWF1rcFtMIGXBgNVHSMEgY8wgYyAFATT3tP98MjrwlmSh/sf1z5y+O35
oXGkbzBtMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEhMB8GA1UE
ChMYTmludGVuZG8gb2YgQW1lcmljYSBJbmMuMQswCQYDVQQLEwJJUzEZMBcGA1UE
AxMQTmludGVuZG8gQ0EgLSBHM4IBATA7BgNVHR8ENDAyMDCgLqAshipodHRwOi8v
Y3JsLm5pbnRlbmRvLmNvbS9uaW50ZW5kby1jYS1nMy5jcmwwDQYJKoZIhvcNAQEL
BQADggEBAEOXZ/3IkNuFUfdxHpP0vrcSCTnDqMk8gsLVbN39BJT8Wqm8e3MFNhS/
Y1YOWgoIPtJp4cd2tXM3cXWzUZgm3SKd1XX/B81PFLEYlk+metUqB4jpF0ApCZs6
RNoXDBTx6XzsC07CA3uaxEdeWjC5Nl29AHuZ1YC/Z+7Da57TwBaa+/APj4y5mGUa
ahbvwpe1t3GSNOS5nBDSeCHAKLmzfnXpliA5qQZxo94RSXIVWK8hilXoFDQCL904
OGpgZnAhz4p3rcJYTq9ub8n6NYr9OJKKbWXfJY1QK4pXFVcIuAph0o/EyzDIEXuT
J4Q4b2km8uI0H4yxsQwUX9Epw6Vbujc=
-----END CERTIFICATE-----
`)

var ctrCommonKey = []byte(`
-----BEGIN PRIVATE KEY-----
MIIEwAIBADANBgkqhkiG9w0BAQEFAASCBKowggSmAgEAAoIBAQDzVXOzfbiNnBw2
lsURKAM2ihVNE/VOXG9jDPxtseckeEiaMQk07b1hJOmOYXvCidzOih4pd8T9HLb5
7L2aT3ZSHNKe+5IdDo/bMUFtw+cTx4GLvegOcSuCLGJZ7VMo/Ubku/A6qsSRWZcs
3+qcuux8WnhA3IVz48u8n0gB9FIHB1EkcCJKAT2cH14jsAZwzoMLDv6n3MTPqs1E
dRVFzkjXfbAmlBrxqhZqo9bssTx7EQLdVJ8MdFU+NlQxGF+gFVKLY1+IDfFiyqS/
4K3cT5vNbzLVoBDvAiyFD9EZwu1VkT62SgDOdd7BUt0q1wLrtXLhrCEn8GorW29z
bqpXfMgnAgMBAAECggEBAMFOTib2JgmhTax0I8OYVM0b7wYXZ9XDit1WMKZ4INaR
E6QidlzszHiC2WO5v5Zw7M/LW2C3++7Tw+xRjOIsZCOhMBUKZy3cJp4LyB2J9mV5
JUm9KL9oWhcEaXFlHp4+bvZA8vu4M4YAdR86FuhBeqLjQArO5NmGypBivNKIpC1d
rwzSMyPddZvur7AsTIK0Ym9SwWN9eK7F2uBkzAneOugOTEAhq2ZnPMByNtjpTvw/
nvgNAB4Ukz/oomtleCaw92SoSlQlYGzVmuhvt4QqOQzS1V+ToauUhmAPmWEHF3IJ
yL1SiY7UmMlqQoGFV6IH4cwLiwm1wk/IF1tkfvR0gmECgYEA/vhTJYn/Gd3XY2tm
vJs5pvmkJJgNxqgunOCFZxrGxf4BLvQGkRnBlT4MzFa1J8ZNtU6VOVbQsS7cD0ke
2Kyv5rdt/7K2/bNdG4Ocnd+GcWpEtF+ik2wEQB48FQvYSQn/w5NKyJlyYOhXKGe3
XhTiyV5rCwiLfRP8pxAO0xLzVhcCgYEA9FEX+hlxYdg2bVfx+geI9uW0COLi7kzc
xxlPjBvdFltYsltSNLP+0BNqHfp3G5PW/XMTAdZnS9033gy4jLU6fAge8FcnLybh
GrpVNCnMY+DLRyVtlu1fcKKH59BK/ElR7aamxqJyw1JO2Z35u251nIbmrMgs5Txp
8aZ4HXCveHECgYEA1mRybddKhTKPwU53FdKkOK4jgo3Ez61tfIYiRl8ykxuRXSze
NLZmm5qQYmXqb+aEQxcvzQYd907CxaujX2hdhG/q854P1uYyPUd+sxVYVBeaa90a
tEGYlV2XAc9y73+T65z3vhOhJLFZUGVdv6NqSw60jZOCzwq2YLfU71E5AcMCgYEA
jQ6c90rlSYaZtfvGu4LKMzJgBZlpSAicl18nrE8SEKxgw2kyRzd88QmkhPZs+kEb
KW3dFXyCWyy36r4RdzvTLnVJ152aBAFAijv2oY1YcnoBI2yanz8hkVhlexOpl4uF
f95t/9UeyWKmH8KzwuF9igfg+vT/5sJAsMJaKzU6OiECgYEAlKK+RKTYwquD+3gK
o1gsGI+nR96Cb1kvfXzsj+V5UkZchew2pOqhrqpPknGIlFCeTDYjN8jqJyX4EljJ
1FTegfhfSe+XR7KOIh8b+d+fgftyRIp3M//BUF1FtwL789f/VakaIkz6Ret/+8tA
3UHqKKGtgSRL9kiTMVYy64pBG9A=
-----END PRIVATE KEY-----
`)

var clientInfo = ClientInformation{
	ClientID:     "ea25c66c26b403376b4c5ed94ab9cdea",
	ClientSecret: "d137be62cb6a2b831cad8c013b92fb55",
	DeviceCert:   "",
	Environment:  "L1",
	Country:      "US",
	Region:       "2",
	SysVersion:   "1111",
	Serial:       "1",
	DeviceID:     "1",
	DeviceType:   "",
	PlatformID:   "1",
	FPDVersion:   "0000",
}

func TestNewAccountServerClient(t *testing.T) {

	keyPair, err := tls.X509KeyPair(ctrCommonCert, ctrCommonKey)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	expectedOutput := &AccountServerClient{
		APIEndpoint: "https://account.nintendo.net/v1/api",
		HTTPClient: &fasthttp.Client{
			TLSConfig: &tls.Config{
				Certificates:       []tls.Certificate{keyPair},
				ClientAuth:         tls.RequireAndVerifyClientCert,
				InsecureSkipVerify: true,
			},
		},
		ClientInformation: clientInfo,
	}

	output, err := NewAccountServerClient("https://account.nintendo.net/v1/api", ctrCommonCert, ctrCommonKey, clientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	if output.ClientInformation != expectedOutput.ClientInformation {

		t.Errorf("invalid output")

	}

}

func TestDoesUserExist(t *testing.T) {

	client, err := NewAccountServerClient("https://account.nintendo.net/v1/api", ctrCommonCert, ctrCommonKey, clientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	output, exml, err := client.DoesUserExist("scott0852")
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	if len(exml.Errors) != 0 {

		t.Errorf("expected no error to occur, instead got %v\n", exml.Errors[0])

	}

	if output != true {

		t.Errorf("invalid output")

	}

}

func TestGetEULA(t *testing.T) {

	client, err := NewAccountServerClient("https://account.nintendo.net/v1/api", ctrCommonCert, ctrCommonKey, clientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	output, exml, err := client.GetEULA("US", "@latest")
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	if len(exml.Errors) != 0 {

		t.Errorf("expected no error to occur, instead got %v\n", exml.Errors[0])

	}

	for _, agreement := range output.Agreements {

		if agreement.Language == "en" && agreement.Title.Data == "Nintendo Network Services Agreement" {

			return

		}

	}

	t.Errorf("invalid output")

}

func TestGetPIDs(t *testing.T) {

	client, err := NewAccountServerClient("https://account.nintendo.net/v1/api", ctrCommonCert, ctrCommonKey, clientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	output, exml, err := client.GetPIDs([]string{"scott0852"})
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	if len(exml.Errors) != 0 {

		t.Errorf("expected no error to occur, instead got %v\n", exml.Errors[0])

	}

	if output[0] != 1794841894 {

		t.Errorf("invalid output")

	}

}

func TestGetNNIDs(t *testing.T) {

	client, err := NewAccountServerClient("https://account.nintendo.net/v1/api", ctrCommonCert, ctrCommonKey, clientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	output, exml, err := client.GetNNIDs([]int64{1794841894})
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	if len(exml.Errors) != 0 {

		t.Errorf("expected no error to occur, instead got %v\n", exml.Errors[0])

	}

	if output[0] != "SCOTT0852" {

		t.Errorf("invalid output")

	}

}

func TestGetMiis(t *testing.T) {

	expectedOutput := []uint8{0xd0, 0xea, 0x35, 0x97, 0xcc, 0x40, 0xfd, 0x77}

	client, err := NewAccountServerClient("https://account.nintendo.net/v1/api", ctrCommonCert, ctrCommonKey, clientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	output, exml, err := client.GetMiis([]int64{1794841894})
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	if len(exml.Errors) != 0 {

		t.Errorf("expected no error to occur, instead got %v\n", exml.Errors[0])

	}

	for i, b := range output.Miis[0].Mii.AuthorID {

		if expectedOutput[i] != b {

			t.Errorf("invalid output")

		}

	}

}

// temporary test function, will be removed later
func TestVerifyMii(t *testing.T) {

	client, err := NewAccountServerClient("https://account.nintendo.net/v1/api", ctrCommonCert, ctrCommonKey, clientInfo)
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	output, _, err := client.GetMiis([]int64{1794841894})
	if err != nil {

		t.Errorf("expected no error to occur, instead got %v\n", err)

	}

	t.Logf("%#v", output.Miis[0].Data)

	if output.Miis[0].Data != base64.StdEncoding.EncodeToString(output.Miis[0].Mii.Encode()) {

		t.Logf("original data: %#v", output.Miis[0].Data)
		t.Logf("encoded data: %#v", base64.StdEncoding.EncodeToString(output.Miis[0].Mii.Encode()))
		t.Errorf("encoded Mii != original data")

	}

}
