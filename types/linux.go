package types

type DpkgInfoObject struct {
	Object
	Name   OVALEntityObjectString `xml:"name"`
	Filter []Filter               `xml:"filter"`
}

type DpkgInfoState struct {
	State
	Name    OVALEntityStateString    `xml:"name"`
	Arch    OVALEntityStateString    `xml:"arch"`
	Epoch   OVALEntityStateAnySimple `xml:"epoch"`
	Release OVALEntityStateAnySimple `xml:"release"`
	Version OVALEntityStateAnySimple `xml:"version"`
	EVR     OVALEntityStateEVRString `xml:"evr"`
}

type DpkgInfoTest struct {
	Test
	Object ObjectRef `xml:"object"`
	State  StateRef  `xml:"state"`
}

type RpmInfoObject struct {
	Object
	Behaviors []RpmInfoBehaviors
	Name      OVALEntityObjectString `xml:"name"`
	Filter    []Filter               `xml:"filter"`
}

type RpmInfoState struct {
	State
	Name           OVALEntityStateString    `xml:"name"`
	Arch           OVALEntityStateString    `xml:"arch"`
	Epoch          OVALEntityStateAnySimple `xml:"epoch"`
	Release        OVALEntityStateAnySimple `xml:"release"`
	Version        OVALEntityStateAnySimple `xml:"version"`
	EVR            OVALEntityStateEVRString `xml:"evr"`
	SignatureKeyID OVALEntityStateString
	ExtendedName   OVALEntityStateString
	Filepath       OVALEntityStateString
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
	ObjectRef OVALObjectID `xml:"object_ref,attr"`
}

type StateRef struct {
	StateRef OVALStateID `xml:"state_ref,attr"`
}
