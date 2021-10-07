package model

type Policy struct {
	Subject string `json:"role" validate:"required"`
	Object  string `json:"path" validate:"required"`
	Action  string `json:"action" validate:"required"`
	CompanyID string `json:"company_id,omitempty"`
}

type PolicyUpdate struct {
	Old Policy `json:"old,omitempty" validate:"required"`
	New Policy `json:"new,omitempty" validate:"required"`
}
