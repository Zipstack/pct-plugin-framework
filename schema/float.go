package schema

type Float64Attribute struct {
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

type FloatAttribute = Float64Attribute

func (s FloatAttribute) IsRequired() bool {
	return s.Required
}

func (s FloatAttribute) IsComputed() bool {
	return s.Computed
}

func (s FloatAttribute) GetDescription() string {
	return s.Description
}

func (s FloatAttribute) GetDeprecationMessage() string {
	return s.DeprecationMessage
}
