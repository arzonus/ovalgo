package models

type DebianOVALDefinitions struct {
	Generator   Generator        `xml:"generator"`
	Definitions []Definition     `xml:"definitions>definition"`
	Tests       []DpkgInfoTest   `xml:"tests>dpkginfo_test"`
	Objects     []DpkgInfoObject `xml:"objects>dpkginfo_object"`
	States      []DpkgInfoState  `xml:"states>dpkginfo_state"`
}
