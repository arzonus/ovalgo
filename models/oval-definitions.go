package models

// Short types Enumeration, Strings and etc

type Enumeration string
type FilterActionEnumeration string

type EntityObjectString string
type EntityStateAnySimple string
type EntityStateEVRString string
type EntityStateString string

type SetOperatorEnumeration string
type DatatypeFormatEnumeration string
type ArithmeticEnumeration string

type OVALDefinitions struct {
	Generator   Generator    `xml:"generator"`
	Definitions []Definition `xml:"definitions>definition"`
	Tests       []Test       `xml:"tests>test"`
	Objects     []Object     `xml:"objects>object"`
	States      []State      `xml:"states>state"`
	Variables   []Variable   `xml:"variables>variable"`
}

type Definition struct {
	ID         DefinitionID     `xml:"id,attr"`
	Class      ClassEnumeration `xml:"class,attr"`
	Version    uint32           `xml:"version,attr"`
	Deprecated bool             `xml:"deprecated,attr"`
	Metadata   Metadata         `xml:"metadata"`
	Notes      []Note           `xml:"notes>note"`
	Criteria   Criteria         `xml:"criteria"`
}

type Metadata struct {
	Title       string      `xml:"title"`
	Affected    []Affected  `xml:"affected"`
	References  []Reference `xml:"reference"`
	Description string      `xml:"description"`
}

type Affected struct {
	Family    Enumeration `xml:"family,attr"`
	Platforms []string    `xml:"platform"`
	Products  []string    `xml:"product"`
}

type Reference struct {
	Source string `xml:"source,attr"`
	RefID  string `xml:"ref_id,attr"`
	RefURL URI    `xml:"ref_url,attr"`
}

type Note string

type Criteria struct {
	Operator           OperatorEnumeration `xml:"operator,attr"`
	Negate             bool                `xml:"negate,attr"`
	Comment            NonEmptyString      `xml:"comment,attr"`
	Criterias          []Criteria          `xml:"criteria"`
	Criterions         []Criterion         `xml:"criterion"`
	ExtendDefinitions  []ExtendDefinition  `xml:"extend_definition"`
	ApplicabilityCheck bool                `xml:"applicability_check,attr"`
}

type Criterion struct {
	TestRef            TestID         `xml:"test_ref,attr"`
	Negate             bool           `xml:"negate,attr"`
	Comment            NonEmptyString `xml:"comment,attr"`
	ApplicabilityCheck bool           `xml:"applicability_check,attr"`
}

type ExtendDefinition struct {
	DefinitionRef      DefinitionID   `xml:"definition_ref,attr"`
	Negate             bool           `xml:"negate,attr"`
	Comment            NonEmptyString `xml:"comment,attr"`
	ApplicabilityCheck bool           `xml:"applicability_check,attr"`
}

type Test struct {
	ID             TestID               `xml:"id,attr"`
	Version        uint32               `xml:"version,attr"`
	CheckExistence ExistenceEnumeration `xml:"check_existence,attr"`
	Check          CheckEnumeration     `xml:"check,attr"`
	StateOperator  OperatorEnumeration  `xml:"state_operator,attr"`
	Comment        NonEmptyString       `xml:"comment,attr"`
	Deprecated     bool                 `xml:"deprecated,attr"`
	Notes          []Note               `xml:"notes>note"`
}

type Object struct {
	ID         ObjectID       `xml:"id,attr"`
	Version    uint32         `xml:"version,attr"`
	Comment    NonEmptyString `xml:"comment,attr"`
	Deprecated bool           `xml:"deprecated,attr"`
	Notes      []Note         `xml:"notes>note"`
}

type State struct {
	ID         StateID             `xml:"id,attr"`
	Operator   OperatorEnumeration `xml:"operator,attr"`
	Version    uint32              `xml:"version,attr"`
	Comment    NonEmptyString      `xml:"comment,attr"`
	Deprecated bool                `xml:"deprecated,attr"`
	Notes      []Note              `xml:"notes>note"`
}

type Set struct {
	SetOperator      SetOperatorEnumeration
	Sets             []Set
	ObjectReferences []ObjectID
	Filters          []Filter
}

type Filter struct {
	Action FilterActionEnumeration
	Value  StateID
}

type Variable struct {
	ID         VariableID
	Version    uint32
	Datatype   DatatypeEnumeration
	Comment    NonEmptyString
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
	Operation   OperationEnumeration
	Hint        string
}

type Restriction struct {
	Operation OperationEnumeration
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
	Datatype SimpleDataEnumeration
	Value    string
}

type ObjectComponent struct {
	ItemField   NonEmptyString
	RecordField NonEmptyString
	ObjectRef   ObjectID
}

type VariableComponent struct {
	VarRef VariableID
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
	ArithmeticOperation ArithmeticEnumeration
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
	Format1 DatatypeFormatEnumeration
	Format2 DatatypeFormatEnumeration
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
	Datatype  DatatypeEnumeration
	Operation OperationEnumeration
	Mask      bool
	VarRef    VariableID
	CheckRef  CheckEnumeration
}
