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

package utils

import (
	"strconv"
	"unicode/utf8"
)

// ApplyByteSliceAtOffset applies a byte slice to another byte slice at the specified offset, overwriting any existing indexes already in the slice
func ApplyByteSliceAtOffset(src, dest []byte, offset int) []byte {

	for i, byt := range src {

		dest[offset+i] = byt

	}
	return dest

}

// ConvertInt64SliceToStringSlice converts a slice of int64s to a slice of strings
func ConvertInt64SliceToStringSlice(il []int64) (sl []string) {

	sl = []string{}
	for _, i := range il {

		sl = append(sl, strconv.Itoa(int(i)))

	}

	return

}

// CRC16 computes the crc16 checksum of a byte array
func CRC16(data []byte) (hash uint16) {

	var flag uint16
	hash = 0

	for _, c := range data {

		for i := 0; i < 8; i++ {

			flag = hash & 0x8000
			hash = (hash << 1) & 0xFFFF
			if flag != 0x00 {

				hash ^= 0x1021

			}

		}
		hash ^= uint16(c)

	}

	return

}

// DecodeUTF8StringFromBytes decodes a string from a byte array encoded in utf8
func DecodeUTF8StringFromBytes(b []byte) (s string) {

	s = ""
	var (
		r    rune
		size int
	)
	for len(b) > 0 {

		r, size = utf8.DecodeRune(b)
		b = b[size:]
		if r == 0 {

			continue

		}
		s = s + string(r)

	}
	return

}

// EncodeBytesFromUTF8String encodes a byte array in utf8 from a string
func EncodeBytesFromUTF8String(s string) (b []byte) {

	b = []byte{}
	var tb []byte
	for i := 0; i < len(s); i++ {

		tb = make([]byte, utf8.RuneLen(rune(s[i])))
		_ = utf8.EncodeRune(tb, rune(s[i]))
		b = append(b, tb...)
		
	}
	return

}
