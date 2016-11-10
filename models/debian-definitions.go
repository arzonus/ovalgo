package models

import (
	"github.com/arzonus/ovalgo/utils/ovaltime"
)

type DebianOVALDefinitions struct {
	Generator   Generator          `xml:"generator"`
	Definitions []DebianDefinition `xml:"definitions>definition"`
	Tests       []DpkgInfoTest     `xml:"tests>dpkginfo_test"`
	Objects     []DpkgInfoObject   `xml:"objects>dpkginfo_object"`
	States      []DpkgInfoState    `xml:"states>dpkginfo_state"`
}

type DebianDefinition struct {
	Definition
	Metadata DebianMetadata `xml:"metadata"`
}

type DebianMetadata struct {
	Metadata
	Debian DebianMetadataDebian `xml:"debian"`
}

type DebianMetadataDebian struct {
	Date     ovaltime.Time `xml:"date"`
	MoreInfo string        `xml:"moreinfo"`
}
