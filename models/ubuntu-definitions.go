package models

type UbuntuOVALDefinitions struct {
	OVALDefinitions
	Definitions     []DebianDefinition `xml:"definitions>definition"`
	DpkgInfoTests   []DpkgInfoTest     `xml:"tests>dpkginfo_test"`
	FamilyTests     []FamilyTest       `xml:"tests>family_test"`
	UnameTests      []UnameTest        `xml:"tests>uname_test"`
	UnknownTests    []UnknownTest      `xml:"tests>unknown_test"`
	DpkgInfoObjects []DpkgInfoObject   `xml:"objects>dpkginfo_object"`
	DpkgInfoStates  []DpkgInfoState    `xml:"states>dpkginfo_state"`
}
