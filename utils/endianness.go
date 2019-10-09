/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

package utils

import "encoding/binary"

// Swapu64Big converts part of a byte slice from little-endian uint64 to big-endian uint64 starting at the specified offset, and covering eight indexes
func Swapu64Big(data []byte, offset int) []byte {

	byteSection := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.BigEndian.PutUint64(byteSection, binary.LittleEndian.Uint64(data[offset:offset+8]))
	return ApplyByteSliceAtOffset(byteSection, data, offset)

}

// Swapu64Little converts part of a byte slice from big-endian uint64 to little-endian uint64 starting at the specified offset, and covering eight indexes
func Swapu64Little(data []byte, offset int) []byte {

	byteSection := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.LittleEndian.PutUint64(byteSection, binary.BigEndian.Uint64(data[offset:offset+8]))
	return ApplyByteSliceAtOffset(byteSection, data, offset)

}

// Swapu32Big converts part of a byte slice from little-endian uint32 to big-endian uint32 starting at the specified offset, and covering four indexes
func Swapu32Big(data []byte, offset int) []byte {

	byteSection := []byte{0, 0, 0, 0}
	binary.BigEndian.PutUint32(byteSection, binary.LittleEndian.Uint32(data[offset:offset+4]))
	return ApplyByteSliceAtOffset(byteSection, data, offset)

}

// Swapu32Little converts part of a byte slice from big-endian uint32 to little-endian uint32 starting at the specified offset, and covering four indexes
func Swapu32Little(data []byte, offset int) []byte {

	byteSection := []byte{0, 0, 0, 0}
	binary.LittleEndian.PutUint32(byteSection, binary.BigEndian.Uint32(data[offset:offset+4]))
	return ApplyByteSliceAtOffset(byteSection, data, offset)

}

// Swapu16Big converts part of a byte slice from little-endian uint16 to big-endian uint16 starting at the specified offset, and covering two indexes
func Swapu16Big(data []byte, offset int) []byte {

	byteSection := []byte{0, 0}
	binary.BigEndian.PutUint16(byteSection, binary.LittleEndian.Uint16(data[offset:offset+2]))
	return ApplyByteSliceAtOffset(byteSection, data, offset)

}

// Swapu16Little converts part of a byte slice from big-endian uint16 to little-endian uint16 starting at the specified offset, and covering two indexes
func Swapu16Little(data []byte, offset int) []byte {

	byteSection := []byte{0, 0}
	binary.LittleEndian.PutUint16(byteSection, binary.BigEndian.Uint16(data[offset:offset+2]))
	return ApplyByteSliceAtOffset(byteSection, data, offset)

}
