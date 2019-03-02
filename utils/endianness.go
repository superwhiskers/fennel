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
