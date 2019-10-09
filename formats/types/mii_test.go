/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

package types

import (
	"encoding/base64"
	"testing"
)

var data = "AwEAQNDqNZfMQP131K+wv1n8kW4jgAAApltTAEMATwBUAFQAMAA4ADUAMgAAAGc5AgA5B7RIRBL3IsQGrQwTagwAOCmxMUhQUwBjAG8AdAB0ACAATQAuAAAAAAAAAA50"

func BenchmarkParseMii(b *testing.B) {

	b.ReportAllocs()

	data, err := base64.StdEncoding.DecodeString(data)
	if err != nil {

		b.Fatalf("expected no error to occur, instead got %v\n", err)

	}

	var out *Mii
	for n := 0; n < b.N; n++ {

		out = ParseMii(data)

	}

	_ = out

}
