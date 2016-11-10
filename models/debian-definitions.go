package models

import (
	"github.com/arzonus/ovalgo/utils/ovaltime"
)

type DebianOVALDefinitions struct {
	OVALDefinitions
	Definitions          []DebianDefinition    `xml:"definitions>definition"`
	DpkgInfoTests        []DpkgInfoTest        `xml:"tests>dpkginfo_test"`
	TextFileContentTests []TextFileContentTest `xml:"tests>textfilecontent_test"`
	UnameTests           []UnameTest           `xml:"tests>uname_test"`
	DpkgInfoObjects      []DpkgInfoObject      `xml:"objects>dpkginfo_object"`
	DpkgInfoStates       []DpkgInfoState       `xml:"states>dpkginfo_state"`
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
