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
	MachineClass  ValueOperation `xml:"machine_class"`
	NodeName      ValueOperation `xml:"node_name"`
	OSName        ValueOperation `xml:"os_name"`
	OSRelease     ValueOperation `xml:"os_release"`
	OSVersion     ValueOperation `xml:"os_version"`
	ProcessorType ValueOperation `xml:"processor_type"`
}
