package schema

type MapAttribute struct {
	// Definitions for the key value arguments.
	Attributes map[string]Attribute

	// Whether the value must be present in configuration or not.
	// Both Required and Computed cannot be true.
	Required bool

	// Whether the configuration value is set by the provider or not.
	// Both Required and Computed cannot be true.
	Computed bool

	// Whether the configuration value is optional or not.
	Optional bool

	// Exactly one of the given attributes should be present.
	ExactlyOneOf []string

	// Plain text description that can be used in various tooling.
	Description string

	// Warning messages to display if this arribute is getting
	// deprecated in an upcoming version.
	DeprecationMessage string

	// TODO One or more validators for the type.
	// Validators []Validator

	// TODO Schema based modifications which can alter the plan.
	// Modifiers []Modifier
}

func (s MapAttribute) IsRequired() bool {
	return s.Required
}

func (s MapAttribute) IsComputed() bool {
	return s.Computed
}

func (s MapAttribute) IsOptional() bool {
	return s.Optional
}

func (s MapAttribute) GetDescription() string {
	return s.Description
}

func (s MapAttribute) GetDeprecationMessage() string {
	return s.DeprecationMessage
}
