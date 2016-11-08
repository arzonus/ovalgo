package types

import (
	"github.com/arzonus/ovalgo/utils/rfctime"
)

type UbuntuDefinitions struct {
	Generator   UbuntuGenerator  `xml:"generator"`
	Definitions []Definition     `xml:"definitions>definition"`
	Tests       []DpkgInfoTest   `xml:"tests>dpkginfo_test"`
	Objects     []DpkgInfoObject `xml:"objects>dpkginfo_object"`
	States      []DpkgInfoState  `xml:"states>dpkginfo_state"`
}

type UbuntuGenerator struct {
	ProductName    []string     `xml:"product_name"`
	ProductVersion []string     `xml:"product_version"`
	SchemaVersion  string       `xml:"schema_version"`
	Timestamp      rfctime.Time `xml:"timestamp"`
	Any            string       `xml:"xsd:any"`
}
