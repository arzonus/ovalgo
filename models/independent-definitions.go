package models

type TextFileContentTest struct {
	Object ObjectRef `xml:"object"`
	State  StateRef  `xml:"state"`
}

type FamilyTest struct {
	Object ObjectRef `xml:"object"`
	State  StateRef  `xml:"state"`
}

type UnknownTest Test
