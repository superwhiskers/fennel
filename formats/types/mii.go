/*

libninty - nintendo network utility library for golang
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
	//"bytes"
	"encoding/binary"
)

// applies a slice to a slice at the specified offset
func applySliceAtOffset(src, dest []byte, offset int) []byte {

	for i, byt := range src {

		dest[offset+i] = byt

	}
	return dest

}


// uint32 little-endian -> big-endian
// offsets should be zero-indexed
func swap32Big(data []byte, offset int) []byte {

	byteSection := []byte{0, 0, 0, 0}
	binary.BigEndian.PutUint32(byteSection, binary.LittleEndian.Uint32(data[offset:offset+4]))
	return applySliceAtOffset(byteSection, data, offset)

}

// uint32 big-endian -> little-endian
// offsets should be zero-indexed
func swap32Little(data []byte, offset int) []byte {

	byteSection := []byte{0, 0, 0, 0}
	binary.LittleEndian.PutUint32(byteSection, binary.BigEndian.Uint32(data[offset:offset+4]))
	return applySliceAtOffset(byteSection, data, offset)

}

// uint16 little-endian -> big-endian
// offsets should be zero-indexed
func swap16Big(data []byte, offset int) []byte {

	byteSection := []byte{0, 0}
	binary.BigEndian.PutUint16(byteSection, binary.LittleEndian.Uint16(data[offset:offset+2]))
	return applySliceAtOffset(byteSection, data, offset)

}

// uint16 big-endian -> little-endian
// offsets should be zero-indexed
func swap16Little(data []byte, offset int) []byte {

	byteSection := []byte{0, 0}
	binary.LittleEndian.PutUint16(byteSection, binary.BigEndian.Uint16(data[offset:offset+2]))
	return applySliceAtOffset(byteSection, data, offset)

}

// Mii contains all of the data that a mii can have
type Mii struct {
	// unknown fields
	Unknown1  [4]byte
	Unknown2  [4]byte
	Unknown3  [4]byte
	Unknown4  byte
	Unknown5  [2]byte
	Unknown6  byte
	Unknown7  byte
	Unknown8  uint8
	Unknown9  byte
	Unknown10 [2]byte

	// attributes
	BirthPlatform [4]byte
	FontRegion    [4]byte
	RegionMove    [2]byte
	Copyable      bool
	MiiVersion    uint8
	AuthorID      [8]uint8
	MiiID         [10]uint8
	LocalOnly     bool
	Color         [4]byte
	BirthDay      [5]byte
	BirthMonth    [4]byte
	Gender        byte
	MiiName       string
	Size          uint8
	Fatness       uint8
	AuthorName    string
	Checksum      uint16

	// face
	BlushType [4]byte
	FaceStyle [4]byte
	FaceColor [3]byte
	FaceType  [4]byte

	// hair
	HairMirrored bool
	HairColor    [3]byte
	HairType     uint8

	// eyes
	EyeThickness [3]byte
	EyeScale     [4]byte
	EyeColor     [3]byte
	EyeType      [6]byte
	EyeHeight    [7]byte
	EyeDistance  [4]byte
	EyeRotation  [5]byte

	// eyebrow
	EyebrowThickness [4]byte
	EyebrowScale     [4]byte
	EyebrowColor     [3]byte
	EyebrowType      [5]byte
	EyebrowHeight    [7]byte
	EyebrowDistance  [4]byte
	EyebrowRotation  [5]byte

	// nose
	NoseHeight [7]byte
	NoseScale  [4]byte
	NoseType   [5]byte

	// mouth
	MouthThickness [3]byte
	MouthScale     [4]byte
	MouthColor     [3]byte
	MouthType      [6]byte
	MouthHeight    [5]byte

	// mustache
	MustacheType   [3]byte
	MustacheHeight [6]byte
	MustacheScale  [4]byte

	// beard
	BeardColor [3]byte
	BeardType  [3]byte

	// glasses
	GlassesHeight [5]byte
	GlassesScale  [4]byte
	GlassesColor  [3]byte
	GlassesType   [4]byte

	// mole
	MoleY       [5]byte
	MoleX       [5]byte
	MoleScale   [4]byte
	MoleEnabled bool
}

// NilMii is a Mii with no data
var NilMii = Mii{
	Unknown1:         [4]byte{0, 0, 0, 0},
	Unknown2:         [4]byte{0, 0, 0, 0},
	Unknown3:         [4]byte{0, 0, 0, 0},
	Unknown4:         0x00,
	Unknown5:         [2]byte{0, 0},
	Unknown6:         0x00,
	Unknown7:         0x00,
	Unknown8:         uint8(0),
	Unknown9:         0x00,
	Unknown10:        [2]byte{0, 0},
	BirthPlatform:    [4]byte{0, 0, 0, 0},
	FontRegion:       [4]byte{0, 0, 0, 0},
	RegionMove:       [2]byte{0, 0},
	Copyable:         false,
	MiiVersion:       uint8(0),
	AuthorID:         [8]uint8{uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0)},
	MiiID:            [10]uint8{uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0)},
	LocalOnly:        false,
	Color:            [4]byte{0, 0, 0, 0},
	BirthDay:         [5]byte{0, 0, 0, 0, 0},
	BirthMonth:       [4]byte{0, 0, 0, 0},
	Gender:           0x00,
	MiiName:          "",
	Size:             uint8(0),
	Fatness:          uint8(0),
	AuthorName:       "",
	Checksum:         uint16(0),
	BlushType:        [4]byte{0, 0, 0, 0},
	FaceStyle:        [4]byte{0, 0, 0, 0},
	FaceColor:        [3]byte{0, 0, 0},
	FaceType:         [4]byte{0, 0, 0, 0},
	HairMirrored:     false,
	HairColor:        [3]byte{0, 0, 0},
	HairType:         uint8(0),
	EyeThickness:     [3]byte{0, 0, 0},
	EyeScale:         [4]byte{0, 0, 0, 0},
	EyeColor:         [3]byte{0, 0, 0},
	EyeType:          [6]byte{0, 0, 0, 0, 0, 0},
	EyeHeight:        [7]byte{0, 0, 0, 0, 0, 0, 0},
	EyeDistance:      [4]byte{0, 0, 0, 0},
	EyeRotation:      [5]byte{0, 0, 0, 0, 0},
	EyebrowThickness: [4]byte{0, 0, 0, 0},
	EyebrowScale:     [4]byte{0, 0, 0, 0},
	EyebrowColor:     [3]byte{0, 0, 0},
	EyebrowType:      [5]byte{0, 0, 0, 0, 0},
	EyebrowHeight:    [7]byte{0, 0, 0, 0, 0, 0, 0},
	EyebrowDistance:  [4]byte{0, 0, 0, 0},
	EyebrowRotation:  [5]byte{0, 0, 0, 0, 0},
	NoseHeight:       [7]byte{0, 0, 0, 0, 0, 0, 0},
	NoseScale:        [4]byte{0, 0, 0, 0},
	NoseType:         [5]byte{0, 0, 0, 0, 0},
	MouthThickness:   [3]byte{0, 0, 0},
	MouthScale:       [4]byte{0, 0, 0, 0},
	MouthColor:       [3]byte{0, 0, 0},
	MouthType:        [6]byte{0, 0, 0, 0, 0, 0},
	MouthHeight:      [5]byte{0, 0, 0, 0, 0},
	MustacheType:     [3]byte{0, 0, 0},
	MustacheHeight:   [6]byte{0, 0, 0, 0, 0, 0},
	MustacheScale:    [4]byte{0, 0, 0, 0},
	BeardColor:       [3]byte{0, 0, 0},
	BeardType:        [3]byte{0, 0, 0},
	GlassesHeight:    [5]byte{0, 0, 0, 0, 0},
	GlassesScale:     [4]byte{0, 0, 0, 0},
	GlassesColor:     [3]byte{0, 0, 0},
	GlassesType:      [4]byte{0, 0, 0, 0},
	MoleY:            [5]byte{0, 0, 0, 0, 0},
	MoleX:            [5]byte{0, 0, 0, 0, 0},
	MoleScale:        [4]byte{0, 0, 0, 0},
	MoleEnabled:      false,
}

// ParseMii takes a mii as a byte array and parses it to a Mii
func ParseMii(miiByte []byte) Mii {

	/*section := miiByte[:95]
	bigEndian := []byte{}

	// run buffer.Next a bit here*/
	
	return NilMii

}
