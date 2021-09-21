package model

type Permision struct {
	Subject string `json:"subject,omitempty" binding:"required"`
	Object  string `json:"object,omitempty" binding:"required"`
	Action  string `json:"action,omitempty" binding:"required"`
}
