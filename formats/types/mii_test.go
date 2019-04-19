package types

import (
	"testing"
	"encoding/base64"
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
