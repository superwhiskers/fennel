/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

package types

import "testing"

func TestServicetoken(t *testing.T) {

	var (
		expectedServicetoken       = "3YysZBU8Xm+Uqb60uTU69cksB4SrTGcVJtU9JlzJhMYcbnpBq25KgWlrPt18zOfa+JHJzF6Ha36NuGjQ+BCrRsfRmnGSMz5muA7GLV195inaPza2AQLoEdRp4qjsklle1oqswldQS8m3pAHpZAHYYOORoJe3KLF6uEvleJRiCfk="
		expectedParsedServicetoken = "dd8cac64153c5e6f94a9beb4b9353af5c92c0784ab4c671526d53d265cc984c61c6e7a41ab6e4a81696b3edd7ccce7daf891c9cc5e876b7e8db868d0f810ab46c7d19a7192333e66b80ec62d5d7de629da3f36b60102e811d469e2a8ec92595ed68aacc257504bc9b7a401e96401d860e391a097b728b17ab84be578946209f9"
	)

	servicetoken, err := ParseServicetoken(expectedServicetoken)

	if err != nil {

		t.Errorf("couldn't parse the servicetoken to a safe format. error: %v\n", err)

	}

	if servicetoken != expectedParsedServicetoken {

		t.Errorf("the parsed servicetoken doesn't match the expected one")

	}

	t.Logf("got parsed servicetoken: %s\n", servicetoken)

}
