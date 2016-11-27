package models

type TextFileContentTest struct {
	Test
	Object ObjectRef  `xml:"object"`
	States []StateRef `xml:"state"`
}

type FamilyTest struct {
	Test
	Object ObjectRef  `xml:"object"`
	States []StateRef `xml:"state"`
}

type UnknownTest struct {
	Test
}

type TextFileContentState struct {
	State
	Filepath      string `xml:"filepath"`
	Path          string `xml:"path"`
	Filename      string `xml:"filename"`
	Pattern       string `xml:"pattern"`
	Instance      uint   `xml:"instance"`
	Text          string `xml:"text"`
	Subexpression string `xml:"subexpression"`
	WindowsView   string `xml:"windows_view"`
	Operation     string `xml:"operation,attr"`
	Line          Line   `xml:"line"`
}

type Line struct {
	Line      string `xml:",innerxml"`
	Operation string `xml:"operation,attr"`
}

type FamilyState struct {
	State
	Family string `xml:"family"`
}

type TextFileContentObject struct {
	Object
	Filepath string `xml:"filepath"`
	Path     string `xml:"path"`
	Filename string `xml:"filename"`
	Pattern  string `xml:"pattern"`
	Instance string `xml:"instance"`
}

type FamilyObject struct {
	Object
}
