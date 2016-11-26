package models

// Short types Enumeration, Strings and etc

type OVALDefinitions struct {
	Generator   Generator    `xml:"generator"`
	Definitions []Definition `xml:"definitions>definition"`
	Tests       []Test       `xml:"tests>test"`
	Objects     []Object     `xml:"objects>object"`
	States      []State      `xml:"states>state"`
	Variables   []Variable   `xml:"variables>variable"`
}

type Definition struct {
	ID         string   `xml:"id,attr"`
	Class      string   `xml:"class,attr"`
	Version    uint     `xml:"version,attr"`
	Deprecated bool     `xml:"deprecated,attr"`
	Metadata   Metadata `xml:"metadata"`
	Notes      []string `xml:"notes>note"`
	Criteria   Criteria `xml:"criteria"`
}

type Metadata struct {
	Title       string      `xml:"title"`
	Affected    []Affected  `xml:"affected"`
	References  []Reference `xml:"reference"`
	Description string      `xml:"description"`
}

type Affected struct {
	Family    string   `xml:"family,attr"`
	Platforms []string `xml:"platform"`
	Products  []string `xml:"product"`
}

type Reference struct {
	Source string `xml:"source,attr"`
	RefID  string `xml:"ref_id,attr"`
	RefURL string `xml:"ref_url,attr"`
}

type Note string

type Criteria struct {
	Operator           string             `xml:"operator,attr"`
	Negate             bool               `xml:"negate,attr"`
	Comment            string             `xml:"comment,attr"`
	Criterias          []Criteria         `xml:"criteria"`
	Criterions         []Criterion        `xml:"criterion"`
	ExtendDefinitions  []ExtendDefinition `xml:"extend_definition"`
	ApplicabilityCheck bool               `xml:"applicability_check,attr"`
}

type Criterion struct {
	TestRef            string `xml:"test_ref,attr"`
	Negate             bool   `xml:"negate,attr"`
	Comment            string `xml:"comment,attr"`
	ApplicabilityCheck bool   `xml:"applicability_check,attr"`
}

type ExtendDefinition struct {
	DefinitionRef      string `xml:"definition_ref,attr"`
	Negate             bool   `xml:"negate,attr"`
	Comment            string `xml:"comment,attr"`
	ApplicabilityCheck bool   `xml:"applicability_check,attr"`
}

type Test struct {
	ID             string   `xml:"id,attr"`
	Version        uint     `xml:"version,attr"`
	CheckExistence string   `xml:"check_existence,attr"`
	Check          string   `xml:"check,attr"`
	StateOperator  string   `xml:"state_operator,attr"`
	Comment        string   `xml:"comment,attr"`
	Deprecated     bool     `xml:"deprecated,attr"`
	Notes          []string `xml:"notes>note"`
}

type Object struct {
	ID         string   `xml:"id,attr"`
	Version    uint     `xml:"version,attr"`
	Comment    string   `xml:"comment,attr"`
	Deprecated bool     `xml:"deprecated,attr"`
	Notes      []string `xml:"notes>note"`
}

type State struct {
	ID         string   `xml:"id,attr"`
	Operator   string   `xml:"operator,attr"`
	Version    uint     `xml:"version,attr"`
	Comment    string   `xml:"comment,attr"`
	Deprecated bool     `xml:"deprecated,attr"`
	Notes      []string `xml:"notes>note"`
}

type Set struct {
	SetOperator      string
	Sets             []Set
	ObjectReferences []string
	Filters          []Filter
}

type Filter struct {
	Action string
	Value  string
}

type Variable struct {
	ID         string
	Version    uint32
	Datatype   string
	Comment    string
	Deprecated bool
}

type ExternalVariable struct {
	PossibleValue       PossibleValue
	PossibleRestriction PossibleRestriction
}

type PossibleValue struct {
	Value string
	Hint  string
}

type PossibleRestriction struct {
	Restriction []Restriction
	Operation   string
	Hint        string
}

type Restriction struct {
	Operation string
	Value     string
}

type ConstantVariable struct {
	Values []Value
}

type Value struct {
	value string
}

type LocalVariable struct {
	Components []ComponentGroup
}

type ComponentGroup struct {
	ObjectComponents   []ObjectComponent
	VariableComponents []VariableComponent
	LiteralComponent   []LiteralComponent
	Functions          []FunctionComponent
}

type LiteralComponent struct {
	Datatype string
	Value    string
}

type ObjectComponent struct {
	ItemField   string
	RecordField string
	ObjectRef   string
}

type VariableComponent struct {
	VarRef string
}

type FunctionComponent struct {
	Arithmetic     ArithmeticFunction
	Begin          BeginFunction
	Concat         ConcatFunction
	Count          CountFunction
	End            EndFunction
	EscapeRegex    EscapeRegexFunction
	Split          SplitFunction
	Substring      SubstringFunction
	TimeDifference TimeDifferenceFunction
	Unique         UniqueFunction
	RegexCapture   RegexCaptureFunction
}

type ArithmeticFunction struct {
	ArithmeticOperation string
	Values              []ComponentGroup
}

type BeginFunction struct {
	Character string
	Value     ComponentGroup
}
type ConcatFunction struct {
	Values []ComponentGroup
}
type CountFunction struct {
	Values []ComponentGroup
}
type EndFunction struct {
	Character string
	Value     ComponentGroup
}
type EscapeRegexFunction struct {
	Value ComponentGroup
}
type SplitFunction struct {
	Delimiter string
	Value     ComponentGroup
}
type SubstringFunction struct {
	SubstringStart  int
	SubstringLength int
	Value           ComponentGroup
}
type TimeDifferenceFunction struct {
	Format1 string
	Format2 string
	Values  []ComponentGroup
}

type UniqueFunction struct {
	Values []ComponentGroup
}
type RegexCaptureFunction struct {
	Pattern string
	Value   ComponentGroup
}

type EntityAttributeGroup struct {
	Datatype  string
	Operation string
	Mask      bool
	VarRef    string
	CheckRef  string
}
