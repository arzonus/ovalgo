package types

import "time"

type OVALDefinitionID string
type OVALObjectID string
type OVALStateID string
type OVALTestID string

type OVALEnumeration string
type OVALClassEnumeration string
type OVALCheckEnumeration string
type OVALOperatorEnumeration string
type OVALExistenceEnumeration string
type OVALFilterActionEnumeration string

type OVALNonEmptyString string

type OVALEntityObjectString string
type OVALEntityStateAnySimple string
type OVALEntityStateEVRString string
type OVALEntityStateString string

type URI string

type OVALDefinitions struct {
	Generator   Generator    `xml:"generator"`
	Definitions []Definition `xml:"definitions>definition"`
	Tests       []Test       `xml:"tests>test"`
	Objects     []Object     `xml:"objects>object"`
	States      []State      `xml:"states>state"`
}

type Generator struct {
	ProductName    string    `xml:"product_name"`
	ProductVersion string    `xml:"product_version"`
	SchemaVersion  string    `xml:"schema_version"`
	Timestamp      time.Time `xml:"timestamp"`
	Any            string    `xml:"xsd:any"`
}

type Definition struct {
	ID         OVALDefinitionID     `xml:"id,attr"`
	Class      OVALClassEnumeration `xml:"class,attr"`
	Version    uint32               `xml:"version,attr"`
	Deprecated bool                 `xml:"deprecated,attr"`
	Metadata   Metadata             `xml:"metadata"`
	Notes      []Note               `xml:"notes>note"`
	Criteria   Criteria             `xml:"criteria"`
}

type Metadata struct {
	Title       string      `xml:"title"`
	Affected    []Affected  `xml:"affected"`
	References  []Reference `xml:"reference"`
	Description string      `xml:"description"`
}

type Affected struct {
	Family    OVALEnumeration `xml:"family,attr"`
	Platforms []string        `xml:"platform"`
	Products  []string        `xml:"product"`
}

type Reference struct {
	Source string `xml:"source,attr"`
	RefID  string `xml:"ref_id,attr"`
	RefURL URI    `xml:"ref_url,attr"`
}

type Note string

type Criteria struct {
	Operator           OVALOperatorEnumeration `xml:"operator,attr"`
	Negate             bool                    `xml:"negate,attr"`
	Comment            OVALNonEmptyString      `xml:"comment,attr"`
	Criterias          []Criteria              `xml:"criteria"`
	Criterions         []Criterion             `xml:"criterion"`
	ExtendDefinitions  []ExtendDefinition      `xml:"extend_definition"`
	ApplicabilityCheck bool                    `xml:"applicability_check,attr"`
}

type Criterion struct {
	TestRef            OVALTestID         `xml:"test_ref,attr"`
	Negate             bool               `xml:"negate,attr"`
	Comment            OVALNonEmptyString `xml:"comment,attr"`
	ApplicabilityCheck bool               `xml:"applicability_check,attr"`
}

type ExtendDefinition struct {
	DefinitionRef      OVALDefinitionID   `xml:"definition_ref,attr"`
	Negate             bool               `xml:"negate,attr"`
	Comment            OVALNonEmptyString `xml:"comment,attr"`
	ApplicabilityCheck bool               `xml:"applicability_check,attr"`
}

type Test struct {
	ID             OVALTestID               `xml:"id,attr"`
	Version        uint32                   `xml:"version,attr"`
	CheckExistence OVALExistenceEnumeration `xml:"check_existence"`
	Check          OVALCheckEnumeration     `xml:"check"`
	StateOperator  OVALOperatorEnumeration  `xml:"state_operator"`
	Comment        OVALNonEmptyString       `xml:"comment,attr"`
	Deprecated     bool                     `xml:"deprecated,attr"`
	Notes          []Note                   `xml:"notes>note"`
}

type Object struct {
	ID         OVALObjectID       `xml:"id,attr"`
	Version    uint32             `xml:"version,attr"`
	Comment    OVALNonEmptyString `xml:"comment,attr"`
	Deprecated bool               `xml:"deprecated,attr"`
	Notes      []Note             `xml:"notes>note"`
}

type State struct {
	ID         OVALStateID             `xml:"id,attr"`
	Operator   OVALOperatorEnumeration `xml:"operator,attr"`
	Version    uint32                  `xml:"version,attr"`
	Comment    OVALNonEmptyString      `xml:"comment,attr"`
	Deprecated bool                    `xml:"deprecated,attr"`
	Notes      []Note                  `xml:"notes>note"`
}

type Filter struct {
	Action OVALFilterActionEnumeration
	Value  OVALStateID
}
