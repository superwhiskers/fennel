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
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Parampack is a data structure that represents a nintendo parampack
type Parampack struct {
	TitleID            string
	AccessKey          string
	PlatformID         int
	RegionID           int
	LanguageID         int
	CountryID          int
	AreaID             int
	NetworkRestriction int
	FriendRestriction  int
	RatingRestriction  int
	RatingOrganization int
	TransferableID     string
	TimezoneName       string
	UTCOffset          int
	RemasterVersion    int
}

// NilParampack is a Parampack with no data
var NilParampack = Parampack{}

// FormatString formats the Parampack as a string
func (p Parampack) FormatString() string {

	return fmt.Sprintf("\\title_id\\%s\\access_key\\%s\\platform_id\\%d\\region_id\\%d\\language_id\\%d\\country_id\\%d\\area_id\\%d\\network_restriction\\%d\\friend_restriction\\%d\\rating_restriction\\%d\\rating_organization\\%d\\transferable_id\\%s\\tz_name\\%s\\utc_offset\\%d\\remaster_version\\%d\\", p.TitleID, p.AccessKey, p.PlatformID, p.RegionID, p.LanguageID, p.CountryID, p.AreaID, p.NetworkRestriction, p.FriendRestriction, p.RatingRestriction, p.RatingOrganization, p.TransferableID, p.TimezoneName, p.UTCOffset, p.RemasterVersion)

}

// FormatSource formats the Parampack in the source format
func (p Parampack) FormatSource() string {

	return base64.StdEncoding.EncodeToString([]byte(p.FormatString()))

}

// ParseStringParampack takes a parampack as a string and parses it to a Parampack
func ParseStringParampack(parampack string) Parampack {

	var (
		splitParampack     = strings.Split(parampack, "\\")
		titleID            = "0000000000000000"
		accessKey          = ""
		platformID         = 0
		regionID           = 0
		languageID         = 0
		countryID          = 0
		areaID             = 0
		networkRestriction = 0
		friendRestriction  = 0
		ratingRestriction  = 0
		ratingOrganization = 0
		transferableID     = ""
		timezoneName       = ""
		utcOffset          = 0
		remasterVersion    = 0
	)

	for ind, ele := range splitParampack {

		switch ele {

		case "title_id":

			// TODO: add a 3ds tid unstringifier
			// titleids are special
			/*
				unstringifiedTID, err := unstringifyTID(splitParampack[ind+1])
				if err != nil {
					unstringifiedTID = "0000000000000000"
				}
			*/
			titleID = splitParampack[ind+1]

		case "access_key":
			accessKey = splitParampack[ind+1]

		case "platform_id":
			tmp, err := strconv.Atoi(splitParampack[ind+1])
			if err != nil {
				tmp = 0
			}
			platformID = tmp

		case "region_id":
			tmp, err := strconv.Atoi(splitParampack[ind+1])
			if err != nil {
				tmp = 0
			}
			regionID = tmp

		case "language_id":
			tmp, err := strconv.Atoi(splitParampack[ind+1])
			if err != nil {
				tmp = 0
			}
			languageID = tmp

		case "country_id":
			tmp, err := strconv.Atoi(splitParampack[ind+1])
			if err != nil {
				tmp = 0
			}
			countryID = tmp

		case "area_id":
			tmp, err := strconv.Atoi(splitParampack[ind+1])
			if err != nil {
				tmp = 0
			}
			areaID = tmp

		case "network_restriction":
			tmp, err := strconv.Atoi(splitParampack[ind+1])
			if err != nil {
				tmp = 0
			}
			networkRestriction = tmp

		case "friend_restriction":
			tmp, err := strconv.Atoi(splitParampack[ind+1])
			if err != nil {
				tmp = 0
			}
			friendRestriction = tmp

		case "rating_restriction":
			tmp, err := strconv.Atoi(splitParampack[ind+1])
			if err != nil {
				tmp = 0
			}
			ratingRestriction = tmp

		case "rating_organization":
			tmp, err := strconv.Atoi(splitParampack[ind+1])
			if err != nil {
				tmp = 0
			}
			ratingOrganization = tmp

		case "transferable_id":
			transferableID = splitParampack[ind+1]

		case "tz_name":
			timezoneName = splitParampack[ind+1]

		case "utc_offset":
			tmp, err := strconv.Atoi(splitParampack[ind+1])
			if err != nil {
				tmp = 0
			}
			utcOffset = tmp

		case "remaster_version":
			tmp, err := strconv.Atoi(splitParampack[ind+1])
			if err != nil {
				tmp = 0
			}
			remasterVersion = tmp

		}

	}

	returnableParampack := Parampack{
		TitleID:            titleID,
		AccessKey:          accessKey,
		PlatformID:         platformID,
		RegionID:           regionID,
		LanguageID:         languageID,
		CountryID:          countryID,
		AreaID:             areaID,
		NetworkRestriction: networkRestriction,
		FriendRestriction:  friendRestriction,
		RatingRestriction:  ratingRestriction,
		RatingOrganization: ratingOrganization,
		TransferableID:     transferableID,
		TimezoneName:       timezoneName,
		UTCOffset:          utcOffset,
		RemasterVersion:    remasterVersion,
	}

	return returnableParampack

}

// ParseParampack takes a parampack in the source format and parses it to a Parampack
func ParseParampack(parampack string) (Parampack, error) {

	paramStripped := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, parampack)

	decodedParampack, err := base64.StdEncoding.DecodeString(paramStripped)
	if err != nil {

		return NilParampack, err

	}

	return ParseStringParampack(string(decodedParampack[:])), nil

}
