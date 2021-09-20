package model

type Permision struct {
	Subject    string `json:"role" validate:"required"`
	Object    string `json:"path" validate:"required"`
	Action   string `json:"action" validate:"required"`
}