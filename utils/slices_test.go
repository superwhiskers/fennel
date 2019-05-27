package utils

import "testing"

func BenchmarkCRC16(b *testing.B) {

	b.ReportAllocs()
	
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x010}

	for i := 0; i < b.N; i++ {

		_ = CRC16(data)

	}

}
