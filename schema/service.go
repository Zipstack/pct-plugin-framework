package schema

type ServiceRequest struct {
	TypeName       string
	ConfigContents string
	PlanContents   string
	StateContents  string
	ResourceData   string
}

type ServiceResponse struct {
	TypeName         string
	SchemaContents   string
	StateID          string
	StateContents    string
	StateLastUpdated string
	ResourceServices map[string]string
	ResourceData     string
	ErrorsContents   string
}

type ProviderService interface {
	Metadata(req *ServiceRequest) *ServiceResponse
	Schema() *ServiceResponse
	Configure(*ServiceRequest) *ServiceResponse
	Resources() *ServiceResponse
	UpdateResourceServices(map[string]string)
}

type ResourceService interface {
	Metadata(req *ServiceRequest) *ServiceResponse
	Configure(*ServiceRequest) *ServiceResponse
	Schema() *ServiceResponse
	Create(*ServiceRequest) *ServiceResponse
	Read(*ServiceRequest) *ServiceResponse
	Update(*ServiceRequest) *ServiceResponse
	Delete(*ServiceRequest) *ServiceResponse
}

func ErrorResponse(err error) *ServiceResponse {
	return &ServiceResponse{
		TypeName:       "",
		SchemaContents: "",
		StateContents:  "",
		ErrorsContents: err.Error(),
	}
}
