package parser

type SwaggerInfo struct {
	Title          *string             `json:"title,omitempty"`
	Description    *string             `json:"description,omitempty"`
	TermsOfService *string             `json:"termsOfService,omitempty"`
	Contact        *SwaggerInfoContact `json:"contact,omitempty"`
	License        *SwaggerInfoLicense `json:"license,omitempty"`
	Version        *string             `json:"version,omitempty"`
}
