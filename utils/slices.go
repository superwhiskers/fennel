/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

package utils

import (
	"strconv"
	"unicode/utf8"
)

// ApplyByteSliceAtOffset applies a byte slice to another byte slice at the specified offset, overwriting any existing indexes already in the slice
func ApplyByteSliceAtOffset(src, dest []byte, offset int) []byte {

	var (
		i = 0
		n = len(src)
	)
	{
	apply_loop:
		dest[offset+i] = src[i]
		i++
		if i < n {

			goto apply_loop

		}
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

	var (
		flag uint16

		i = 0
		n = len(data)
	)
	{
	sum_loop:
		i2 := 0
		{
		hash_add_loop:
			flag = hash & 0x8000
			hash = (hash << 1) & 0xFFFF
			if flag != 0x00 {

				hash ^= 0x1021

			}
			i2++
			if i2 < 8 {
				
				goto hash_add_loop

			}
		}
		hash ^= uint16(data[i])
		i++
		if i < n {

			goto sum_loop

		}
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
		s += string(r)

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
