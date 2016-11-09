package models

import "time"

type URI string

// Enumeration types
type CheckEnumeration string
type ClassEnumeration string
type SimpleDataEnumeration string
type ComplexDatatypeEnumeration string
type DatatypeEnumeration string
type ExistenceEnumeration string
type MessageLevelEnumeration string
type FamilyEnumeration string
type OperationEnumeration string
type OperatorEnumeration string

// ID Patterns
type DefinitionID string
type VariableID string
type ObjectID string
type StateID string
type TestID string

// Strings
type EmptyString string
type NonEmptyString string

// GeneratorType -> Generator
type Generator struct {
	ProductName    string    `xml:"product_name"`
	ProductVersion string    `xml:"product_version"`
	SchemaVersion  string    `xml:"schema_version"`
	Timestamp      time.Time `xml:"timestamp"`
	Any            string    `xml:"xsd:any"`
}

// MessageType -> Message
type Message struct {
	Level   MessageLevelEnumeration
	Message string
}
