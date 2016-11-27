package models

type UnameTest struct {
	Test
	Object ObjectRef  `xml:"object"`
	States []StateRef `xml:"state"`
}

type UnameObject struct {
	Object
}

type UnameState struct {
	State
	MachineClass  string `xml:"machine_class"`
	NodeName      string `xml:"node_name"`
	OSName        string `xml:"os_name"`
	OSRelease     string `xml:"os_release"`
	OSVersion     string `xml:"os_version"`
	ProcessorType string `xml:"processor_type"`
}
