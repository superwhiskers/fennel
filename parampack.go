/*

parampack.go -
contains things for handling parampacks

*/

package libninty

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*

Parampack implements a struct for containing data housed in nintendo parampacks in a golang-compatible format

*/
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

/*

NilParampack is a parampack type that contains no data

*/
var NilParampack = Parampack{
	TitleID:            "0000000000000000",
	AccessKey:          "",
	PlatformID:         0,
	RegionID:           0,
	LanguageID:         0,
	CountryID:          0,
	AreaID:             0,
	NetworkRestriction: 0,
	FriendRestriction:  0,
	RatingRestriction:  0,
	RatingOrganization: 0,
	TransferableID:     "",
	TimezoneName:       "",
	UTCOffset:          0,
	RemasterVersion:    0,
}

/*

StringifyParampack is a method of the Parampack type that returns a stringified version of the parampack

*/
func (p Parampack) StringifyParampack() string {

	// it's just return the stringified version of the parampack
	return fmt.Sprintf("\\title_id\\%s\\access_key\\%s\\platform_id\\%d\\region_id\\%d\\language_id\\%d\\country_id\\%d\\area_id\\%d\\network_restriction\\%d\\friend_restriction\\%d\\rating_restriction\\%d\\rating_organization\\%d\\transferable_id\\%s\\tz_name\\%s\\utc_offset\\%d\\remaster_version\\%d\\", p.TitleID, p.AccessKey, p.PlatformID, p.RegionID, p.LanguageID, p.CountryID, p.AreaID, p.NetworkRestriction, p.FriendRestriction, p.RatingRestriction, p.RatingOrganization, p.TransferableID, p.TimezoneName, p.UTCOffset, p.RemasterVersion)

}

/*

EncodeParampack is a method of the Parampack type that encodes a stringified parampack into base64

*/
func (p Parampack) EncodeParampack() string {

	// convert the stringified parampack to base64
	return base64.StdEncoding.EncodeToString([]byte(p.StringifyParampack()))

}

/*

UnstringifyParampack takes a stringified parampack and places it into a struct

*/
func UnstringifyParampack(parampack string) Parampack {

	// split it by backslashes
	splitParampack := strings.Split(parampack, "\\")

	// variables to be placed into the struct
	titleID := "0000000000000000"
	accessKey := ""
	platformID := 0
	regionID := 0
	languageID := 0
	countryID := 0
	areaID := 0
	networkRestriction := 0
	friendRestriction := 0
	ratingRestriction := 0
	ratingOrganization := 0
	transferableID := ""
	timezoneName := ""
	utcOffset := 0
	remasterVersion := 0

	// iterate over the split parampack
	for ind, ele := range splitParampack {

		// check if it is one of the parts of a parameter pack
		// and assign its value to the corresponding variable
		switch ele {

		case "title_id":

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

	// finally, formulate a parampack struct
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

/*

DecodeParampack takes a base64ed parampack and decodes it into a struct

*/
func DecodeParampack(parampack string) (Parampack, error) {

	// strip spaces
	paramStripped := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, parampack)

	// decode it from base64
	decodedParampack, err := base64.StdEncoding.DecodeString(paramStripped)

	// if there is an error
	if err != nil {

		// exit the function and return the error
		return NilParampack, err

	}

	// and return it
	return UnstringifyParampack(string(decodedParampack[:])), nil

}
