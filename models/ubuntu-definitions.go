package models

type UbuntuOVALDefinitions struct {
	OVALDefinitions
	Definitions []UbuntuDefinition `xml:"definitions>definition"`

	DpkgInfoTests []DpkgInfoTest `xml:"tests>dpkginfo_test"`
	UnameTests    []UnameTest    `xml:"tests>uname_test"`
	FamilyTests   []FamilyTest   `xml:"tests>family_test"`
	UnknownTests  []UnknownTest  `xml:"tests>unknown_test"`

	DpkgInfoObjects []DpkgInfoObject `xml:"objects>dpkginfo_object"`
	UnameObjects    []UnameObject    `xml:"objects>uname_object"`
	FamilyObjects   []FamilyObject   `xml:"objects>family_object"`

	DpkgInfoStates []DpkgInfoState `xml:"states>dpkginfo_state"`
	UnameStates    []UnameState    `xml:"states>uname_state"`
	FamilyStates   []FamilyState   `xml:"states>family_state"`
}

type UbuntuDefinition struct {
	Definition
	Metadata UbuntuMetadata `xml:"metadata"`
}

type UbuntuMetadata struct {
	Metadata
}
