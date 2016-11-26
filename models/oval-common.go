package models

import (
	"github.com/arzonus/ovalgo/utils/ovaltime"
)

// GeneratorType -> Generator
type Generator struct {
	ProductName    string        `xml:"product_name"`
	ProductVersion string        `xml:"product_version"`
	SchemaVersion  string        `xml:"schema_version"`
	Timestamp      ovaltime.Time `xml:"timestamp"`
	Any            string        `xml:"xsd:any"`
}

// MessageType -> Message
type Message struct {
	Level   string
	Message string
}
