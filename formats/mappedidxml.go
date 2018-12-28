package formats

import "encoding/xml"

// NillMappedIDsXML represents a MappedIDsXML with no value
var NilMappedIDsXML = MappedIDsXML{
	MappedIDs: []MappedIDsXMLMappedID{},
}

// MappedIDsXML represents a nintendo network mappedids xml sheet
type MappedIDsXML struct {
	XMLName   xml.Name               `xml:"mapped_ids"`
	MappedIDs []MappedIDsXMLMappedID `xml:"mapped_id"`
}

// MappedIDsXMLMappedID represents a mapped id of a MappedIDsXML
type MappedIDsXMLMappedID struct {
	InID  string `xml:"in_id"`
	OutID string `xml:"out_id"`
}

