/*

fennel - nintendo network utility library for golang
Copyright (C) 2018 superwhiskers <whiskerdev@protonmail.com>

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

package types

import (
	"unicode/utf8"
	
	"github.com/superwhiskers/fennel/utils"
)

// Mii contains all of the data that a mii can have
type Mii struct {
	// unknown fields
	Unknown1  uint64
	Unknown2  uint64
	Unknown3  uint64
	Unknown4  byte
	Unknown5  []byte
	Unknown6  byte
	Unknown7  byte
	Unknown8  uint8
	Unknown9  byte
	Unknown10 []byte

	// attributes
	BirthPlatform uint64
	FontRegion    uint64
	RegionMove    uint64
	Copyable      bool
	MiiVersion    uint8
	AuthorID      []uint8
	MiiID         []uint8
	LocalOnly     bool
	Color         uint64
	BirthDay      uint64
	BirthMonth    uint64
	Gender        byte
	MiiName       string
	Size          uint8
	Fatness       uint8
	AuthorName    string
	Checksum      uint16

	// face
	BlushType []byte
	FaceStyle []byte
	FaceColor []byte
	FaceType  []byte

	// hair
	HairMirrored bool
	HairColor    []byte
	HairType     uint8

	// eyes
	EyeThickness []byte
	EyeScale     []byte
	EyeColor     []byte
	EyeType      []byte
	EyeHeight    []byte
	EyeDistance  []byte
	EyeRotation  []byte

	// eyebrow
	EyebrowThickness []byte
	EyebrowScale     []byte
	EyebrowColor     []byte
	EyebrowType      []byte
	EyebrowHeight    []byte
	EyebrowDistance  []byte
	EyebrowRotation  []byte

	// nose
	NoseHeight []byte
	NoseScale  []byte
	NoseType   []byte

	// mouth
	MouthThickness []byte
	MouthScale     []byte
	MouthColor     []byte
	MouthType      []byte
	MouthHeight    []byte

	// mustache
	MustacheType   []byte
	MustacheHeight []byte
	MustacheScale  []byte

	// beard
	BeardColor []byte
	BeardType  []byte

	// glasses
	GlassesHeight []byte
	GlassesScale  []byte
	GlassesColor  []byte
	GlassesType   []byte

	// mole
	MoleY       []byte
	MoleX       []byte
	MoleScale   []byte
	MoleEnabled bool
}

// NilMii is a Mii with no data
var NilMii = Mii{}

// swaps the endianness of a mii binary format to little-endian
func swapMiiEndiannessToLittle(data []byte) []byte {

	data = utils.Swapu32Little(data, 0x00)

	for i := 0x18; i <= 0x2E; i += 2 {

		data = utils.Swapu16Little(data, i)

	}

	for i := 0x30; i <= 0x5C; i += 2 {

		data = utils.Swapu16Little(data, i)

	}

	data = utils.Swapu16Little(data, 0x5C)

	return data

}

// ParseMii takes a mii as a byte array and parses it to a Mii
func ParseMii(miiByte []byte) Mii {

	buf := utils.NewByteBuffer(swapMiiEndiannessToLittle(miiByte))
	btf := utils.NewBitfield(miiByte...)
	mii := Mii{}

	mii.BirthPlatform = btf.ReadBitsNext(4) // 0x00
	mii.Unknown1 = btf.ReadBitsNext(4)      // 0x00.4
	mii.Unknown2 = btf.ReadBitsNext(4)      // 0x01
	mii.Unknown3 = btf.ReadBitsNext(4)      // 0x01.4
	mii.FontRegion = btf.ReadBitsNext(4)    // 0x02
	mii.RegionMove = btf.ReadBitsNext(2)    // 0x02.4
	mii.Unknown4 = btf.ReadBitNext()        // 0x02.6
	if btf.ReadBitNext() == 0x00 {          // 0x02.7

		mii.Copyable = false

	} else {

		mii.Copyable = true

	}

	// seek the byte buffer to offset 0x03 to align it with the bitfield's offset
	buf.Seek(0x03, false)

	mii.MiiVersion = buf.ReadBytesNext(1)[0] // 0x03
	mii.AuthorID = buf.ReadBytesNext(8)      // 0x04
	mii.MiiID = buf.ReadBytesNext(10)        // 0x0C
	mii.Unknown5 = buf.ReadBytesNext(2)      // 0x16

	// seek the bitfield to offset 0x16*8 to align it with the byte buffer's offset
	// (the bitfield takes offsets in terms of bits)
	btf.Seek(0x18*8, false) // TODO: verify this offset

	mii.Unknown6 = btf.ReadBitNext()     // 0x16
	mii.Unknown7 = btf.ReadBitNext()     // 0x16.1
	mii.Color = btf.ReadBitsNext(4)      // 0x16.2
	mii.BirthDay = btf.ReadBitsNext(5)   // 0x16.6
	mii.BirthMonth = btf.ReadBitsNext(4) // 0x17.3
	mii.Gender = btf.ReadBitNext()       // 0x17.7

	// seek the byte buffer back to the proper spot
	buf.Seek(0x18, false)

	// TODO: optimize this for speed and make sure it works
	tmp := buf.ReadBytesNext(20)
	tmp2 := []rune{}
	for len(tmp) > 0 {

		r, size := utf8.DecodeRune(tmp)
		tmp2 = append(tmp2, r)
		tmp = tmp[size:]

	}
	mii.MiiName = string(tmp2)

	return mii

}
