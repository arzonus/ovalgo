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
	Filepath      ValueOperation `xml:"filepath"`
	Path          ValueOperation `xml:"path"`
	Filename      ValueOperation `xml:"filename"`
	Pattern       ValueOperation `xml:"pattern"`
	Instance      ValueOperation `xml:"instance"`
	Text          ValueOperation `xml:"text"`
	Subexpression ValueOperation `xml:"subexpression"`
	WindowsView   ValueOperation `xml:"windows_view"`
	Line          ValueOperation `xml:"line"`
}

type FamilyState struct {
	State
	Family ValueOperation `xml:"family"`
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
