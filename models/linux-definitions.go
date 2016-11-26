package models

type DpkgInfoObject struct {
	Object
	Name   string   `xml:"name"`
	Filter []Filter `xml:"filter"`
}

type DpkgInfoState struct {
	State
	Name    string `xml:"name"`
	Arch    string `xml:"arch"`
	Epoch   string `xml:"epoch"`
	Release string `xml:"release"`
	Version string `xml:"version"`
	EVR     EVR    `xml:"evr"`
}

type EVR struct {
	EVR       string `xml:",innerxml"`
	Operation string `xml:"operation,attr"`
}

type DpkgInfoTest struct {
	Test
	Object ObjectRef `xml:"object"`
	State  StateRef  `xml:"state"`
}

type RpmInfoObject struct {
	Object
	Behaviors []RpmInfoBehaviors
	Name      string   `xml:"name"`
	Filter    []Filter `xml:"filter"`
}

type RpmInfoState struct {
	State
	Name           string `xml:"name"`
	Arch           string `xml:"arch"`
	Epoch          string `xml:"epoch"`
	Release        string `xml:"release"`
	Version        string `xml:"version"`
	EVR            string `xml:"evr"`
	SignatureKeyID string
	ExtendedName   string
	Filepath       string
}

type RpmInfoTest struct {
	Test
	Object ObjectRef `xml:"object"`
	State  StateRef  `xml:"state"`
}

type RpmInfoBehaviors struct {
	Filepath bool
}

type ObjectRef struct {
	ObjectRef string `xml:"object_ref,attr"`
}

type StateRef struct {
	StateRef string `xml:"state_ref,attr"`
}
