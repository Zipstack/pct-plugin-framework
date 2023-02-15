package schema

type StringAttribute struct {
	// Whether a value must be entered or not.
	// Both Required and Computed cannot be true.
	Required bool

	// Only the provider is able to set its value.
	// Both Required and Computed cannot be true.
	Computed bool

	// Setting it to true will obscure the value in CLI output.
	// It does not affect how the values are stored.
	Sensitive bool

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

func (s StringAttribute) IsRequired() bool {
	return s.Required
}

func (s StringAttribute) IsComputed() bool {
	return s.Computed
}

func (s StringAttribute) IsSensitive() bool {
	return s.Sensitive
}

func (s StringAttribute) GetDescription() string {
	return s.Description
}

func (s StringAttribute) GetDeprecationMessage() string {
	return s.DeprecationMessage
}
