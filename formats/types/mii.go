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
	BlushType uint64
	FaceStyle uint64
	FaceColor uint64
	FaceType  uint64

	// hair
	HairMirrored bool
	HairColor    uint64
	HairType     uint8

	// eyes
	EyeThickness uint64
	EyeScale     uint64
	EyeColor     uint64
	EyeType      uint64
	EyeHeight    uint64
	EyeDistance  uint64
	EyeRotation  uint64

	// eyebrow
	EyebrowThickness uint64
	EyebrowScale     uint64
	EyebrowColor     uint64
	EyebrowType      uint64
	EyebrowHeight    uint64
	EyebrowDistance  uint64
	EyebrowRotation  uint64

	// nose
	NoseHeight uint64
	NoseScale  uint64
	NoseType   uint64

	// mouth
	MouthThickness uint64
	MouthScale     uint64
	MouthColor     uint64
	MouthType      uint64
	MouthHeight    uint64

	// mustache
	MustacheType   uint64
	MustacheHeight uint64
	MustacheScale  uint64

	// beard
	BeardColor uint64
	BeardType  uint64

	// glasses
	GlassesHeight uint64
	GlassesScale  uint64
	GlassesColor  uint64
	GlassesType   uint64

	// mole
	MoleY       uint64
	MoleX       uint64
	MoleScale   uint64
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

	for i := 0x30; i <= 0x48; i += 2 {

		data = utils.Swapu16Little(data, i)

	}

	for i := 0x48; i <= 0x5C; i += 2 {

		data = utils.Swapu16Little(data, i)

	}

	data = utils.Swapu16Little(data, 0x5C)

	return data

}

// converts a bit to a bool
func bitToBool(data byte) bool {

	if data == 0x00 {

		return false

	}
	return true

}

// converts a uint64 to bool (nani)
func u64ToBool(data uint64) bool {

	if data == 0x00 {

		return false

	}
	return true

}

// ParseMii takes a mii as a byte array and parses it to a Mii
// TODO: potentially hardcode offsets for `Seek` calls
func ParseMii(miiByte []byte) Mii {

	buf := utils.NewByteBuffer(swapMiiEndiannessToLittle(miiByte))
	btf := utils.NewBitfield(miiByte...)
	mii := Mii{}

	mii.BirthPlatform = btf.ReadBitsNext(4)
	mii.Unknown1 = btf.ReadBitsNext(4)
	mii.Unknown2 = btf.ReadBitsNext(4)
	mii.Unknown3 = btf.ReadBitsNext(4)
	mii.FontRegion = btf.ReadBitsNext(4)
	mii.RegionMove = btf.ReadBitsNext(2)
	mii.Unknown4 = btf.ReadBitNext()
	mii.Copyable = bitToBool(btf.ReadBitNext())

	buf.Seek(btf.Offset()/8, false)

	mii.MiiVersion = buf.ReadBytesNext(1)[0]
	mii.AuthorID = buf.ReadBytesNext(8)
	mii.MiiID = buf.ReadBytesNext(10)
	mii.Unknown5 = buf.ReadBytesNext(2)

	btf.Seek(buf.Offset()*8, false)

	mii.Unknown6 = btf.ReadBitNext()
	mii.Unknown7 = btf.ReadBitNext()
	mii.Color = btf.ReadBitsNext(4)
	mii.BirthDay = btf.ReadBitsNext(5)
	mii.BirthMonth = btf.ReadBitsNext(4)
	mii.Gender = btf.ReadBitNext()

	buf.Seek(btf.Offset()/8, false)

	// TODO: potentially optimize this
	tmp := buf.ReadBytesNext(20)
	mii.MiiName = ""
	for len(tmp) > 0 {

		r, size := utf8.DecodeRune(tmp)
		tmp = tmp[size:]
		if r == 0 {

			continue

		}
		mii.MiiName = mii.MiiName + string(r)

	}

	mii.Fatness = buf.ReadBytesNext(1)[0]
	mii.Size = buf.ReadBytesNext(1)[0]

	btf.Seek(buf.Offset()*8, false)

	mii.BlushType = btf.ReadBitsNext(4)
	mii.FaceStyle = btf.ReadBitsNext(4)
	mii.FaceColor = btf.ReadBitsNext(3)
	mii.FaceType = btf.ReadBitsNext(4)
	mii.LocalOnly = bitToBool(btf.ReadBitNext())
	mii.HairMirrored = u64ToBool(btf.ReadBitsNext(5))
	mii.HairColor = btf.ReadBitsNext(3)

	buf.Seek(btf.Offset()/8, false)

	mii.HairType = buf.ReadBytesNext(1)[0]

	btf.Seek(buf.Offset()*8, false)

	mii.EyeThickness = btf.ReadBitsNext(3)
	mii.EyeScale = btf.ReadBitsNext(4)
	mii.EyeColor = btf.ReadBitsNext(3)
	mii.EyeType = btf.ReadBitsNext(6)
	mii.EyeHeight = btf.ReadBitsNext(7)
	mii.EyeDistance = btf.ReadBitsNext(4)
	mii.EyeRotation = btf.ReadBitsNext(5)

	mii.EyebrowThickness = btf.ReadBitsNext(4)
	mii.EyebrowScale = btf.ReadBitsNext(4)
	mii.EyebrowColor = btf.ReadBitsNext(3)
	mii.EyebrowType = btf.ReadBitsNext(5)
	mii.EyebrowHeight = btf.ReadBitsNext(7)
	mii.EyebrowDistance = btf.ReadBitsNext(4)
	mii.EyebrowRotation = btf.ReadBitsNext(5)

	mii.NoseHeight = btf.ReadBitsNext(7)
	mii.NoseScale = btf.ReadBitsNext(4)
	mii.NoseType = btf.ReadBitsNext(5)

	mii.MouthThickness = btf.ReadBitsNext(3)
	mii.MouthScale = btf.ReadBitsNext(4)
	mii.MouthColor = btf.ReadBitsNext(3)
	mii.MouthType = btf.ReadBitsNext(6)

	buf.Seek(btf.Offset()/8, false)
	
	mii.Unknown8 = buf.ReadBytesNext(1)[0]

	btf.Seek(buf.Offset()*8, false)

	mii.MustacheType = btf.ReadBitsNext(3)
	mii.MouthHeight = btf.ReadBitsNext(5)
	mii.MustacheHeight = btf.ReadBitsNext(6)
	mii.MustacheScale = btf.ReadBitsNext(4)
	mii.BeardColor = btf.ReadBitsNext(3)
	mii.BeardType = btf.ReadBitsNext(3)

	mii.GlassesHeight = btf.ReadBitsNext(5)
	mii.GlassesScale = btf.ReadBitsNext(4)
	mii.GlassesColor = btf.ReadBitsNext(3)
	mii.GlassesType = btf.ReadBitsNext(4)
	mii.Unknown9 = btf.ReadBitNext()
	mii.MoleY = btf.ReadBitsNext(5)
	mii.MoleX = btf.ReadBitsNext(5)
	mii.MoleScale = btf.ReadBitsNext(4)
	mii.MoleEnabled = bitToBool(btf.ReadBitNext())

	buf.Seek(btf.Offset()/8, false)

	// TODO: potentially optimize this too
	tmp = buf.ReadBytesNext(20)
	mii.AuthorName = ""
	for len(tmp) > 0 {

		r, size := utf8.DecodeRune(tmp)
		tmp = tmp[size:]
		if r == 0 {

			continue

		}
		mii.AuthorName = mii.AuthorName + string(r)

	}

	mii.Unknown10 = buf.ReadBytesNext(2)
	mii.Checksum = buf.ReadComplexNext(1, utils.Unsigned16, utils.LittleEndian).([]uint16)[0]

	// TODO: add proper checksum validation
	
	return mii

}
