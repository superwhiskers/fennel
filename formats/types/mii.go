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

// converts a bool to a bit
func boolToBit(data bool) byte {

	if data == false {

		return 0x00

	}
	return 0x01

}

// converts a uint64 to bool (nani)
func u64ToBool(data uint64) bool {

	if data == 0x00 {

		return false

	}
	return true

}

// converts a bool to a uint64 (owo)
func boolTou64(data bool) uint64 {

	if data == false {

		return 0x00

	}
	return 0x01

}

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
	HairMirrored uint64
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

// ParseMii takes a mii as a byte array and returns a parsed Mii
func ParseMii(miiByte []byte) *Mii {

	mii := &Mii{}
	mii.Parse(miiByte)
	return mii

}

// Parse takes a mii as a byte array and parses it into a Mii
// TODO: potentially hardcode offsets for `Seek` calls
func (mii *Mii) Parse(miiByte []byte) {

	buf := crunch.NewBuffer(swapMiiEndiannessToLittle(miiByte))

	mii.BirthPlatform = buf.ReadBitsNext(4)
	mii.Unknown1 = buf.ReadBitsNext(4)
	mii.Unknown2 = buf.ReadBitsNext(4)
	mii.Unknown3 = buf.ReadBitsNext(4)
	mii.FontRegion = buf.ReadBitsNext(4)
	mii.RegionMove = buf.ReadBitsNext(2)
	mii.Unknown4 = buf.ReadBitNext()
	mii.Copyable = bitToBool(buf.ReadBitNext())

	buf.AlignByte()

	mii.MiiVersion = buf.ReadByteNext()
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

	mii.MiiName = utils.DecodeUTF8StringFromBytes(buf.ReadBytesNext(20))
	mii.Fatness = buf.ReadByteNext()
	mii.Size = buf.ReadByteNext()

	buf.AlignBit()

	mii.BlushType = buf.ReadBitsNext(4)
	mii.FaceStyle = buf.ReadBitsNext(4)
	mii.FaceColor = buf.ReadBitsNext(3)
	mii.FaceType = buf.ReadBitsNext(4)
	mii.LocalOnly = bitToBool(buf.ReadBitNext())
	mii.HairMirrored = buf.ReadBitsNext(5)
	mii.HairColor = buf.ReadBitsNext(3)

	buf.AlignByte()

	mii.HairType = buf.ReadByteNext()

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
	
	mii.Unknown8 = buf.ReadByteNext()

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

	mii.AuthorName = utils.DecodeUTF8StringFromBytes(buf.ReadBytesNext(20))
	mii.Unknown10 = buf.ReadBytesNext(2)
	mii.Checksum = buf.ReadComplexNext(1, crunch.Unsigned16, crunch.LittleEndian).([]uint16)[0]

	// TODO: add proper checksum validation

}

// Encode takes a Mii and encodes it as a byte array
// TODO: potentially hardcode offsets for `Seek` calls
func (mii *Mii) Encode() []byte {

	buf := crunch.NewBuffer(make([]byte, 0x60))

	buf.SetBitsNext(mii.BirthPlatform, 4)
	buf.SetBitsNext(mii.Unknown1, 4)
	buf.SetBitsNext(mii.Unknown2, 4)
	buf.SetBitsNext(mii.Unknown3, 4)
	buf.SetBitsNext(mii.FontRegion, 4)
	buf.SetBitsNext(mii.RegionMove, 2)
	buf.SetBitNext(mii.Unknown4)
	buf.SetBitNext(boolToBit(mii.Copyable))

	buf.AlignByte()

	buf.WriteByteNext(mii.MiiVersion)
	buf.WriteBytesNext(mii.AuthorID)
	buf.WriteBytesNext(mii.MiiID)
	buf.WriteBytesNext(mii.Unknown5)

	buf.AlignBit()

	buf.SetBitNext(mii.Unknown6)
	buf.SetBitNext(mii.Unknown7)
	buf.SetBitsNext(mii.Color, 4)
	buf.SetBitsNext(mii.BirthDay, 5)
	buf.SetBitsNext(mii.BirthMonth, 4)
	buf.SetBitNext(mii.Gender)

	buf.AlignByte()

	buf.WriteBytesNext(utils.EncodeBytesFromUTF8String(mii.MiiName))
	buf.WriteByteNext(mii.Size)
	buf.WriteByteNext(mii.Fatness)

	buf.AlignBit()

	buf.SetBitsNext(mii.BlushType, 4)
	buf.SetBitsNext(mii.FaceStyle, 4)
	buf.SetBitsNext(mii.FaceColor, 3)
	buf.SetBitsNext(mii.FaceType, 4)
	buf.SetBitNext(boolToBit(mii.LocalOnly))
	buf.SetBitsNext(mii.HairMirrored, 5)
	buf.SetBitsNext(mii.HairColor, 3)

	buf.AlignByte()

	buf.WriteByteNext(mii.HairType)

	buf.AlignBit()

	buf.SetBitsNext(mii.EyeThickness, 3)
	buf.SetBitsNext(mii.EyeScale, 4)
	buf.SetBitsNext(mii.EyeColor, 3)
	buf.SetBitsNext(mii.EyeType, 6)
	buf.SetBitsNext(mii.EyeHeight, 7)
	buf.SetBitsNext(mii.EyeDistance, 4)
	buf.SetBitsNext(mii.EyeRotation, 5)

	buf.SetBitsNext(mii.EyebrowThickness, 4)
	buf.SetBitsNext(mii.EyebrowScale, 4)
	buf.SetBitsNext(mii.EyebrowColor, 3)
	buf.SetBitsNext(mii.EyebrowType, 5)
	buf.SetBitsNext(mii.EyebrowHeight, 7)
	buf.SetBitsNext(mii.EyebrowDistance, 4)
	buf.SetBitsNext(mii.EyebrowRotation, 5)

	buf.SetBitsNext(mii.NoseHeight, 7)
	buf.SetBitsNext(mii.NoseScale, 4)
	buf.SetBitsNext(mii.NoseType, 5)
	
	buf.SetBitsNext(mii.MouthThickness, 3)
	buf.SetBitsNext(mii.MouthScale, 4)
	buf.SetBitsNext(mii.MouthColor, 3)
	buf.SetBitsNext(mii.MouthType, 6)

	buf.AlignByte()

	buf.WriteByteNext(mii.Unknown8)

	buf.AlignBit()

	buf.SetBitsNext(mii.MustacheType, 3)
	buf.SetBitsNext(mii.MouthHeight, 5)
	buf.SetBitsNext(mii.MustacheHeight, 6)
	buf.SetBitsNext(mii.MustacheScale, 4)
	buf.SetBitsNext(mii.BeardColor, 3)
	buf.SetBitsNext(mii.BeardType, 3)

	buf.SetBitsNext(mii.GlassesHeight, 5)
	buf.SetBitsNext(mii.GlassesScale, 4)
	buf.SetBitsNext(mii.GlassesColor, 3)
	buf.SetBitsNext(mii.GlassesType, 4)
	buf.SetBitNext(mii.Unknown9)
	buf.SetBitsNext(mii.MoleY, 5)
	buf.SetBitsNext(mii.MoleX, 5)
	buf.SetBitsNext(mii.MoleScale, 4)
	buf.SetBitNext(boolToBit(mii.MoleEnabled))

	buf.AlignByte()
	
	buf.WriteBytesNext(utils.EncodeBytesFromUTF8String(mii.AuthorName))
	buf.WriteBytesNext(mii.Unknown10)
	buf.WriteBytes(0x00, swapMiiEndiannessToBig(buf.Bytes()))
	buf.WriteComplexNext([]uint16{utils.CRC16(buf.Bytes())}, crunch.Unsigned16, crunch.LittleEndian)

	return buf.Bytes()
	
}
