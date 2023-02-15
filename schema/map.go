package schema

type MapAttribute struct {
	// Definitions for the key value arguments.
	Attributes map[string]Attribute

	// Whether a value must be entered or not.
	// Both Required and Computed cannot be true.
	Required bool

	// Only the provider is able to set its value.
	// Both Required and Computed cannot be true.
	Computed bool

	// Plain text description that can be used in various tooling.
	Description string

	// Warning messages to display if this arribute is getting
	// deprecated in an upcoming version.
	DeprecationMessage string

	// TODO One or more validators for the type.
	// Validators []Validator

	// TODO Schema based modifications which can alter the plan.
	// Modifiers []Planmodifier
}

func (s MapAttribute) IsRequired() bool {
	return s.Required
}

func (s MapAttribute) IsComputed() bool {
	return s.Computed
}

func (s MapAttribute) GetDescription() string {
	return s.Description
}

func (s MapAttribute) GetDeprecationMessage() string {
	return s.DeprecationMessage
}
