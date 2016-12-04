package models

type DpkgInfoObject struct {
	Object
	Name   string   `xml:"name"`
	Filter []Filter `xml:"filter"`
}

type DpkgInfoState struct {
	State
	Name    ValueOperation `xml:"name"`
	Arch    ValueOperation `xml:"arch"`
	Epoch   ValueOperation `xml:"epoch"`
	Release ValueOperation `xml:"release"`
	Version ValueOperation `xml:"version"`
	EVR     ValueOperation `xml:"evr"`
}

type DpkgInfoTest struct {
	Test
	Object ObjectRef  `xml:"object"`
	States []StateRef `xml:"state"`
}

type RpmInfoObject struct {
	Object
	Behaviors []RpmInfoBehaviors
	Name      string   `xml:"name"`
	Filter    []Filter `xml:"filter"`
}

type RpmInfoState struct {
	State
	Name           ValueOperation `xml:"name"`
	Arch           ValueOperation `xml:"arch"`
	Epoch          ValueOperation `xml:"epoch"`
	Release        ValueOperation `xml:"release"`
	Version        ValueOperation `xml:"version"`
	EVR            ValueOperation `xml:"evr"`
	SignatureKeyID ValueOperation
	ExtendedName   ValueOperation
	Filepath       ValueOperation
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
