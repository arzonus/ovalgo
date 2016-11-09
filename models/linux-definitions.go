package models

type DpkgInfoObject struct {
	Object
	Name   EntityObjectString `xml:"name"`
	Filter []Filter           `xml:"filter"`
}

type DpkgInfoState struct {
	State
	Name    EntityStateString    `xml:"name"`
	Arch    EntityStateString    `xml:"arch"`
	Epoch   EntityStateAnySimple `xml:"epoch"`
	Release EntityStateAnySimple `xml:"release"`
	Version EntityStateAnySimple `xml:"version"`
	EVR     EntityStateEVRString `xml:"evr"`
}

type DpkgInfoTest struct {
	Test
	Object ObjectRef `xml:"object"`
	State  StateRef  `xml:"state"`
}

type RpmInfoObject struct {
	Object
	Behaviors []RpmInfoBehaviors
	Name      EntityObjectString `xml:"name"`
	Filter    []Filter           `xml:"filter"`
}

type RpmInfoState struct {
	State
	Name           EntityStateString    `xml:"name"`
	Arch           EntityStateString    `xml:"arch"`
	Epoch          EntityStateAnySimple `xml:"epoch"`
	Release        EntityStateAnySimple `xml:"release"`
	Version        EntityStateAnySimple `xml:"version"`
	EVR            EntityStateEVRString `xml:"evr"`
	SignatureKeyID EntityStateString
	ExtendedName   EntityStateString
	Filepath       EntityStateString
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
	ObjectRef ObjectID `xml:"object_ref,attr"`
}

type StateRef struct {
	StateRef StateID `xml:"state_ref,attr"`
}
