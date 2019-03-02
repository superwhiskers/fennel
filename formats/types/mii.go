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

package types

import (
	"github.com/superwhiskers/fennel/utils"
	"github.com/superwhiskers/crunch"
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

// swaps the endianness of a mii binary format to big-endian
func swapMiiEndiannessToBig(data []byte) []byte {

	data = utils.Swapu32Big(data, 0x00)

	for i := 0x18; i <= 0x2E; i += 2 {

		data = utils.Swapu16Big(data, i)

	}

	for i := 0x30; i <= 0x48; i += 2 {

		data = utils.Swapu16Big(data, i)

	}

	for i := 0x48; i <= 0x5C; i += 2 {

		data = utils.Swapu16Big(data, i)

	}

	data = utils.Swapu16Big(data, 0x5C)

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

	buf := crunch.NewBuffer(swapMiiEndiannessToLittle(miiByte))
	mii := Mii{}

	mii.BirthPlatform = buf.ReadBitsNext(4)
	mii.Unknown1 = buf.ReadBitsNext(4)
	mii.Unknown2 = buf.ReadBitsNext(4)
	mii.Unknown3 = buf.ReadBitsNext(4)
	mii.FontRegion = buf.ReadBitsNext(4)
	mii.RegionMove = buf.ReadBitsNext(2)
	mii.Unknown4 = buf.ReadBitNext()
	mii.Copyable = bitToBool(buf.ReadBitNext())

	buf.AlignByte()

	mii.MiiVersion = buf.ReadBytesNext(1)[0]
	mii.AuthorID = buf.ReadBytesNext(8)
	mii.MiiID = buf.ReadBytesNext(10)
	mii.Unknown5 = buf.ReadBytesNext(2)

	buf.AlignBit()

	mii.Unknown6 = buf.ReadBitNext()
	mii.Unknown7 = buf.ReadBitNext()
	mii.Color = buf.ReadBitsNext(4)
	mii.BirthDay = buf.ReadBitsNext(5)
	mii.BirthMonth = buf.ReadBitsNext(4)
	mii.Gender = buf.ReadBitNext()

	buf.AlignByte()

	mii.MiiName = utils.DecodeUTF8String(buf.ReadBytesNext(20))
	mii.Fatness = buf.ReadBytesNext(1)[0]
	mii.Size = buf.ReadBytesNext(1)[0]

	buf.AlignBit()

	mii.BlushType = buf.ReadBitsNext(4)
	mii.FaceStyle = buf.ReadBitsNext(4)
	mii.FaceColor = buf.ReadBitsNext(3)
	mii.FaceType = buf.ReadBitsNext(4)
	mii.LocalOnly = bitToBool(buf.ReadBitNext())
	mii.HairMirrored = u64ToBool(buf.ReadBitsNext(5))
	mii.HairColor = buf.ReadBitsNext(3)

	buf.AlignByte()

	mii.HairType = buf.ReadBytesNext(1)[0]

	buf.AlignBit()

	mii.EyeThickness = buf.ReadBitsNext(3)
	mii.EyeScale = buf.ReadBitsNext(4)
	mii.EyeColor = buf.ReadBitsNext(3)
	mii.EyeType = buf.ReadBitsNext(6)
	mii.EyeHeight = buf.ReadBitsNext(7)
	mii.EyeDistance = buf.ReadBitsNext(4)
	mii.EyeRotation = buf.ReadBitsNext(5)

	mii.EyebrowThickness = buf.ReadBitsNext(4)
	mii.EyebrowScale = buf.ReadBitsNext(4)
	mii.EyebrowColor = buf.ReadBitsNext(3)
	mii.EyebrowType = buf.ReadBitsNext(5)
	mii.EyebrowHeight = buf.ReadBitsNext(7)
	mii.EyebrowDistance = buf.ReadBitsNext(4)
	mii.EyebrowRotation = buf.ReadBitsNext(5)

	mii.NoseHeight = buf.ReadBitsNext(7)
	mii.NoseScale = buf.ReadBitsNext(4)
	mii.NoseType = buf.ReadBitsNext(5)

	mii.MouthThickness = buf.ReadBitsNext(3)
	mii.MouthScale = buf.ReadBitsNext(4)
	mii.MouthColor = buf.ReadBitsNext(3)
	mii.MouthType = buf.ReadBitsNext(6)

	buf.AlignByte()
	
	mii.Unknown8 = buf.ReadBytesNext(1)[0]

	buf.AlignBit()

	mii.MustacheType = buf.ReadBitsNext(3)
	mii.MouthHeight = buf.ReadBitsNext(5)
	mii.MustacheHeight = buf.ReadBitsNext(6)
	mii.MustacheScale = buf.ReadBitsNext(4)
	mii.BeardColor = buf.ReadBitsNext(3)
	mii.BeardType = buf.ReadBitsNext(3)

	mii.GlassesHeight = buf.ReadBitsNext(5)
	mii.GlassesScale = buf.ReadBitsNext(4)
	mii.GlassesColor = buf.ReadBitsNext(3)
	mii.GlassesType = buf.ReadBitsNext(4)
	mii.Unknown9 = buf.ReadBitNext()
	mii.MoleY = buf.ReadBitsNext(5)
	mii.MoleX = buf.ReadBitsNext(5)
	mii.MoleScale = buf.ReadBitsNext(4)
	mii.MoleEnabled = bitToBool(buf.ReadBitNext())

	buf.AlignByte()

	mii.AuthorName = utils.DecodeUTF8String(buf.ReadBytesNext(20))
	mii.Unknown10 = buf.ReadBytesNext(2)
	mii.Checksum = buf.ReadComplexNext(1, crunch.Unsigned16, crunch.LittleEndian).([]uint16)[0]

	// TODO: add proper checksum validation
	
	return mii

}
