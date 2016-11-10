package models

type UnameTest struct {
	Object ObjectRef `xml:"object"`
	State  StateRef  `xml:"state"`
}
