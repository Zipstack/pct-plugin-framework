package schema

type Schema struct {
	// All the attributes.
	Attributes map[string]Attribute

	// All the attributes.
	Blocks map[string]Block

	// Plain text description that can be used in various tooling.
	Description string

	// Warning messages to display if this schema is getting
	// deprecated in an upcoming version.
	DeprecationMessage string

	// Current version of the schema.
	Version int64
}

type Attribute interface {
	GetDeprecationMessage() string
	GetDescription() string
	IsRequired() bool
	IsComputed() bool
	IsOptional() bool
}

type Block interface {
	GetDeprecationMessage() string
	GetDescription() string
	IsRequired() bool
}
