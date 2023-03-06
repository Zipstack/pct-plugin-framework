package schema

type NestedBlock struct {
	// Defnition for the child block.
	// Both NestedBlock and Attributes cannot be specified.
	Block map[string]Block

	// Definitions for the key value arguments.
	Attributes map[string]Attribute

	// Whether the value must be present in configuration or not.
	Required bool

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

func (s NestedBlock) IsRequired() bool {
	return s.Required
}

func (s NestedBlock) GetDescription() string {
	return s.Description
}

func (s NestedBlock) GetDeprecationMessage() string {
	return s.DeprecationMessage
}
