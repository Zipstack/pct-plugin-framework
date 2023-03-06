package schema

type BoolAttribute struct {
	// Whether the value must be present in configuration or not.
	// Both Required and Computed cannot be true.
	Required bool

	// Whether the configuration value is set by the provider or not.
	// Both Required and Computed cannot be true.
	Computed bool

	// Whether the configuration value is optional or not.
	Optional bool

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

func (s BoolAttribute) IsRequired() bool {
	return s.Required
}

func (s BoolAttribute) IsComputed() bool {
	return s.Computed
}

func (s BoolAttribute) IsOptional() bool {
	return s.Optional
}

func (s BoolAttribute) GetDescription() string {
	return s.Description
}

func (s BoolAttribute) GetDeprecationMessage() string {
	return s.DeprecationMessage
}
