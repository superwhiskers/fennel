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

package formats

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
