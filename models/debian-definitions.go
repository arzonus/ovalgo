package models

import (
	"github.com/arzonus/ovalgo/utils/ovaltime"
)

type DebianOVALDefinitions struct {
	OVALDefinitions
	Definitions []DebianDefinition `xml:"definitions>definition"`

	DpkgInfoTests        []DpkgInfoTest        `xml:"tests>dpkginfo_test"`
	TextFileContentTests []TextFileContentTest `xml:"tests>textfilecontent_test"`
	UnameTests           []UnameTest           `xml:"tests>uname_test"`

	DpkgInfoObjects        []DpkgInfoObject        `xml:"objects>dpkginfo_object"`
	TextFileContentObjects []TextFileContentObject `xml:"objects>textfilecontent_object"`
	UnameObjects           []UnameObject           `xml:"objects>uname_object"`

	DpkgInfoStates        []DpkgInfoState        `xml:"states>dpkginfo_state"`
	TextFileContentStates []TextFileContentState `xml:"states>textfilecontent_state"`
	UnameStates           []UnameState           `xml:"states>uname_state"`
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
